package routes

import (
	"fmt"
	"os"
	"testing"
	"warehouse/simulation"
	"warehouse/utils"
)

func TestRunMonteCarloSim(t *testing.T) {
	simInput := simulation.InitYear14()

	t.Run("RunMonteCarloSim Stückselbstkosten", func(t *testing.T) {
		simData := simulation.SimData{
			Input:   simInput,
			Results: simulation.SimResults{},
		}
		simData.GetResults()

		StückSelbstkosten := simData.Results.Selbstkosten / (simData.Results.AbsatzmengeInlandGesamt + simData.Results.AbsatzmengeAuslandGesamt)

		t.Log("Got:", StückSelbstkosten)

		err := utils.PrintSelbstkostenTableYear14(os.Stdout, simData)
		if err != nil {
			t.Error(err)
		}
	})
}

func BenchSingleRunMonteCarloSim(b *testing.B, steps int, simInput simulation.SimInput) {
	b.Run(fmt.Sprintf("%d steps", steps), func(b *testing.B) {
		batchSize := utils.CalculateOptimalBatchSizeForCPUTasks(steps)

		for b.Loop() {
			runMonteCarloSim(steps, simInput, batchSize)
		}
	})
}

func BenchmarkRunMonteCarloSim(b *testing.B) {
	simInput := simulation.InitYear14()

	BenchSingleRunMonteCarloSim(b, 100_000, simInput)
	BenchSingleRunMonteCarloSim(b, 1_000_000, simInput)
	BenchSingleRunMonteCarloSim(b, 5_000_000, simInput)
}
