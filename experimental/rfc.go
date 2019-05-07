package main

import (
	"fmt"
)

// Sign returns the sign of an integer.
// func sign(n int) int {
// 	switch {
// 	case n > 0:
// 		return 1
// 	case n < 0:
// 		return -1
// 	}
// 	return 0
// }

// Abs returns the absolute value of an integer.
func abs(n int) int {
	// modulo := func(x, n int) int {
	// 	return (x%n + n) % n
	// }
	// m := modulo(n, n*n-n+2)
	// p := modulo(m, n*n+2)
	// return 2*p - n
	if n < 0 {
		return -n
	}
	return n
}

// func ms2() {
// 	start := 3
// 	stop := 8
// 	delta := stop - start
// 	period := 2 * abs(delta)
// 	for i := 0; i < 20; i++ {
// 		v := start + abs(delta-(i+delta)%period)
// 		// v = stop - abs(delta-i%period)
// 		fmt.Print(v, " ")
// 	}
// 	fmt.Println()
// }

func mirrorSequenceUnit(start, pivot int, count int) []int {
	// r is the range, but we can't use that word
	r := abs(pivot - start)
	period := 2 * r

	nn := make([]int, count)
	for i := range nn {
		if pivot == start {
			nn[i] = pivot
		} else {
			if pivot > start {
				nn[i] = pivot - abs(r-i%period)
				// nn[i] = start + abs(r-(i+r)%period)
			} else if pivot < start {
				nn[i] = pivot + abs(r-i%period)
				// nn[i] = start - abs(r-(i+r)%period)
			}
		}
	}
	return nn
}

func mirrorSequenceGen(start, pivot int) func() int {
	// r is the range, but we can't use that word
	r := abs(pivot - start)
	period := 2 * r

	var i int

	return func() int {
		n := i
		i++

		if pivot > start {
			return pivot - abs(r-n%period)
			// return start + abs(r-(n+r)%period)
		} else if pivot < start {
			return pivot + abs(r-n%period)
			// return start - abs(r-(n+r)%period)
		}

		return pivot
	}
}

// func crescendoPyramidalSubsequence(min, max int) []int {
// 	fmt.Println("diffs := ", deltaSeqDiffs(min, max, max-min))
// 	// nn := make([]int, 2*n-1)
// 	nn := []int{}
// 	// midpoint := 2*n + 1
// 	// fmt.Println(midpoint)
// 	for _, d := range deltaSeqDiffs(min, max, 2*(max-min)+1) {
// 		nn = append(nn, max-d)
// 	}
// 	// for i := 0; i < n; i++ {
// 	// 	nn[i] = i + 1
// 	// 	nn[n+i-1] = n - i
// 	// }
// 	return nn
// }

// func decrescendoPyramidalSubsequence(min, max int) []int {
// 	nn := []int{}
// 	for _, d := range deltaSeqDiffs(min, max, 2*(max-min)+1) {
// 		nn = append(nn, min+d)
// 	}
// 	// nn := make([]int, 2*n-1)
// 	// for i := 0; i < n; i++ {
// 	// 	nn[i] = n - i
// 	// 	nn[n+i-1] = i + 1
// 	// }
// 	return nn
// }

func rfcEncode(s string, numRails int) string {
	rails := [][]int{}
	for i := 0; i < numRails; i++ {
		rail := []int{}
		rails = append(rails, rail)
	}

	// fmt.Println("hey:", crescendoPyramidalSubsequence(2, 6))
	// fmt.Println("hey:", decrescendoPyramidalSubsequence(2, 6))

	// seqLen := 2 * (numRails - 1)
	// for i := 1; i <= seqLen; i++ {
	// 	x := seqLen - i%numRails

	// 	fmt.Println(x)
	// }

	// chars := []rune(s)
	// ctr := 0
	// for ctr < len(chars) {
	// 	rails[]
	// }

	return s
}

// f(3)=4 1232 12321
// f(4)=6 123432
// f(5)=8 12345432

// f(2)=4
// f(3)=6
// f(4)=8

// 1..3
// -1010-

func rfcDecode(s string, rails int) string {
	return s
}

func main() {
	// ms2()
	f := mirrorSequenceGen(1, 5)
	fmt.Println("SEQ:", f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f())

	x := mirrorSequenceUnit(1, 5, 15)
	y := mirrorSequenceUnit(5, 1, 15)
	z := mirrorSequenceUnit(5, 5, 15)
	fmt.Println("SEQ:", x)
	fmt.Println("SEQ:", y)
	fmt.Println("SEQ:", z)
	// rails := 3
	// p := deltaSeqDiffFunc(3, 8)
	// mytestmirror()
	//fmt.Println("msu := ", mirrorSequenceUnit(-1, -7))
	// fmt.Println("hey!!!", deltaSeqDiffs(2, 6, 100))
	// fmt.Println("hey!!!", abs(0))
	// fmt.Println("hey!!!", abs(1), abs(2), abs(3), abs(5), abs(8), abs(13), abs(21))
	// fmt.Println("hey!!!", abs(-1), abs(-2), abs(-3), abs(-5), abs(-8), abs(-13), abs(-21))
	// msg := "WEAREDISCOVEREDFLEEATONCE"
	// e := rfcEncode(msg, rails)
	// d := rfcDecode(e, rails)
	// fmt.Println("m =", msg)
	// fmt.Println("e =", e)
	// fmt.Println("d =", d)
}
