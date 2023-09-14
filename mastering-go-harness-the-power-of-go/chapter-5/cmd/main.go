package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Reader struct {
	bufio *bufio.Reader
}

func main() {
	reader := Reader{bufio: bufio.NewReader(os.Stdin)}

	var conn string
dbDetails:
	fmt.Println("Please provide: hostname port username password db / (or leave empty for defaults)")

	for {
		input := reader.Read()
		if len(input) < 1 {
			input = []string{"127.0.0.1", "5432", "postgres", "postgres1234", "go"}
		} else if len(input) < 6 {
			goto dbDetails
		}

		host := input[0]
		p := input[1]
		user := input[2]
		pass := input[3]
		database := input[4]
		port, err := strconv.Atoi(p)

		if err != nil {
			fmt.Println("Not a valid port number:", err)
			goto dbDetails
		}

		conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, database)

		break
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		goto dbDetails
	}

	defer db.Close()

	query := `SELECT table_name FROM information_schema.tables WHERE
	table_schema = 'public' ORDER BY table_name`
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Query", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)

		if err != nil {
			fmt.Println("Scan", err)
			return
		}

		fmt.Println("*", name)
	}
}

func (r *Reader) Read() []string {
	input, err := r.bufio.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	return strings.Fields(input)
}
