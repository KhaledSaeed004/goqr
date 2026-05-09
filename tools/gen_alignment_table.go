//go:build ignore

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AlignmentCenters []int

func main() {
	res, err := http.Get("https://www.thonky.com/qr-code-tutorial/alignment-pattern-locations")
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
	out.WriteString("var AlignmentTable = map[int][]int{\n")
	out.WriteString("\t1: {},\n")

	doc.Find("table tbody tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		if row.Children().Length() < 2 {
			return
		}

		version := strings.TrimPrefix(row.Find("td").First().Text(), "QR Version ")
		versionNum, err := strconv.Atoi(version)
		if err != nil {
			return
		}

		out.WriteString(fmt.Sprintf("\t%d: {", versionNum))

		alignmentCenters := AlignmentCenters{}

		row.Find("td").Each(func(j int, td *goquery.Selection) {
			if j == 0 {
				return
			}

			centerModule, err := strconv.Atoi(strings.TrimSpace(td.Text()))
			if err != nil {
				return
			}
			alignmentCenters = append(alignmentCenters, centerModule)
		})

		for k, center := range alignmentCenters {
			if k > 0 {
				out.WriteString(", ")
			}
			out.WriteString(fmt.Sprintf("%d", center))
		}
		out.WriteString("},\n")
	})

	out.WriteString("}\n")
	fmt.Println(out.String())
}
