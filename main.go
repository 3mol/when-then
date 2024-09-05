package main

import (
	"fmt"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

import (
	"bufio"
	"encoding/json"
	"os"
)

type Job struct {
	When string `json:"when"`
	Then string `json:"then"`
}

func main() {
	file := getFile()
	defer file.Close()

	jobs := parseConfig(file)
	if jobs == nil {
		return
	}

	processInput(jobs)
}

func parseConfig(file *os.File) []Job {
	var jobs []Job
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&jobs); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}
	return jobs
}

func processInput(jobs []Job) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputLine := scanner.Text()
		for _, job := range jobs {
			if job.When == inputLine {
				fmt.Println(job.Then)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}

func getFile() *os.File {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "config.json"
	}
	file, err1 := os.Open(filename)
	if err1 != nil {
		fmt.Println("Error opening file:", err1)
	}
	return file
}
