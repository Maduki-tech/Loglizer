package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Logs struct {
	errors   int
	warnings int
	infos    int
}

func main() {
	fileArg := os.Args[1]
	file, err := os.OpenFile(fileArg, os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line)
	}

	generateBarChart(logs.infos, logs.errors, logs.warnings)
}

var logs Logs

func parseLine(line string) {
	logType := strings.Split(line, ":")[0]
	switch {
	case logType == "INFO":
		parseForInfo(line)
	case logType == "WARN":
		parseForWarn(line)
	case logType == "ERROR":
		parseForErrors(line)
	default:
		fmt.Println("UNKNOWN")
	}
}

func parseForInfo(line string) {
	fmt.Println(line)
	logs.infos++
}

func parseForWarn(line string) {
	fmt.Println(line)
	logs.warnings++
}

func parseForErrors(line string) {
	fmt.Println(line)
	logs.errors++
}
