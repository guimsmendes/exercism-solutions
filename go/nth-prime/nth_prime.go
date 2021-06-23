package prime

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	var count, prime int
	for i := 2; ; i++ {
		if isPrime(i) {
			prime = i
			count++
		}
		if count == n {
			break
		}
	}
	return prime, true
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
