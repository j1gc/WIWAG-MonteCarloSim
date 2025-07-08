package routes

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"slices"
	"sync"
	"time"
	"warehouse/histogram"
	"warehouse/simulation"
	"warehouse/utils"
)

//type Handlers struct {
//	DB db.Database
//}

func RandomNormFloatInRange(maxValue float64, minValue float64) float64 {
	// Standard deviation
	// so that 99,7 % of random values fall into the range by calculating the formular of the third sigma interval:
	// solve(0+3*stdDev = 1, stdDev)
	// otherwise by default a stdDev of 1*stdDev would make it so that only 68 % of values fall into that range
	// if the value still doesn't fall into the truncated range then use the accept rejection method

	mean := (maxValue + minValue) / 2
	stdDev := (maxValue - mean) / 3

	// accept rejection algorithm
	for {
		randomValue := rand.NormFloat64()*stdDev + mean

		// check if the value is in the range: min <= x <= max
		if randomValue <= maxValue && randomValue >= minValue {
			return randomValue
		}
	}
}

func RandomNormFloatInPercentageRangeUnder(maxValue float64, percentage float64) float64 {
	// prevents that an infinit loop is created when providing a value outside this range
	if percentage > 1 || percentage < 0 {
		// programm should fail to get faster feedback on the error
		log.Fatalf("Percentage must be in domain [0, 1], but was provided %f", percentage)
	}

	minValue := maxValue * (1 - percentage)

	return RandomNormFloatInRange(maxValue, minValue)
}

func runSimulationStep(input simulation.SimInput) simulation.SimResults {
	currentInput := input

	originalAbsatzmengeInlandDetailhandelHoheQualität := input.AbsatzmengeInlandDetailhandelHoheQualität
	originalAbsatzmengeInlandOnlineshopHoheQualität := input.AbsatzmengeInlandOnlineshopHoheQualität

	originalAbsatzmengeAuslandDetailhandelHoheQualität := input.AbsatzmengeAuslandDetailhandelHoheQualität
	originalAbsatzmengeAuslandOnlineshopHoheQualität := input.AbsatzmengeAuslandOnlineshopHoheQualität

	originalMaterialverbrauch := input.Materialverbrauch

	detailhandelInland := RandomNormFloatInPercentageRangeUnder(originalAbsatzmengeInlandDetailhandelHoheQualität, 0.2)

	detailhandelAusland := RandomNormFloatInPercentageRangeUnder(originalAbsatzmengeAuslandDetailhandelHoheQualität, 0.3)
	onlineShopInland := RandomNormFloatInPercentageRangeUnder(originalAbsatzmengeInlandOnlineshopHoheQualität, 0.2)
	onlineShopAusland := RandomNormFloatInPercentageRangeUnder(originalAbsatzmengeAuslandOnlineshopHoheQualität, 0.3)
	materialVerbrauch := RandomNormFloatInPercentageRangeUnder(originalMaterialverbrauch, 0.3)

	currentInput.AbsatzmengeInlandDetailhandelHoheQualität = detailhandelInland
	currentInput.AbsatzmengeAuslandDetailhandelHoheQualität = detailhandelAusland

	currentInput.AbsatzmengeInlandOnlineshopHoheQualität = onlineShopInland
	currentInput.AbsatzmengeAuslandOnlineshopHoheQualität = onlineShopAusland

	currentInput.Materialverbrauch = materialVerbrauch

	simulationData := simulation.SimData{
		Input:   currentInput,
		Results: simulation.SimResults{},
	}
	results := simulationData.GetResults()

	//stückSelbstkosten := results.Selbstkosten / (results.AbsatzmengeInlandGesamt + results.AbsatzmengeAuslandGesamt)
	return results
}

func runMonteCarloSim(steps int, simInput simulation.SimInput, batchSize int) []simulation.SimResults {
	var results []simulation.SimResults

	// a mutex is needed because the result slice is being written to by multiple go routines
	var wg sync.WaitGroup
	var mu sync.Mutex
	for batchStart := 0; batchStart < steps; batchStart += batchSize {
		batchEnd := batchStart + batchSize
		// check if the last batch is reached
		// if is, set batchEnd to steps so that the last batchSize is smaller
		if batchEnd > steps {
			batchEnd = steps
		}

		wg.Add(1)
		// start up a go routine for every batch
		go func(batchStart int, batchEnd int) {
			defer wg.Done()

			currentBatchSize := batchEnd - batchStart
			currentBatchResults := make([]simulation.SimResults, currentBatchSize)

			for i := 0; i < currentBatchSize; i++ {
				currentBatchResults[i] = runSimulationStep(simInput)
			}

			mu.Lock()
			results = append(results, currentBatchResults...)
			mu.Unlock()

		}(batchStart, batchEnd)
	}
	wg.Wait()

	return results
}

func RunSimulation(c echo.Context) error {
	defaultSimInput := simulation.InitYear14()

	c.Logger().Print("Connection!")
	simulationSteps := 100_000 // (defaultSimInput.getAbsatzmengeAuslandGesamt() + defaultSimInput.getAbsatzmengeInlandGesamt())

	// Parse JSON body
	err := json.NewDecoder(c.Request().Body).Decode(&simulationSteps)
	if err != nil {
		c.Logger().Error("Decode error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input JSON")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.Logger().Error("Decode error:", err)
		}
	}(c.Request().Body)

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("Cache-Control", "no-cache")

	select {
	case <-c.Request().Context().Done(): // Handle client disconnect
		c.Logger().Print("Client disconnected")
		return nil
	default:
		timeStart := time.Now()
		batchSize := utils.CalculateOptimalBatchSizeForCPUTasks(simulationSteps)
		simResults := runMonteCarloSim(simulationSteps, defaultSimInput, batchSize)

		// reports the time that the Monte Carlo Simulation took
		c.Logger().Print("Simulation took:" + time.Now().Sub(timeStart).String())

		var resultStück []float64
		for _, val := range simResults {
			resultStück = append(resultStück, val.Selbstkosten/(val.AbsatzmengeInlandGesamt+val.AbsatzmengeAuslandGesamt))
		}

		minV := slices.Min(resultStück)
		maxV := slices.Max(resultStück)
		binAmount := histogram.CalculateNumberOfBins(simulationSteps)

		binEdges := histogram.CalculateBinEdges(binAmount, minV, maxV)
		binIndexes := histogram.Digitize(resultStück, binEdges)
		bins := histogram.CreateBins(resultStück, binIndexes, binEdges, binAmount)

		return c.JSON(http.StatusOK, bins)

		//go func() {
		//	insertTime := time.Now()
		//	err = h.DB.InsertSimulationsInDB(simResults)
		//	fmt.Println("Insert took:", time.Now().Sub(insertTime).String())
		//}()
		//if err != nil {
		//	return err
		//}
	}
}
