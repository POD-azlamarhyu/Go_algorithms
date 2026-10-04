package src

import (
	"log/slog"
	"sort"
)

func arrayPermutation(array []int) (*[][]int){
	var helper func([]int, int)
	var rdata [][]int

	helper = func(nums []int, n int){
		if n == len(nums){
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			rdata = append(rdata, tmp)
			return 
		}
		for i := n; i < len(nums); i++{
			nums[n], nums[i] = nums[i], nums[n]
			helper(nums, n+1)
			nums[n], nums[i] = nums[i], nums[n]
		}
	}

	helper(array, 0)
	return &rdata
}

func nextPermutationArray[T Data](array []T) ([]T, bool){
	n := len(array)
	a := array
	i := n - 2
	for i >= 0 && a[i] >= a[i+1]{
		i--
	}

	if i < 0 {
		return nil, false
	}

	j := n - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]

	for l, r := i+1 , n-1; l < r; l, r= l+1, r-1{
		a[l], a[r] = a[r], a[l]
	}
	return a, true
}

func nextPermutations(array []int) (bool){
	n := len(array)
	a := array

	i := n - 2
	for i >= 0 && a[i] >= a[i+1]{
		i--
	}

	if i < 0 {
		return false
	}

	j := n - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]

	for l, r := i+1, n-1; l < r; l, r= l+1, r-1{
		a[l], a[r] = a[r], a[l]
	}
	return true
}

func permutations(array []int) [][]int {
	sort.Ints(array)

	var rdata [][]int

	for {
		p := make([]int, len(array))
		copy(p, array)
		rdata = append(rdata, p)

		if !nextPermutations(array){
			break
		}
	}

	return rdata
}

func addBoth(a []int) []int{
	rdata := make([]int, len(a) + 2)
	rdata[0] = 0
	copy(rdata[1:], a)
	rdata[len(rdata) - 1] = 0

	return rdata
}

func PermutationSearch(data [][]int, Target int) (*int, *[]int){
	n := len(data)
	slog.Info("Start Permutaion Search", slog.Any("data legnth", n))

	cities := make([]int, n-1)

	for i := 1; i < n; i++ {
		cities[i-1]=i
	}
	minDist := int(^uint(0) >> 1)

	var route []int
	var total int
	var shortRoute []int

	perm := permutations(cities)
	slog.Info("Permutations array", slog.Any("data", perm))

	for i := 0; i < len(perm); i++{
		route = addBoth(perm[i])
		total = 0

		for j := 0; j < len(route) - 1; j++{
			total += data[route[j]][route[j+1]]
		}
		
		if total < minDist {
			minDist = total
			shortRoute = route
		}
	}
	return &minDist,&shortRoute
}