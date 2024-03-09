package main

import (
	"log"
	"os"
)

var (

    InfoLog   	*log.Logger
    ErrorLog   	*log.Logger
)

func init() {

    InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
