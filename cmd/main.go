package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
    path := flag.String("f", "", "file to parse")
    flag.Parse()

    if path == nil || *path == "" {
        fmt.Fprintln(os.Stderr, "Flag -f is not provided")
        return;
    }

	archive, err := zip.OpenReader(*path)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Cannot extract files from %s\n", *path)
        return
    }
    defer archive.Close()

    doc, err := archive.Open("word/document.xml")
    if err != nil {
        panic(err)
    }
    defer doc.Close()

	fileContent, err := xmlquery.Parse(doc)
	if err != nil {
		panic(err)
	}
	for _, row := range fileContent.SelectElements("//w:tr") {
		for _, cell := range row.SelectElements("//w:tc") {
			fmt.Printf("%s; ", strings.TrimSpace(cell.InnerText()))
		}
		fmt.Print("\n")
	}
}
