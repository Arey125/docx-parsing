package main

import (
	"flag"
	"fmt"
	"os"

    "docx-parsing/internal/table"
)

func main() {
	path := flag.String("f", "", "file to parse")
	flag.Parse()

	if path == nil || *path == "" {
		fmt.Fprintln(os.Stderr, "Flag -f is not provided")
		return
	}

    tables, err := table.ExtractTables(*path)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    for _, tableNode := range tables {
        table.PrintTable(tableNode)
        fmt.Println()
    }
}
