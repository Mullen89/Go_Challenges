package mullen_toolbox

func shiftSliceLeft(shift int, arr []int) []int {
	for i := 0; i < shift; i++ {
		x := arr[0]          // get the 0 index element from slice
		arr = arr[1:]        // remove the 0 index element from slice
		arr = append(arr, x) // print 0 index element
	}
	return arr
}

func bubbleSort(arr []int) {
	totalNumSwaps := 0
	for i := 0; i < len(arr); i++ {
		numSwaps := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numSwaps++
				totalNumSwaps++
			}
		}
		if numSwaps == 0 {
			break
		}
	}
	// fmt.Printf("Array is sorted in %d swaps.\n", totalNumSwaps)
	// fmt.Printf("First Element: %d\n", arr[0])
	// fmt.Printf("Last Element: %d\n", arr[len(arr)-1])
}

func factorial(n int) int {
	if n <= 1 {
		return n
	} else {
		n = n * factorial(n-1)
	}
	return n
}

func randNum(range int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(range))
}

//checks if a number is a power of two.
func PowersOfTwo(i int) bool {
	return i != 0 && (i-1)&i == 0
}