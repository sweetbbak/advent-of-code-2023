package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func process(str string, r *regexp.Regexp) (int, error) {
	matches := r.FindAllString(str, -1)
	numerals := strings.Join(matches, "")

	first := numerals[0]
	last := numerals[len(numerals)-1]

	fmt.Printf("First %v Last %v\n", string(first), string(last))

	n := fmt.Sprintf("%v%v", string(first), string(last))
	return strconv.Atoi(n)
}

func addup(n []int) int {
	var ans int
	ans = 0
	for _, num := range n {
		ans = ans + num
	}
	return ans
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile("[0-9]+")

	var numbys []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := scanner.Text()
		num, err := process(a, r)
		if err != nil {
			fmt.Println(err)
		} else {
			numbys = append(numbys, num)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbys)
	answer := addup(numbys)
	fmt.Printf("The answer is: [%v]\n", answer)
}
