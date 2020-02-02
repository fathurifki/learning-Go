package main

import (
 "log"
 "os"
)

func main(){
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connect string: %s\n", connStr)
}


