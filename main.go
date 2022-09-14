package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Укажите файл для чтения: ")
	var fin string
	fmt.Scanln(&fin)
	fmt.Print("Укажите файл для записи результата: ")
	var fout string
	fmt.Scanln(&fout)

	content, err := ioutil.ReadFile("./" + fin)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`.*|\n`)
	lines := re.FindAllString(string(content), -1)
	if lines == nil {
		log.Fatal("входной файл пустой")
	}
	for idx, _ := range lines {
		lines[idx] = strings.TrimSpace(lines[idx])
		//log.Println("строки файла:", lines[idx])
	}

	expr := regexp.MustCompile(`^[-+]?((\d+[.]\d+)|\d+)[-+*\/]{1}((\d+[.]\d+)|\d+)[=][?]$`)
	digitExpr := regexp.MustCompile(`(\d+[.]\d+)|\d+`)
	mathExpr := regexp.MustCompile(`[-+*\/]{1}`)
	replaceExpr := regexp.MustCompile(`\?`)
	var linesOut string
	for _, s := range lines {
		exp := expr.FindAllString(s, -1)
		if exp == nil {
			continue
		}
		//log.Printf("мат. выражение строки %d: %s", idx, s)

		n := digitExpr.FindAllString(s, -1)
		n1, err := strconv.ParseFloat(n[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		n2, err := strconv.ParseFloat(n[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		do := mathExpr.FindString(s)

		switch do {
		case "+":
			linesOut += replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1+n2))
		case "-":
			linesOut += replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1-n2))
		case "*":
			linesOut += replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1*n2))
		case "/":
			if n2 == 0 {
				log.Printf("делитель не должен равняться нулю!")
				continue
			}
			linesOut += replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1/n2))
		}
	}

	log.Println(linesOut)
	err = ioutil.WriteFile("./"+fout, []byte(linesOut), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
