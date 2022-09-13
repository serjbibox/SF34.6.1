package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./math.in")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(content))
	re := regexp.MustCompile(`[0-9]`)
	submatches := re.FindAllStringSubmatch(string(content), -1)

	for _, s := range submatches {
		classes := strings.Split(s[1], " ")
		for _, c := range classes {
			fmt.Println("Найден класс", c)
		}
	}
}
