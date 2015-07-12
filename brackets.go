//Implement an algorithm for matching braces in a string
package main

import "fmt"
import "log"
import "bufio"
import "os"

type Stack struct {
	data []string
}

func (s *Stack) push(element string) {
	s.data = append(s.data, element)
}

func (s *Stack) pop() (element string) {
	if len(s.data) > 0 {
		element, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	}
	return
}

func (s *Stack) peek() string {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	} else {
		return " "
	}

}

func (s *Stack) size() int {
	return len(s.data)
}

func matchingBraces(input string) bool {
	// get a slice with all the chars in the input string
	chars := input
	s := Stack{}

	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case '(', '[', '{':
			s.push(string(chars[i]))
		case ')', ']', '}':

			if s.size() > 0 {
				starting := s.peek()[0]
				switch

			}

			if s.peek()[0] != '(' {
				return false
			} else {
				_ = s.pop()
			}
		case ']':
			if s.peek()[0] != '[' {
				return false
			} else {
				_ = s.pop()
			}
		case '}':
			if s.peek()[0] != '{' {
				return false
			} else {
				_ = s.pop()
			}
		}
	}

	// string is balanced only if stack is empty (all brackets matched)
	if s.size() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(matchingBraces(input))
	}
}
