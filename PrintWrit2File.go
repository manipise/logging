package main

import (
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	format := "txt"
	var data = [][]string{
		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "0"},        /*  initializers for row indexed by 0 */
		{"11", "12", "13", "14", "15", "16", "17", "8", "9", "10", "0"}, /*  initializers for row indexed by 1 */
		{"21", "22", "23", "24", "25", "26", "27", "8", "9", "10", "0"}, /*  initializers for row indexed by 2 */
	}
	if format == "" {
		printPattern(data, format)
	}

	if format == "txt" {
		printPattern(data, format)

	}
}
func printPattern(data [][]string, format string) {
	w := io.MultiWriter(os.Stdout)
	if format == "txt" {
  name:=report.txt
  
		file, err := os.Create(name)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		w = io.MultiWriter(os.Stdout, file)
	}
  	if format == "csv" {
  name:=report.csv
  
		file, err := os.Create(name)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		w = io.MultiWriter(os.Stdout, file)
	}

	fmt.Fprintf(w, "%s", fmt.Sprintln("data Report"))
	fmt.Fprintf(w, "%s", fmt.Sprintln("ExecutionId: ", data[0][0]))
	fmt.Fprintf(w, "%s", fmt.Sprintln("SpecId: ", data[0][1]))
	fmt.Fprintf(w, "%s", fmt.Sprintln("Operation: ", data[0][2]))
	fmt.Fprintf(w, "%s", fmt.Sprintln("ExecutedBy: ", data[0][3]))
	fmt.Fprintf(w, "%s", fmt.Sprintln("ExecStatus: ", data[0][4]))
	fmt.Fprintf(w, "%s", fmt.Sprintln("ExecutedStart: ", data[0][5]))
	fmt.Fprintf(w, "%s", fmt.Sprintln("ExecutedEnd: ", data[0][6]))

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"PoolName", "JobStatus", "Status", "Message"})
	table.SetBorder(false)
	var trimval [][]string
	for i := 0; i < len(data); i++ {
		trimval = append(trimval, data[i][7:])
	}
	table.AppendBulk(trimval) // Add Bulk Data
	table.Render()

}
