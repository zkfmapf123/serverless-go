package filesystem

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintTable[T interface{}](list map[string]T, appendProperties []string, appendPropertiesFunc func(k string, v T) []string) {
	var tableData [][]string
	for k, v := range list {
		tableData = append(tableData, appendPropertiesFunc(k, v))
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(appendProperties)

	for _, row := range tableData {
		table.Append(row)
	}

	table.Render()
}
