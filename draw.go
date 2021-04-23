package draw

import "math/rand"

// Pool for drawcard
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Weight returns weight of the elements at index i.
	Weight(i int) int
}

// totalWeight returns total weight of the elements
func totalWeight(data Interface) int {
	w := 0
	for i := 0; i < data.Len(); i++ {
		weight := data.Weight(i)
		if weight <= 0 {
			panic("draw:totalWeight weight value failed, it's must 0")
		}
		w += data.Weight(i)
	}
	return w
}

func slice(length int) []int {
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = i
	}
	return slice
}

// randRange returns as an int, a non-negative pseudo-random number in [min,max]
func randRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Rand returns the index of elements
// it return all if count >= data.length and !repeated
func Rand(data Interface, count int, repeated bool) []int {
	if count == 0 || data.Len() == 0 {
		return nil
	}

	result := make([]int, 0, count)
	if !repeated && count >= data.Len() {
		return slice(data.Len())
	}

	var (
		totalWeight = totalWeight(data)
		curCount    = 0
		index       = slice(data.Len())
	)

	for curCount < count {
		weight := randRange(1, totalWeight)
		for i, idx := range index {
			weight -= data.Weight(idx)
			if weight <= 0 {
				result = append(result, idx)
				curCount += 1
				if !repeated {
					index = append(index[:i], index[i+1:]...)
				}
				break
			}
		}

	}

	return result
}
