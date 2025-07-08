package histogram

import (
	"math"
)

// TODO: combine functions in this file for better performance
// TODO: Add a struct for a more unified histogram creation api
// TODO: Redo functions, ad comments, cuz the current state is bug prone and written to complicated

type Bin struct {
	values              []float64
	X0, X1              float64
	NumberOfPointsInBin int
}

func CalculateNumberOfBins(dataPointAmount int) int {
	// calculated by using Sturges rule:
	//	k = 1 + log2(n);
	// 	TODO: Research other rules as well
	k := 1 + math.Log2(float64(dataPointAmount))

	return int(math.Ceil(k))
}

func CalculateBinEdges(binAmount int, min, max float64) []float64 {
	binWidth := (max - min) / float64(binAmount)

	binEdges := make([]float64, binAmount+1)
	for i := 0; i <= binAmount; i++ {
		binEdges[i] = min + binWidth*float64(i)
	}

	return binEdges
}

func Digitize(data, binEdges []float64) []int {
	// TODO: look into a faster solution, but this is at the moment fast enough,
	// since the simulation logic is bottlenecked by the speed of go.
	// So the little time gotten from improving this is marginal

	bins := make([]int, len(data))
	for i, value := range data {
		for j := 1; j < len(binEdges); j++ {
			if binEdges[j-1] <= value && value <= binEdges[j] {
				bins[i] = j - 1
				break
			}
		}
	}

	return bins
}

func CreateBins(data []float64, binIndexes []int, binEdges []float64, binAmount int) []Bin {
	// seams buggy
	bins := make([]Bin, binAmount)
	for i := 0; i < len(binIndexes); i++ {
		idx := binIndexes[i]
		bins[idx].values = append(bins[idx].values, data[i])
		bins[idx].NumberOfPointsInBin += 1
	}

	for i := 0; i < len(bins); i++ {
		bins[i].X0 = binEdges[i]
		bins[i].X1 = binEdges[i+1]
	}

	return bins
}
