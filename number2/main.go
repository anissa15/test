package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input text to get the palindrome")
	fmt.Println("---------------------")

	for {
		fmt.Print("input: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		palindrome := findPalindrome(text)
		fmt.Println("output:", palindrome)

	}
}

func findPalindrome(text string) string {
	spaceIndexs := getSpaceIndex(text)

	a := strings.ReplaceAll(text, " ", "")
	a = strings.ToLower(a)
	b := revert(a)

	maps := getPalindromeMaps(a, b)
	palindrome, index := getFirstLongPalindrome(maps)

	result := addSpace(palindrome, index, a, spaceIndexs)
	return result
}

func getSpaceIndex(text string) []int {
	var indexs []int
	for i := 0; i < len(text); i++ {
		if text[i:i+1] == " " {
			indexs = append(indexs, i)
		}
	}
	return indexs
}

func addSpace(palindrome string, index int, source string, spaceIndexs []int) string {
	for _, v := range spaceIndexs {
		a := source[:v]
		b := source[v:]
		source = a + " " + b

		if index >= v {
			index += 1
			continue
		}

		if v-index >= len(palindrome) {
			break
		}

		a = palindrome[:v-index]
		b = palindrome[v-index:]
		palindrome = a + " " + b
	}
	return palindrome
}

func getFirstLongPalindrome(maps map[string]int) (str string, index int) {
	var length int

	replace := func(k string, v int) {
		str = k
		length = len(k)
		index = v
	}

	for k, v := range maps {
		if len(k) > length {
			replace(k, v)
		} else if len(k) == length {
			if v < index {
				replace(k, v)
			}
		}
	}
	return
}

func getPalindromeMaps(a, b string) map[string]int {
	maps := make(map[string]int)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				p := getPalindrome(a, b, i, j)
				_, exist := maps[p]
				if !exist {
					maps[p] = i
				}
			}
		}
	}
	return maps
}

func getPalindrome(a, b string, i, j int) string {
	var result string
	for {
		if i >= len(a) || j >= len(b) {
			break
		}

		if a[i] == b[j] {
			result += a[i : i+1]
		} else {
			break
		}

		i += 1
		j += 1
	}
	return result
}

func revert(text string) string {
	var new string
	for i := len(text) - 1; i >= 0; i-- {
		new += text[i : i+1]
	}
	return new
}
