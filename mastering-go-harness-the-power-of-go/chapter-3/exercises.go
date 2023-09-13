package exercises

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

type Record struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type PhoneBookEntry = map[int]Record

func TelMatcher(num string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(num)
}
func PhoneBook() {
	lines, err := ReadCSVFile("./phoneBook.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	records := make([]Record, len(lines))

	createRecords(&records, &lines)
}

func createRecords(records *[]Record, lines *[][]string) {
	for i, row := range *lines {
		(*records)[i] = Record{Name: row[0], Surname: row[1], Tel: row[2], LastAccess: row[3]}
	}
}

func ReadCSVFile(csvPath string) ([][]string, error) {
	_, err := os.Stat(csvPath)

	if err != nil {
		return nil, err
	}

	csvFile, err := os.Open(csvPath)

	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func SaveCSVFile(csvPath string, data []Record) error {
	csvFile, err := os.Create(csvPath)

	if err != nil {
		return err
	}

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	for _, record := range data {
		temp := []string{record.Name, record.Surname, record.Tel, record.LastAccess}

		err = csvWriter.Write(temp)
	}

	if err != nil {
		return err
	}

	csvWriter.Flush()

	return nil
}
