package draw

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Slice struct {
	items []*Item
}

type Item struct {
	ID     string
	Weight int
}

func (s *Slice) Len() int {
	return len(s.items)
}

func (s *Slice) Weight(i int) int {
	return s.items[i].Weight
}

func init() {
	rand.Seed(time.Now().Unix())
}

func TestRand(t *testing.T) {
	s := &Slice{items: []*Item{
		{ID: "11", Weight: 10},
		{ID: "22", Weight: 15},
		{ID: "33", Weight: 20},
		{ID: "44", Weight: 25},
		{ID: "55", Weight: 30},
	}}

	// repeated
	result := Rand(s, 3, true)
	fmt.Println(result)

	result = Rand(s, 3, false)
	fmt.Println(result)
}
