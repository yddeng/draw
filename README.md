# Draw

抽卡、抽奖、掉落库 drawcard 、lottery

## Interface

```
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Weight returns weight of the elements at index i.
	Weight(i int) int
}
```

## Usage

Rand returns the index of elements

`func Rand(data Interface, count int, repeated bool) []int `

RandRangeInt returns as an int, a non-negative pseudo-random number in [min,max]

`func RandRangeInt(min, max int) int`

## Example

```
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
}

```