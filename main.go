package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dbNamePtr := flag.String("db", "", "SQLite database to use")
	flag.Parse()
	if *dbNamePtr == "" {
		fmt.Println("Database is required")
		os.Exit(1)
	}

}
