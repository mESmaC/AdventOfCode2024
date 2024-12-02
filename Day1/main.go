package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sterilize()
}

/*
* Variable count to count from 0-5
* Bool to switch arrays
* iterate through the sterilized text with a for loop, every five characters switch to a new array
* print arrays to their own files
 */

func sterilize() {
	file, err := os.Open("./list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	count := 0
	swap := false

	hold := ""

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		char := []rune(scanner.Text())[0]

		// fmt.Printf("Char: %q ASCII: %d\n", scanner.Text(), char)

		// fmt.Printf("Count: %d Swap: %v Hold: %q\n", count, swap, hold)

		if unicode.IsSpace(char) || char == '\r' {
			if hold != "" { // Finalize and append the current number
				val, _ := strconv.Atoi(hold)
				if swap == false {
					list1 = append(list1, val)
				} else {
					list2 = append(list2, val)
				}
				hold = ""    // Reset `hold`
				count = 0    // Reset `count`
				swap = !swap // Toggle `swap`
			}
			continue
		}

		if count < 5 {
			hold += scanner.Text()
			count++
		} else if count == 5 && swap == false {
			val, _ := strconv.Atoi(hold)
			list1 = append(list1, val)
			hold = ""
			count = 0
			swap = !swap
		} else if count == 5 && swap == true {
			val, _ := strconv.Atoi(hold)
			list2 = append(list2, val)
			hold = ""
			count = 0
			swap = !swap
		}
	}

	if hold != "" {
		val, _ := strconv.Atoi(hold)
		if swap == false {
			list1 = append(list1, val)
		} else {
			list2 = append(list2, val)
		}
	}

	//fmt.Println("List1: ", list1)
	//fmt.Println("List2: ", list2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sortAndCalc(list1, list2)
}

func sort(s []int) []int {
	l := 0
	r := len(s) - 1
	help(s, l, r)
	return s
}

func help(s []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2

		help(s, l, m)
		help(s, m+1, r)

		merge(s, l, m, r)
	}
}

func merge(s []int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m

	sl := make([]int, n1)
	sr := make([]int, n2)

	for i := 0; i < n1; i++ {
		sl[i] = s[l+i]
	}
	for j := 0; j < n2; j++ {
		sr[j] = s[m+1+j]
	}

	i := 0
	j := 0
	k := l

	for i < n1 && j < n2 {
		if sl[i] <= sr[j] {
			s[k] = sl[i]
			i++
		} else {
			s[k] = sr[j]
			j++
		}
		k++
	}

	for i < n1 {
		s[k] = sl[i]
		i++
		k++
	}

	for j < n2 {
		s[k] = sr[j]
		j++
		k++
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortAndCalc(l1 []int, l2 []int) {
	sl1 := sort(l1)
	sl2 := sort(l2)

	fmt.Printf("Length of sl1: %d, Length of sl2: %d\n", len(sl1), len(sl2))

	sod := 0 // SOD = Sum of differences

	//fmt.Println("Sorted l1: ", sl1)
	//fmt.Println("Sorted l2: ", sl2)

	for i := 0; i < len(sl1); i++ {
		sod += abs(sl1[i] - sl2[i])
	}

	fmt.Println(sod)
}
