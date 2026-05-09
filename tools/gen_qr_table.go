//go:build ignore

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type BlockGroup struct {
	NumBlocks     int
	DataCodewords int
}

func main() {
	res, err := http.Get("https://www.thonky.com/qr-code-tutorial/error-correction-table")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error creating document:", err)
		return
	}

	var out strings.Builder
	out.WriteString("package internal\n\n")
	out.WriteString("import \"github.com/KhaledSaeed004/goqr/ecc\"\n\n")
	out.WriteString("var QRTable = map[int]map[ecc.ECLevel]QRSpec{\n")

	doc.Find("table tbody tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		version := (i-1)/4 + 1

		if (i-1)%4 == 0 {
			out.WriteString(fmt.Sprintf("\t%d: {\n", version))
		}

		cells := row.Find("td")
		if cells.Length() < 5 {
			return
		}

		level := strings.Split(cells.Eq(0).Text(), "-")[1]
		ecCodewordsPerBlock := cells.Eq(2).Text()

		out.WriteString(fmt.Sprintf("\t\tecc.%s: {ECCodewordsPerBlock: %s, Groups: []BlockGroup{", level, ecCodewordsPerBlock))
		blockGroups := []BlockGroup{}

		for i := 0; i < 2; i++ {
			groupNumBlocks, numerr := strconv.Atoi(strings.TrimSpace(cells.Eq(3 + i*2).Text()))
			groupDataCodewords, dataerr := strconv.Atoi(strings.TrimSpace(cells.Eq(4 + i*2).Text()))
			if numerr != nil || dataerr != nil {
				continue
			}

			blockGroups = append(blockGroups, BlockGroup{
				NumBlocks:     groupNumBlocks,
				DataCodewords: groupDataCodewords,
			})
		}

		for i, group := range blockGroups {
			if i > 0 {
				out.WriteString(", ")
			}
			out.WriteString(fmt.Sprintf("{NumBlocks: %d, DataCodewords: %d}", group.NumBlocks, group.DataCodewords))
		}
		out.WriteString("}},\n")

		if (i-1)%4 == 3 {
			out.WriteString("\t},\n")
		}
	})

	out.WriteString("}\n")
	fmt.Println(out.String())
}
