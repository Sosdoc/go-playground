//Sample code to read in test cases:
package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"
import "math"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		primes := sieve(int(n + 2))

		for i := 0; i < len(primes); i++ {
			if i < len(primes)-1 {
				fmt.Printf("%d,", primes[i])
			} else {
				fmt.Println(primes[i])
			}
		}
	}
}

func sieve(limit int) (result []int) {
	prime := make([]bool, limit)

	for i := 0; i < limit-1; i++ {
		prime[i] = true
	}

	for i := 2; i < limit; i++ {
		if prime[i] {
			for j := i * i; j < limit; j += i {
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
