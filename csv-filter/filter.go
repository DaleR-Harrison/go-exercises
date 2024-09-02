package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"flag"
	"encoding/csv"
	"regexp"
	"bufio"
	"strings"
)

func main() {
	fileName := flag.String("file", "", "name of csv file")
	flag.Parse()
	newFileName := readInFileToEdit(*fileName)
	readInCsvFile(newFileName)
	
}

func readInFileToEdit(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Failed to read from file", err)
	}

	defer file.Close()

	newFileName := "formatted-address-record.csv"
	newFile, err := os.Create(newFileName)

	if err != nil {
		log.Fatal("Failed to read from file", err)
	}

	defer newFile.Close()

	reader := bufio.NewReader(file)

	var re = regexp.MustCompile(`\s*\\"(\B|\b)`)
	
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Failed to read in file", err)
		}
		
		s := re.ReplaceAllString(line, "\"\"")
		fmt.Println(s)
		fmt.Fprintln(newFile, s)
	}

	return newFileName
}

func readInCsvFile (fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Failed to read from file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	newFile, err := os.Create("filtered-" + fileName)

	if err != nil {
		log.Fatal("Failed to read from file", err)
	}

	csvWriter := csv.NewWriter(newFile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if (record[4] != "BCH" && record[4] != "BTC") {
			trimmedString := strings.Trim(strings.Trim(strings.Trim(record[7], "]"), "["), `"`)
			fmt.Println(trimmedString)
			newRecord := []string{record[0], record[4], record[3], trimmedString}
			csvWriter.Write(newRecord)
		}
	
	}
}
