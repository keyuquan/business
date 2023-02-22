package utils

import (
	"math/rand"
	"sort"
	"time"
)

type Replacer interface {
	Len() int
	Replace([]int) Replacer
	Sort()
	SortKey() string
	Self() []string
}

type StringReplacer []string

// Len implements Replacer.Len
func (sl StringReplacer) Len() int {
	return len(sl)
}

// Replace implements Replacer.Replace
func (sl StringReplacer) Replace(indices []int) Replacer {
	result := make(StringReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = sl[idx]
	}
	return result
}

func (sl StringReplacer) Self() []string {
	return sl
}

func (sl StringReplacer) Sort() {
	sort.Strings(sl)
}

func (sl StringReplacer) SortKey() (res string) {
	sl.Sort()
	for index, v := range sl {
		if index == 0 {
			res += v
		} else {
			res += "," + v
		}

	}
	return
}

// Permutations is generator
func Permutations(list Replacer, selectNum int, repeatable bool, buf int) (c chan Replacer) {
	c = make(chan Replacer, buf)
	go func() {
		defer close(c)
		var permGenerator func([]int, int, int) chan []int
		if repeatable {
			permGenerator = repeatedPermutations
		} else {
			permGenerator = permutations
		}
		indices := make([]int, list.Len(), list.Len())
		for i := 0; i < list.Len(); i++ {
			indices[i] = i
		}
		for perm := range permGenerator(indices, selectNum, buf) {
			c <- list.Replace(perm)
		}
	}()
	return
}

func pop(l []int, i int) (v int, sl []int) {
	v = l[i]
	length := len(l)
	sl = make([]int, length-1, length-1)
	copy(sl, l[:i])
	copy(sl[i:], l[i+1:])
	return
}

//Permtation generator for int slice
func permutations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch selectNum {
		case 1:
			for _, v := range list {
				c <- []int{v}
			}
			return
		case 0:
			return
		case len(list):
			for i := 0; i < len(list); i++ {
				top, subList := pop(list, i)
				for perm := range permutations(subList, selectNum-1, buf) {
					c <- append([]int{top}, perm...)
				}
			}
		default:
			for comb := range combinations(list, selectNum, buf) {
				for perm := range permutations(comb, selectNum, buf) {
					c <- perm
				}
			}
		}
	}()
	return
}

//Repeated permtation generator for int slice
func repeatedPermutations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch selectNum {
		case 1:
			for _, v := range list {
				c <- []int{v}
			}
		default:
			for i := 0; i < len(list); i++ {
				for perm := range repeatedPermutations(list, selectNum-1, buf) {
					c <- append([]int{list[i]}, perm...)
				}
			}
		}
	}()
	return
}

func Combinations(list Replacer, selectNum int, repeatable bool, buf int) (c chan Replacer) {
	c = make(chan Replacer, buf)
	index := make([]int, list.Len(), list.Len())
	for i := 0; i < list.Len(); i++ {
		index[i] = i
	}

	var combGenerator func([]int, int, int) chan []int
	if repeatable {
		combGenerator = repeatedCombinations
	} else {
		combGenerator = combinations
	}

	go func() {
		defer close(c)
		for comb := range combGenerator(index, selectNum, buf) {
			c <- list.Replace(comb)
		}
	}()

	return
}

// combinations generator is for int slice
func combinations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch {
		case selectNum == 0:
			c <- []int{}
		case selectNum == len(list):
			c <- list
		case len(list) < selectNum:
			return
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range combinations(list[i+1:], selectNum-1, buf) {
					c <- append([]int{list[i]}, subComb...)
				}
			}
		}
	}()
	return
}

// repeatedCombination is generator for int slice
func repeatedCombinations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		if selectNum == 1 {
			for v := range list {
				c <- []int{v}
			}
			return
		}
		for i := 0; i < len(list); i++ {
			for subComb := range repeatedCombinations(list[i:], selectNum-1, buf) {
				c <- append([]int{list[i]}, subComb...)
			}
		}
	}()
	return
}

// Shuffle 随机数组位置
func Shuffle(slice []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
