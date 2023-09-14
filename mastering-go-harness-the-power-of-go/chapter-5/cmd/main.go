package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	Username string
	ID       int
}

type Userdata struct {
	Username    string
	Name        string
	Surname     string
	Description string
	ID          int
}

type Reader struct {
	bufio *bufio.Reader
}

var conn string

var db *sql.DB
var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

func main() {
	var err error

	reader := Reader{bufio: bufio.NewReader(os.Stdin)}

dbDetails:
	fmt.Println("Please provide: hostname port username password db / (or leave empty for defaults)")

	for {
		input := reader.Read()
		if len(input) < 1 {
			input = []string{"127.0.0.1", "5432", "postgres", "postgres1234", "go"}
		} else if len(input) < 6 {
			goto dbDetails
		}

		p := input[1]

		Hostname = input[0]
		Port, err = strconv.Atoi(p)
		Username = input[2]
		Password = input[3]
		Database = input[4]

		if err != nil {
			fmt.Println("Not a valid port number:", err)
			goto dbDetails
		}

		break
	}

	db, _ = OpenConnection()
	defer db.Close()
dbTasks:
	for {
		fmt.Println("Options are AddUser|UpdateUser|DeleteUser|ListUsers|Quit")

		input := reader.Read()

		switch input[0] {
		case "AddUser":
			fmt.Println("Type: (username) (name) (surname) (description)")

			input = reader.Read()

			if len(input) < 4 {
				fmt.Println("Not enough arguments")
				break
			}
			user := Userdata{input[0], input[1], input[2], input[3], 0}

			id := AddUser(user)

			if id == -1 {
				fmt.Println("There was an error")
			}
			fmt.Println("User created successfully")

			continue
		case "UpdateUser":
			fmt.Println("Type: (username) (name) (surname) (description)")

			input = reader.Read()

			if len(input) < 4 {
				fmt.Println("Not enough arguments")
				continue
			}
			user := Userdata{input[0], input[1], input[2], input[3], 0}

			err = UpdateUser(user)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("User updated successfully")

			continue
		case "DeleteUser":
			fmt.Println("Type the user id: (id)")

			input = reader.Read()

			if len(input) < 1 {
				fmt.Println("Not enough arguments")
				continue
			}
			userID, err := strconv.Atoi(input[0])

			if err != nil {
				fmt.Println(err)
				continue
			}

			err = DelUser(userID)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("User deleted successfully")

		case "ListUsers":
			fmt.Println("Listing users...")

			users, err := ListUsers()

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(users)
			continue
		}

		goto dbTasks
	}
}

func (r *Reader) Read() []string {
	input, err := r.bufio.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	return strings.Fields(input)
}

func OpenConnection() (*sql.DB, error) {
	conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password, Database)
	adb, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Println("Open():", err)
		return nil, err
	}

	return adb, nil
}

func Exists(username string) (int, error) {
	username = strings.ToLower(username)

	if db == nil {
		return -1, errors.New("there's not an active connection to the DB")
	}

	userID := -1
	statement := `SELECT "id" FROM "users" WHERE username = $1`
	rows, err := db.Query(statement, username)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)

		if err != nil {
			fmt.Println("Scan", err)
			return -1, nil
		}

		userID = id
	}

	defer rows.Close()

	return userID, nil
}

func AddUser(userData Userdata) int {
	userData.Username = strings.ToLower(userData.Username)

	userID, err := Exists(userData.Username)

	if err != nil {
		fmt.Println(err)

		return -1
	}

	if userID != -1 {
		fmt.Println("User already exists:", Username)

		return -1
	}

	insertStatement := `insert into "users" ("username") values ($1)`

	_, err = db.Exec(insertStatement, userData.Username)

	if err != nil {
		fmt.Println(err)

		return -1
	}

	userID, err = Exists(userData.Username)
	if userID == -1 {
		return userID
	}

	if err != nil {
		fmt.Println(err)

		return -1
	}

	insertStatement = `insert into "userdata" ("userid", "name", "surname", "description") values ($1, $2, $3, $4)`

	_, err = db.Exec(insertStatement, userID, userData.Name, userData.Surname,
		userData.Description)

	if err != nil {
		fmt.Println("db.Exec()", err)
		return -1
	}

	return userID
}

func DelUser(userID int) error {
	statement := fmt.Sprintf(`SELECT "username" FROM "users" where id = %d`, userID)
	rows, err := db.Query(statement)

	var username string

	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return err
		}
	}
	defer rows.Close()

	var user, _ = Exists(username)

	if user != userID {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	deleteStatement := `delete from "userdata" where userid=$1`
	_, err = db.Exec(deleteStatement, userID)

	if err != nil {
		return err
	}

	deleteStatement = `delete from "users" where userId=$1`
	_, err = db.Exec(deleteStatement, userID)

	if err != nil {
		return err
	}

	return nil
}

func ListUsers() ([]Userdata, error) {
	Data := []Userdata{}

	rows, err := db.Query(`SELECT "id","username","name","surname","description" FROM "users","userdata" WHERE users.id = userdata.userid`)
	if err != nil {
		return Data, err
	}

	for rows.Next() {
		var id int

		var username, name, surname, description string

		err = rows.Scan(&id, &username, &name, &surname, &description)
		temp := Userdata{ID: id, Username: username, Name: name,
			Surname: surname, Description: description}
		Data = append(Data, temp)

		if err != nil {
			return Data, err
		}
	}

	defer rows.Close()

	return Data, nil
}

func UpdateUser(d Userdata) error {
	userID, _ := Exists(d.Username)
	if userID == -1 {
		return errors.New("User does not exist")
	}

	d.ID = userID
	updateStatement := `update "userdata" set "name"=$1, "surname"=$2,
"description"=$3 where "userid"=$4`

	_, err := db.Exec(updateStatement, d.Name, d.Surname, d.Description,
		d.ID)

	if err != nil {
		return err
	}

	return nil
}
