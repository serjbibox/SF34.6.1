package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Buffer struct {
	content []byte
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.content = append(b.content, p...)
	return len(p), nil
}
func (b *Buffer) String() string {
	return string(b.content)
}

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
	digitExpr := regexp.MustCompile(`[-+]?(\d+[.]\d+)|\d+`)
	mathExpr := regexp.MustCompile(`[-+*\/]{1}`)
	replaceExpr := regexp.MustCompile(`\?`)

	buf := &Buffer{}
	writer := bufio.NewWriter(buf)

	for _, s := range lines {
		mathExp := expr.FindAllString(s, -1)
		if mathExp == nil {
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
			_, _ = writer.Write([]byte(replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1+n2))))
		case "-":
			_, _ = writer.Write([]byte(replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1-n2))))
		case "*":
			_, _ = writer.Write([]byte(replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1*n2))))
		case "/":
			if n2 == 0 {
				log.Printf("делитель не должен равняться нулю!")
				continue
			}
			_, _ = writer.Write([]byte(replaceExpr.ReplaceAllString(s, fmt.Sprintf("%v\n", n1/n2))))
		default:
			continue
		}
		writer.Flush()
	}

	log.Println(buf)

	err = ioutil.WriteFile("./"+fout, buf.content, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
