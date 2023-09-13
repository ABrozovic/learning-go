package main

import (
	"bufio"
	"fmt"
	exercises "mastering-go-exercises/chapter-3"
	"os"
	"strconv"
	"strings"
	"time"
)

var isValidParam = map[string]bool{
	"insert": true,
	"delete": true,
	"list":   true,
	"search": true,
	"save":   true,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Introduce path to phoneBook.csv")

	args, err := argsHelper(reader)
	csvPath := args[0]

	if err != nil {
		fmt.Println("phoneBook.csv not found. Creating one")

		_, err = os.Create(args[0])

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	lines, err := exercises.ReadCSVFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	mapRecords := map[string]exercises.Record{}

	for _, row := range lines {
		mapRecords[row[2]] = exercises.Record{Name: row[0], Surname: row[1], Tel: row[2], LastAccess: row[3]}
	}

	for {
		fmt.Println("Usage: insert|delete|search|list|save")

		args, _ = argsHelper(reader)

		_, ok := isValidParam[args[0]]
		if !ok {
			fmt.Println("Invalid param", args[0])
			continue
		}

		switch args[0] {
		case "insert":
			fmt.Println("Type the following data: Name Surname Telephone")

			args, _ = argsHelper(reader)
			if len(args) < 3 {
				fmt.Println("Missing arguments")
				continue
			}

			record := exercises.Record{Name: args[0], Surname: args[1], Tel: args[2], LastAccess: strconv.FormatInt(time.Now().Unix(), 10)}
			mapRecords[args[2]] = record

			fmt.Println("Record", record, "added successfully")
		case "delete":
			fmt.Println("Type the following data: Telephone")

			args, _ := argsHelper(reader)

			if len(args) < 1 {
				fmt.Println("Missing arguments")
				continue
			}

			_, ok := mapRecords[args[0]]

			if ok {
				delete(mapRecords, args[0])
				fmt.Println("Record deleted successfully")
			} else {
				fmt.Println("Record not found")
			}
		case "list":
			for _, record := range mapRecords {
				fmt.Println(record)
			}
		case "search":
			fmt.Println("Type the following data: Telephone")

			args, _ := argsHelper(reader)

			if len(args) < 1 {
				fmt.Println("Missing arguments")
				continue
			}

			matches := exercises.TelMatcher(args[0])

			if !matches {
				continue
			}

			record, ok := mapRecords[args[0]]

			if ok {
				fmt.Println(record)
			} else {
				fmt.Println("Record not found")
			}
		case "save":
			records := make([]exercises.Record, 0, len(mapRecords))
			for _, record := range mapRecords {
				records = append(records, record)
			}

			err = exercises.SaveCSVFile(csvPath, records)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("CSV was updated")
		}
	}
}

func argsHelper(reader *bufio.Reader) ([]string, error) {
	input, _ := reader.ReadString('\n')
	args := strings.Fields(input)
	_, err := os.Stat(args[0])

	return args, err
}
