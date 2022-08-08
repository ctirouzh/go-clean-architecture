package main

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Println("Cli tool for managing database")
	fmt.Println("	Examples")
	fmt.Println("		go run cmd/migrate.go db-migrate")
	fmt.Println("		go run cmd/migrate.go drop-db")
}

func DBMigrate() {

}

func DropDB() {

}

func main() {
	if len(os.Args) <= 1 {
		Help()
		return
	} else {
		cmd := os.Args[1]
		switch cmd {
		case "db-migrate":
			DBMigrate()
		case "drop-db":
			DropDB()
		default:
			fmt.Printf("Bad command: %s\n", cmd)
		}
	}
}
