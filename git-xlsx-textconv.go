package main

import (
	"fmt"
	"github.com/WeiZhang555/tabwriter"
	"log"
	"os"
	"strings"

	xlsx "github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: git-xlsx-textconv file.xslx")
	}
	excelFileName := os.Args[1]

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 0, '.', tabwriter.Debug)
		for _, row := range sheet.Rows {
			if row == nil {
				continue
			}
			cels := make([]string, len(row.Cells))
			for i, cell := range row.Cells {
				s, _ := cell.String()
				s = strings.Replace(s, "\\", "\\\\", -1)
				s = strings.Replace(s, "\n", "\\n", -1)
				s = strings.Replace(s, "\r", "\\r", -1)
				s = strings.Replace(s, "\t", "\\t", -1)
				cels[i] = s
			}
			fmt.Fprintf(w, "[%s] %s\n", sheet.Name, strings.Join(cels, "\t"))
		}
		w.Flush()
	}
}
