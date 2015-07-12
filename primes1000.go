// compute the sum of the first 1000 prime numbers
package main

import "fmt"
import "math"

func main() {
	//sieve(100000)
	primes := sieveParallel(10)
	for i := 0; i < len(primes); i++ {
		fmt.Print(primes[i])
		if i != len(primes)-1 {
			fmt.Print(",")
		}
	}

}

func markNumbers(n int, array []bool, end chan bool) {
	for i := n * 2; i < len(array); i += n {
		array[i] = false
	}
	end <- true
}

func sieveParallel(limit int) []int {

	prime := make([]bool, limit)
	countChan := make(chan bool)

	for i := 0; i < limit-1; i++ {
		prime[i] = true
	}

	end := int(math.Sqrt(float64(limit)))

	for i := 2; i < end; i++ {
		if prime[i] {
			go markNumbers(i, prime, countChan)
		}
	}

	for i := 2; i < end; i++ {
		_ = <-countChan
	}

	result := make([]int, 0, limit)

	for i := 2; i < limit; i++ {
		if prime[i] {
			result = append(result, i)
		}
	}
	return result
}

func sieve(limit int) (result []int) {
	prime := make([]bool, limit)

	for i := 0; i < limit-1; i++ {
		prime[i] = true
	}

	for i := 2; i < limit; i++ {
		if prime[i] {
			for j := i + i; j < limit; j += i {
				prime[j] = false
			}
		}
	}

	result = make([]int, 0, limit)

	for i := 2; i < limit; i++ {
		if prime[i] {
			result = append(result, i)
		}
	}
	return
}

func sieveCalc(limit int) (result []int) {
	sieve := make([]int, limit)

	for i := 0; i < limit-1; i++ {
		sieve[i] = i + 2
	}

	for i := 0; i < limit; i++ {
		divisor := sieve[i]
		if sieve[i] > 0 {
			for j := i + 1; j < limit; j++ {
				if sieve[j]%divisor == 0 {
					sieve[j] = -1
				}
			}
		}
	}

	result = make([]int, 0, limit)
	//result = append(result, 1)

	for i := 0; i < limit; i++ {
		if sieve[i] > 0 {
			result = append(result, sieve[i])
		}
	}
	return
}

func sieveSum(count int, limit int) (result []int, sum int) {
	sieve := make([]int, limit)
	computed := 0

	for i := 0; i < limit-1; i++ {
		sieve[i] = i + 2
	}

	i := 0
	for computed < count {
		divisor := sieve[i]
		computed++
		if sieve[i] > 0 {
			for j := i + 1; j < limit; j++ {
				if sieve[j]%divisor == 0 {
					sieve[j] = -1
				}
			}
		}
		i++
	}

	result = make([]int, 0, limit)
	result = append(result, 1)
	sum = 0

	for i := 0; i < limit; i++ {
		if sieve[i] > 0 {
			result = append(result, sieve[i])
			sum += sieve[i]

			if len(result) == count+1 {
				break
			}
		}
	}
	return
}
