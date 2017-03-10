package main

import (
    "fmt"
    "os"
    "flag"
    "lucassabreu/pkg/parser"
)

func getHelp() {
    fmt.Println(`A simple command line tool to extract info from code coverage XMLs.`)
    os.Exit(0)
}

func main() {
    fileName := flag.String("f", "fileName", "a XML with the code coverage results")
    flag.Parse()

    fmt.Println("CodeCover: " + *fileName)

    parser.GetObjectFromXML(fileName)
}

