package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hungneox/meowify/cats"
)

func handle(fileName string, translate bool, cat cats.CatInterface) {
	file, err := os.Open(fileName) // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if translate {
			fmt.Println()
			fmt.Println(scanner.Text())
		}
		fmt.Println(cat.MeowifyLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// input := "Roses are red, violets are blue, sugar is sweet, and so are you."
	var cat cats.CatInterface

	fileName := flag.String("f", "", "Input text file")
	translate := flag.Bool("t", false, "Translate line by line")
	lang := flag.String("l", "en", "Choose cat language (en|jp)")

	flag.Parse()

	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *lang == "en" {
		cat = cats.Meow{}
	} else {
		cat = cats.Nyan{}
	}

	handle(*fileName, *translate, cat)
}
