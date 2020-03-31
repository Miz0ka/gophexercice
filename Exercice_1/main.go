package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var goodRep int
	var failRep int
	var csvFilePath string
	questions := make(map[string]string)

	flag.StringVar(&csvFilePath, "csv", ".\\probleme.csv", "Link to the csv file")
	timer := flag.Int("time", 42, "Timer")
	flag.Parse()

	if _, err := os.Stat(csvFilePath); err != nil {
		log.Fatal(err)
	}

	csvFile, _ := os.Open(csvFilePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		questions[line[0]] = line[1]
		fmt.Printf("Q: %s A: %s\n", line[0], line[1])
	}
	/*records, err := csvFile.ReadAll()
	if err != nil {
		log.Fatal(err)
	}*/

	fmt.Print("Initialisation\n")

	quizz(questions, &goodRep, &failRep, *timer)

	fmt.Print("\n", goodRep, failRep, timer)
}

func quizz(quizz map[string]string, goodRep *int, badRep *int, timer int) {
	var input string
	for q, a := range quizz {
		fmt.Println(q + " : ")
		fmt.Scanf("%s", &input)
		if input == a {
			*goodRep++
		} else {
			*badRep++
		}
	}
}
