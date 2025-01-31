package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func read(file *os.File) {
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading record", err)
		}

		fmt.Println(record)
	}
}

func main() {
	firstTasks := [][]string{
		{"task_number", "task_name"},
		{"1", "walk dog"},
	}

	tasks, err := os.Open("tasks.csv")
	if err != nil {
		read(tasks)
	}

	w := csv.NewWriter(os.Stdout)

	for _, task := range firstTasks {
		if err := w.Write(task); err != nil {
			log.Fatalln("error writing tasks")
		}
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
