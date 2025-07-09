package utils

import (
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"warehouse/simulation"
)

func GetEnvBool(key string, defaultValue bool) bool {
	envString, exist := os.LookupEnv(key)
	if !exist {
		log.Printf("Environment variable %s was not set. Using value: %s for it instead", key, strconv.FormatBool(defaultValue))
		return defaultValue
	}
	envBool, err := strconv.ParseBool(envString)
	if err != nil {
		log.Printf("Error parsing environment variable %s: %s\n", key, err)
		return defaultValue
	}
	return envBool
}

func GetEnvString(key string, defaultValue string) string {
	envString, exist := os.LookupEnv(key)
	if !exist {
		log.Printf("Environment variable %s was not set. Using value: %s for it instead", key, defaultValue)
		return defaultValue
	}
	return envString

}

func IsRunningInDocker() bool {
	_, exist := os.LookupEnv("DOCKER")
	if !exist {
		return false
	}
	return true
}

func CalculateOptimalBatchSizeForCPUTasks(taskAmount int) int {
	// Calculates the right batch size by dividing the steps through the number of
	// threads and rounding the result up.
	// A bigger or smaller batch size shouldn't be optimal since the task is CPU-bound
	batchSize := int(math.Ceil(float64(taskAmount / runtime.NumCPU())))
	return batchSize
}

func PrintSelbstkostenTableYear14(writer io.Writer, simData simulation.SimData) error {
	//p := message.NewPrinter(language.English)
	t := NewTablePrinter()

	t.AddHeader("Name", "Value", "Expected", "% Error")

	// Calculates the ingredients for the unit cost
	materialaufwand := simData.Results.MaterialAufwand / simData.Results.AbsatzmengeGesamt
	produktionsaufwand := simData.Results.Produktionsaufwand / simData.Results.AbsatzmengeGesamt
	übrigerPersonalaufwand := simData.Results.PersonalAufwandÜbrigesPersonal / simData.Results.AbsatzmengeGesamt
	werbeaufwand := simData.Results.Werbeaufwand / simData.Results.AbsatzmengeGesamt
	distributionsaufwand := simData.Results.Distributionsaufwand / simData.Results.AbsatzmengeGesamt
	lageraufwand := simData.Results.AufwandLagerkosten / simData.Results.AbsatzmengeGesamt
	forschungUndEntwicklung := simData.Input.AufwandForschungEntwicklung / simData.Results.AbsatzmengeGesamt
	weiterbildungsaufwand := simData.Results.WeiterbildungsAufwand / simData.Results.AbsatzmengeGesamt
	verwaltungsaufwand := simData.Results.Verwaltungskosten / simData.Results.AbsatzmengeGesamt
	übrigerBetrieblicherAufwand := simData.Input.ÜbrigerBetrieblicherAufwand / simData.Results.AbsatzmengeGesamt
	abschreibungen := simData.Results.AbschreibungenLagerräumeUndÜbriges / simData.Results.AbsatzmengeGesamt
	zinsaufwand := simData.Results.Zinsaufwand / simData.Results.AbsatzmengeGesamt
	steuern := simData.Results.Steueraufwand / simData.Results.AbsatzmengeGesamt
	total := simData.Results.Selbstkosten / simData.Results.AbsatzmengeGesamt

	expected := simulation.ExpectedValuesYear14()

	// Calculates the percentage error
	materialaufwandError := math.Abs((materialaufwand - expected.Materialaufwand) / expected.Materialaufwand * 100)
	produktionsaufwandError := math.Abs((produktionsaufwand - expected.Produktionsaufwand) / expected.Produktionsaufwand * 100)
	übrigerPersonalaufwandError := math.Abs((übrigerPersonalaufwand - expected.ÜbrigerPersonalaufwand) / expected.ÜbrigerPersonalaufwand * 100)
	werbeaufwandError := math.Abs((werbeaufwand - expected.Werbeaufwand) / expected.Werbeaufwand * 100)
	distributionsaufwandError := math.Abs((distributionsaufwand - expected.Distributionsaufwand) / expected.Distributionsaufwand * 100)
	lageraufwandError := math.Abs((lageraufwand - expected.Lageraufwand) / expected.Lageraufwand * 100)
	forschungUndEntwicklungError := math.Abs((forschungUndEntwicklung - expected.ForschungUndEntwicklung) / expected.ForschungUndEntwicklung * 100)
	weiterbildungsaufwandError := math.Abs((weiterbildungsaufwand - expected.Weiterbildungsaufwand) / expected.Weiterbildungsaufwand * 100)
	verwaltungsaufwandError := math.Abs((verwaltungsaufwand - expected.Verwaltungsaufwand) / expected.Verwaltungsaufwand * 100)
	übrigerBetrieblicherAufwandError := math.Abs((übrigerBetrieblicherAufwand - expected.ÜbrigerBetrieblicherAufwand) / expected.ÜbrigerBetrieblicherAufwand * 100)
	abschreibungenError := math.Abs((abschreibungen - expected.Abschreibungen) / expected.Abschreibungen * 100)
	zinsaufwandError := math.Abs((zinsaufwand - expected.Zinsaufwand) / expected.Zinsaufwand * 100)
	steuernError := math.Abs((steuern - expected.Steuern) / expected.Steuern * 100)
	totalError := math.Abs((total - expected.Total) / expected.Total * 100)

	t.AddRow("Materialaufwand", materialaufwand, expected.Materialaufwand, materialaufwandError)
	t.AddRow("Produktionsaufwand", produktionsaufwand, expected.Produktionsaufwand, produktionsaufwandError)
	t.AddRow("Übriger Personalaufwand", übrigerPersonalaufwand, expected.ÜbrigerPersonalaufwand, übrigerPersonalaufwandError)
	t.AddRow("Werbeaufwand", werbeaufwand, expected.Werbeaufwand, werbeaufwandError)
	t.AddRow("Distributionsaufwand", distributionsaufwand, expected.Distributionsaufwand, distributionsaufwandError)
	t.AddRow("Lageraufwand", lageraufwand, expected.Lageraufwand, lageraufwandError)
	t.AddRow("Aufwand für Forschung und Entwicklung", forschungUndEntwicklung, expected.ForschungUndEntwicklung, forschungUndEntwicklungError)
	t.AddRow("Aus- und Weiterbildungsaufwand", weiterbildungsaufwand, expected.Weiterbildungsaufwand, weiterbildungsaufwandError)
	t.AddRow("Verwaltungsaufwand", verwaltungsaufwand, expected.Verwaltungsaufwand, verwaltungsaufwandError)
	t.AddRow("Übriger betrieblicher Aufwand", übrigerBetrieblicherAufwand, expected.ÜbrigerBetrieblicherAufwand, übrigerBetrieblicherAufwandError)
	t.AddRow("Abschreibungen Lagerräume und übrige Anlagen", abschreibungen, expected.Abschreibungen, abschreibungenError)
	t.AddRow("Zinsaufwand", zinsaufwand, expected.Zinsaufwand, zinsaufwandError)
	t.AddRow("Steuern", steuern, expected.Steuern, steuernError)
	t.AddRow("-", "-", "-", "-")
	t.AddRow("Total", total, expected.Total, totalError)

	return t.Print(writer)

}
