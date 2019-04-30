package main

import "fmt"

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

func mirrorSequenceUnit(min, max int) []int {
	delta := max - min

	nn := make([]int, 2*delta)
	for i := 0; i < 2*delta; i++ {
		nn[i] = abs(delta - i)
	}
	return nn
}

func deltaSeqDiffs(min, max, rounds int) []int {
	ms := mirrorSequenceUnit(min, max)

	nn := make([]int, rounds)
	for i := 0; i < rounds; i++ {
		nn[i] = ms[i%len(ms)]
	}
	return nn
}

func deltaSeqDiffFunc(min, max int) func() int {
	ms := mirrorSequenceUnit(min, max)

	i := 0
	return func() int {
		n := ms[i%len(ms)]
		i++
		return n
	}
}

func crescendoPyramidalSubsequence(min, max int) []int {
	fmt.Println("diffs := ", deltaSeqDiffs(min, max, max-min))
	// nn := make([]int, 2*n-1)
	nn := []int{}
	// midpoint := 2*n + 1
	// fmt.Println(midpoint)
	for _, d := range deltaSeqDiffs(min, max, 2*(max-min)+1) {
		nn = append(nn, max-d)
	}
	// for i := 0; i < n; i++ {
	// 	nn[i] = i + 1
	// 	nn[n+i-1] = n - i
	// }
	return nn
}

func decrescendoPyramidalSubsequence(min, max int) []int {
	nn := []int{}
	for _, d := range deltaSeqDiffs(min, max, 2*(max-min)+1) {
		nn = append(nn, min+d)
	}
	// nn := make([]int, 2*n-1)
	// for i := 0; i < n; i++ {
	// 	nn[i] = n - i
	// 	nn[n+i-1] = i + 1
	// }
	return nn
}

func rfcEncode(s string, numRails int) string {
	rails := [][]int{}
	for i := 0; i < numRails; i++ {
		rail := []int{}
		rails = append(rails, rail)
	}

	fmt.Println("hey:", crescendoPyramidalSubsequence(2, 6))
	fmt.Println("hey:", decrescendoPyramidalSubsequence(2, 6))

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
	rails := 3
	p := deltaSeqDiffFunc(3, 8)
	fmt.Println("hey1:", p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p(), p())
	fmt.Println("hey!!!", deltaSeqDiffs(2, 6, 100))
	fmt.Println("hey!!!", abs(0))
	fmt.Println("hey!!!", abs(1), abs(2), abs(3), abs(5), abs(8), abs(13), abs(21))
	fmt.Println("hey!!!", abs(-1), abs(-2), abs(-3), abs(-5), abs(-8), abs(-13), abs(-21))
	msg := "WEAREDISCOVEREDFLEEATONCE"
	e := rfcEncode(msg, rails)
	d := rfcDecode(e, rails)
	fmt.Println("m =", msg)
	fmt.Println("e =", e)
	fmt.Println("d =", d)
}
