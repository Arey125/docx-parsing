package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

func printTable(table *xmlquery.Node) {
	for _, row := range table.SelectElements("//w:tr") {
		for _, cell := range row.SelectElements("//w:tc") {
			spanTag := cell.SelectElement("//w:gridSpan")
			span := 1
			if spanTag != nil {
				spanStr := spanTag.SelectAttr("w:val")
				if spanStr != "" {
					span, _ = strconv.Atoi(spanStr)
				}
			}
			fmt.Printf("%s;\t", strings.TrimSpace(cell.InnerText()))
			if span > 1 {
				for i := 1; i < span; i++ {
					fmt.Print(";\t")
				}
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	path := flag.String("f", "", "file to parse")
	flag.Parse()

	if path == nil || *path == "" {
		fmt.Fprintln(os.Stderr, "Flag -f is not provided")
		return
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

    for _, table := range fileContent.SelectElements("//w:tbl") {
        printTable(table)
        fmt.Print("\n")
    }
}
