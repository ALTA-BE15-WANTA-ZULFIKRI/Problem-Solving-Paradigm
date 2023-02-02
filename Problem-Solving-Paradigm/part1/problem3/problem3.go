package main 

import ( 
"fmt"
"sort"
) 

func DragonOfLoowater(dragonHead, knightHeight []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	
	totalHeight := 0
	i, j := 0, 0
	for i < len(dragonHead) && j < len(knightHeight) {
		if dragonHead[i] <= knightHeight[j] {
			totalHeight += knightHeight[j]
			i++
			j++
		} else {
			j++
		}
	}
	
	if i == len(dragonHead) {
		fmt.Println(totalHeight)
	} else {
		fmt.Println("knight fall")
	}
}

func main() {
	DragonOfLoowater([]int{5,4}, [] int{7, 8, 4})    // 11 
	DragonOfLoowater([]int{5,10},[]int{5})           // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}