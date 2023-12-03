package main

import (
	"sort"
	"strconv"
)

type IntBitset struct {
	data []uint64
}

func NewIntBitset() *IntBitset {
	return &IntBitset{
		data: []uint64{},
	}
}

func (ibs *IntBitset) Contains(d uint64) bool {
	for _, v := range ibs.data {
		if v == d {
			return true
		}
	}
	return false
}

func (ibs *IntBitset) ToString() string {
	var res string
	res += "[ "
	for _, v := range ibs.data {
		res += strconv.Itoa(int(v))
		res += " "
	}
	res += "]"
	return res
}

func (ibs *IntBitset) Add(d uint64) {
	ibs.data = append(ibs.data, d)
}

func (ibs *IntBitset) AddRange(d ...uint64) {
	ibs.data = append(ibs.data, d...)
}

func (ibs *IntBitset) AddSlice(d []uint64) {
	ibs.data = append(ibs.data, d...)
}

func (ibs *IntBitset) Reverse() {
	for i, j := 0, len(ibs.data)-1; i < j; i, j = i+1, j-1 {
		ibs.data[i], ibs.data[j] = ibs.data[j], ibs.data[i]
	}
}

func (ibs *IntBitset) SortAsc() {
	sort.Slice(ibs.data, func(i, j int) bool {
		return ibs.data[i] < ibs.data[j]
	})
}

func (ibs *IntBitset) SortDsc() {
	sort.Slice(ibs.data, func(i, j int) bool {
		return ibs.data[i] > ibs.data[j]
	})
}

func (ibs *IntBitset) DeleteAll(d uint64) {
	changed := true
	for changed {
		changed = false
		for i, v := range ibs.data {
			if v == d {
				temp := ibs.data[:i]
				temp2 := ibs.data[i+1:]
				temp = append(temp, temp2...)
				ibs.data = temp
				changed = true
				break
			}
		}
	}
}

func (ibs *IntBitset) DeleteFirst(d uint64) {
	for i, v := range ibs.data {
		if v == d {
			temp := ibs.data[:i]
			temp2 := ibs.data[i+1:]
			temp = append(temp, temp2...)
			ibs.data = temp
			return
		}
	}
}

// func main() {
// 	ibs := NewIntBitset()
// 	fmt.Println(ibs.ToString())
// 	ibs.Add(3)
// 	fmt.Println(ibs.ToString())
// 	ibs.AddSlice([]uint64{5, 5, 4, 5, 6, 5, 7})
// 	fmt.Println(ibs.ToString())
// 	ibs.AddRange(8, 9, 10, 5, 11, 5, 12, 5, 15)
// 	fmt.Println(ibs.ToString())
// 	ibs.DeleteAll(5)
// 	fmt.Println(ibs.ToString())

// 	ibs2 := NewIntBitset()
// 	ibs2.AddRange(8, 3, 4, 6, 1, 0, 9, 12, 2, 2, 45)
// 	fmt.Println(ibs2.ToString())
// 	ibs2.Reverse()
// 	fmt.Println(ibs2.ToString())
// 	ibs2.SortAsc()
// 	fmt.Println(ibs2.ToString())
// 	ibs2.SortDsc()
// 	fmt.Println(ibs2.ToString())
// 	fmt.Println(ibs2.Contains(17))
// 	fmt.Println(ibs2.Contains(4))

// }
