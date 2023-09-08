package exercises

import (
	"errors"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path"
	"path/filepath"
)

func Which(arguments ...string) error {
	if len(arguments) <= 1 {
		fmt.Println("Please proveide an argument.")
		return errors.New("please proveide an argument")
	}

	file := arguments[1]
	pathValue := os.Getenv("PATH")
	pathSplit := filepath.SplitList(pathValue)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			mode := fileInfo.Mode()

			if mode.IsRegular() {
				if mode&0o111 != 0 {
					fmt.Println(fullPath)
				}
			}
		}
	}

	return nil
}

func SysLog() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	if err != nil {
		log.Println(err)
		return
	}

	log.SetOutput(sysLog)
	log.Print("Everything is fine")
}

func CustomLog() {
	logFile := path.Join(os.TempDir(), "mGO.log")
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	iLog := log.New(f, "iLog", log.LstdFlags)
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)

	iLog.Println("Hello there!")
	iLog.Println("Mastering Go")
}

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{
	{Name: "Mihalis", Surname: "Tsoukalos", Tel: "789198489"},
	{Name: "Mary", Surname: "Doe", Tel: "210564891"},
	{Name: "John", Surname: "Black", Tel: "469168482"},
}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}

	return nil
}
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func PhoneBook(args ...string) {
	if len(args) <= 1 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)

		return
	}

	switch args[1] {
	case "search":
		if len(args) <= 2 {
			fmt.Println("Usage: search 'Surname'")
		}

		result := search(args[2])

		if result == nil {
			fmt.Println("Entry not found:", args[2])
			return
		}

		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
