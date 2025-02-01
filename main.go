package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readAll(filename string) ([]map[string]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		log.Fatal("CSV file is empty of only contains headers")
	}

	headers := rows[0]
	var records []map[string]string

	for _, row := range rows[1:] {
		record := make(map[string]string)
		for i, value := range row {
			record[headers[i]] = value
		}
		records = append(records, record)
	}
	return records, nil
}

func writeNewTask(fileName string, newTask string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	if len(rows) == 0 {
		log.Fatal("CSV file is empty")
	}

	var lastID int
	if len(rows) > 1 {
		lastRow := rows[len(rows)-1]
		lastID, err = strconv.Atoi(lastRow[0])
		if err != nil {
			log.Fatal("Error converting last ID:", err)
		}
	} else {
		lastID = 0
	}

	newID := lastID + 1
	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening file for writing:", err)
	}
	defer file.Close()

	newRow := []string{strconv.Itoa(newID), newTask}

	writer := csv.NewWriter(file)
	err = writer.Write(newRow)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
	fmt.Println("New task added with ID:", newID)
}

//	func deleteTask(file *os.File, task int) {
//		file.Seek(0, 0)
//		reader := csv.NewReader(file)
//
//		for {
//			record, err := reader.Read()
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				log.Fatal("Error reading file:", err)
//			}
//			fmt.Printf("%v", record)
//		}
//	}
func main() {
	fmt.Println("Hello, would you like to see tasks and create a new one?")
	fmt.Println("1: read")
	fmt.Println("2: write")
	fmt.Println("3: delete task")

	var option int
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		readAll("tasks.csv")
	case 2:
		fmt.Println("What would you like to write?")
		var task string
		fmt.Scanf("%v", &task)
		writeNewTask("tasks.csv", task)
	// case 3:
	// 	readAll("tasks.csv")
	// 	fmt.Println("Which task do you want deleted?")
	// 	var task int
	// 	fmt.Scanf("%d", task)
	// 	deleteTask("tasks.csv", task)
	default:
		fmt.Println("Invalid Option")
	}
}
