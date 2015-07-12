package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "sort"

func getPermutations(input []string) []string {
	if len(input) == 1 {
		return []string{input[0]}
	} else {
		var res []string
		for _, combination := range getPermutations(input[1:]) {
			for i, _ := range input {
				res = append(res, combination[:i]+input[0]+combination[i:])
			}
		}
		return res
	}
}

func deDuplicate(input []string) []string {
	var result []string
	found := make(map[string]bool)

	for _, val := range input {
		if _, ok := found[val]; !ok {
			result = append(result, val)
			found[val] = true
		}
	}
	return result
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), "")
		p := sort.StringSlice(getPermutations(inputs))
		p = deDuplicate(p)
		sort.Sort(p)

		for i := 0; i < len(p); i++ {
			fmt.Print(p[i])
			if i != len(p)-1 {
				fmt.Print(",")
			}
		}

		fmt.Println("")

	}
}
