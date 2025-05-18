package main

import "fmt"

var option string

type cryptoList struct {
	name      string
	harga     int
	kesulitan int
	reward    int
}
type arrCrypto [999]cryptoList

func main() {
	var n int = 0
	var listCrypto arrCrypto
	
	menu(&n, &listCrypto)
}

func menu(n *int, A *arrCrypto) {
	var idx int
	var start bool = true

	for start {
		fmt.Println("--------------------------MENU--------------------------")
		fmt.Println("1. Add")
		fmt.Println("2. Change")
		fmt.Println("3. Delete")
		fmt.Println("4. Search")
		fmt.Println("5. Sort")
		fmt.Println("6. Exit")
		fmt.Println("Input the number or action you want to perform: ")
		fmt.Scan(&option)
		if option == "add" || option == "Add" || option == "1" {
			Addition(A, n)
		} else if option == "change" || option == "Change" || option == "2" {
			change(A, *n)
		} else if option == "delete" || option == "Delete" || option == "3" {
			delete(A, n)
		} else if option == "search" || option == "Search" || option == "4" {
			fmt.Println("Search by:")
			fmt.Println("1. Binary")
			fmt.Println("2. Sequential")
			fmt.Println("Input the number or action you want to perform: ")
			fmt.Scan(&option)
			if option == "1" || option == "Binary" || option == "binary" {
				idx = searchBin(*A, *n)
				fmt.Println(idx)
				if idx != 1 {
					fmt.Println("Data ditemukan:")
					fmt.Println(A[idx].name, A[idx].harga)
				} else {
					fmt.Println("Data not found")
				}
			} else if option == "2" || option == "Sequential" || option == "sequential" {
				searchSeq(*A, *n)
			}
		} else if option == "sort" || option == "Sort" || option == "5" {
			fmt.Println("Sorting by:")
			fmt.Println("1. Kesulitan")
			fmt.Println("2. Reward")
			fmt.Println("Input the number or action you want to perform: ")
			fmt.Scan(&option)
			if option == "1" || option == "Kesulitan" || option == "kesulitan" {
				fmt.Println("Select sorting type:")
				fmt.Println("1. Selection")
				fmt.Println("2. Insertion ")
				fmt.Println("Input the number or action you want to perform: ")
				if option == "1" || option == "Selection" || option == "selection" {
					selSortKesulitan(A, *n)
				} else if option == "2" || option == "Insertion" || option == "insertion" {
					inSortKesulitan(A, *n)
				}
			} else if option == "2" || option == "Reward" || option == "reward" {
				fmt.Println("Select sorting type:")
				fmt.Println("1. Selection")
				fmt.Println("2. Insertion ")
				fmt.Println("Input the number or action you want to perform: ")
				if option == "1" || option == "Selection" || option == "selection" {
					selSortReward(A, *n)
				} else if option == "2" || option == "Insertion" || option == "insertion" {
					inSortReward(A, *n)
				}
			}
		} else if option == "exit" || option == "Exit" || option == "6" {
			start = false
		} else {
			fmt.Println("Input invalid")
		}
		fmt.Println("--------------------------------------------------------")
	}
}

func Addition(A *arrCrypto, n *int) {
	var i int

	fmt.Println("Please input the number of arrays you want to add: ")
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].name, &A[i].harga, &A[i].kesulitan, &A[i].reward)
	}
}

func change(A *arrCrypto, n int) {
	var idx = -1

	idx = searchBin(*A, n)
	if idx != -1 {
		fmt.Scan(&A[idx].name, &A[idx].harga, &A[idx].kesulitan, &A[idx].reward)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func searchSeq(A arrCrypto, n int) int {
	var s int
	var x string
	var searching int = -1

	fmt.Println("Input the name of the cryptocurrency to search: ")
	fmt.Scan(&x)

	for searching == -1 && s < n {
		if A[s].name == x {
			searching = s
			fmt.Println("Data ditemukan:")
			fmt.Println(A[s].name, A[s].harga, A[s].kesulitan, A[s].reward)
		}
		s++
	}
	return searching
}

func searchBin(A arrCrypto, n int) int {
	var left, right, mid int
	var idx int
	var x string

	fmt.Scan(&x)
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x == A[mid].name {
			idx = mid
		} else if x < A[mid].name {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}

func delete(A *arrCrypto, n *int) {
	var i, j, found int
	var x string

	fmt.Scan(&x)
	found = -1
	i = 0
	for i < *n && found == -1 {
		if A[i].name == x {
			found = i
		}
		i++
	}

	if found == -1 {
		fmt.Println("Data not found")
	} else {
		for j = found; j < *n-1; j++ {
			A[j] = A[j+1]
		}
		*n--
	}
}

func selSortKesulitan(A *arrCrypto, n int) {
	var pass, idx, i int
	var temp cryptoList

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].kesulitan < A[idx].kesulitan {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
func selSortReward(A *arrCrypto, n int) {
	var pass, idx, i int
	var temp cryptoList

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].reward < A[idx].reward {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func inSortKesulitan(A *arrCrypto, n int) {
	var pass, i int
	var temp cryptoList

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.kesulitan < A[i-1].kesulitan {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func inSortReward(A *arrCrypto, n int) {
	var pass, i int
	var temp cryptoList

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.reward < A[i-1].reward {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}
