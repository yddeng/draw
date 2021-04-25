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
	Weight int // 权重
	Count  int // 产出数量
	Wave   int // 产出数量的波动大小
}

func (s *Slice) Len() int {
	return len(s.items)
}

func (s *Slice) Weight(i int) int {
	return s.items[i].Weight
}

func (s *Slice) Print(ret []int) {
	for _, idx := range ret {
		item := s.items[idx]
		count := RandIntAbs(item.Count, item.Wave)
		fmt.Printf("[ID:%s, Count:%d]", item.ID, count)
	}
	fmt.Println()
}

func init() {
	rand.Seed(time.Now().Unix())
}

func TestDraw(t *testing.T) {
	s := &Slice{items: []*Item{
		{ID: "11", Weight: 10, Wave: 2, Count: 10}, // 0.1
		{ID: "22", Weight: 15, Wave: 5, Count: 6},  // 0.15
		{ID: "33", Weight: 20, Wave: 6, Count: 3},  // 0.2
		{ID: "44", Weight: 25, Wave: 4, Count: 5},  // 0.25
		{ID: "55", Weight: 30, Wave: 0, Count: 3},  // 0.3
	}}

	// 允许重复
	result := Draw(s, 4, true)
	s.Print(result)

	// 不允许重复
	result = Draw(s, 4, false)
	s.Print(result)

}
