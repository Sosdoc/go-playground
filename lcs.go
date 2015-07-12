//Find the longest common substring between two strings
package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "math"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ";")
		lcs := lcs(arr[0], arr[1], len(arr[0]), len(arr[1]))
		fmt.Println(lcs)
	}
}

func lcs(a string, b string, alen int, blen int) string {

	lcs := make([][]int, len(a)+1)
	for x := range lcs {
		lcs[x] = make([]int, len(b)+1)
	}

	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if a[i-1] == b[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = int(math.Max(float64(lcs[i-1][j]), float64(lcs[i][j-1])))
			}
		}
	}

	lcs_i := lcs[alen][blen]
	substring := make([]rune, lcs_i)
	i, j := alen, blen

	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			// append character and keep backtracking
			substring[lcs_i-1] = rune(a[i-1])
			i--
			j--
			lcs_i--
		} else if lcs[i-1][j] > lcs[i][j-1] {
			// pick longest and backtrack in that direction
			i--
		} else {
			j--
		}
	}

	return string(substring)
}

// recursive implementation
// uses memoization

func longest(a string, b string) string {
	if len(a) > len(b) {
		return a
	} else {
		return b
	}
}

func findLCS(m map[string]string, a string, b string, i int, j int) string {
	var key string = fmt.Sprintf("%d_%d", i, j)
	_, present := m[key]
	if present {
		fmt.Println(m[key])
		return m[key]
	} else if i < 0 || j < 0 {
		m[key] = ""
	} else if a[i] == b[i] {
		m[key] = findLCS(m, a, b, i-1, j-1) + string(a[i])
	} else {
		m[key] = longest(findLCS(m, a, b, i-1, j), findLCS(m, a, b, i, j-1))
	}

	if i >= 0 && j >= 0 {
		fmt.Printf("%v and %v\n", string(a[i]), string(b[j]))
	}

	return m[key]
}
