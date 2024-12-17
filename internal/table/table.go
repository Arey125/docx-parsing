package table

import (
	"fmt"
    "strconv"
	"strings"
	"archive/zip"

	"github.com/antchfx/xmlquery"
)

func ExtractTables(path string) ([]*xmlquery.Node, error) {
	archive, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer archive.Close()

	doc, err := archive.Open("word/document.xml")
	if err != nil {
        return nil, err
	}
	defer doc.Close()

	fileContent, err := xmlquery.Parse(doc)
	if err != nil {
        return nil, err
	}

    return fileContent.SelectElements("//w:tbl"), nil
}

func PrintTable(table *xmlquery.Node) {
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
