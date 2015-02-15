package problem14

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

// collatzChainLen returns the length of a Collatz chain starting at start.
func collatzChainLen(start int) int {
	next := start
	count := 0
	for {
		if next == 1 {
			return count
		}
		count++
		next = nextCollatz(next)
	}
}

func MaxChain() []int {
	var results = make(map[int][]int)
	max := 0
	for i := 1; i < 1*1000*1000; i++ {
		len := collatzChainLen(i)
		results[len] = append(results[len], i)
	}
	for k, _ := range results {
		if k > max {
			max = k
		}
	}
	return results[max]
}
