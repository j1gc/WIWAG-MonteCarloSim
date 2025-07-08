package utils

import (
	"fmt"
	"io"
	"text/tabwriter"
)

type TablePrinter struct {
	header []any
	rows   [][]any
}

func NewTablePrinter() *TablePrinter {
	return &TablePrinter{}
}

func (t *TablePrinter) AddHeader(header ...any) {
	t.header = header
}

func (t *TablePrinter) AddRow(row ...any) {
	t.rows = append(t.rows, row)
}

func (t *TablePrinter) Print(output io.Writer) error {
	w := tabwriter.NewWriter(output, 5, 2, 2, ' ', 0)

	// print Header
	for _, header := range t.header {
		_, err := fmt.Fprintf(w, "%v\t", header)
		if err != nil {
			return err
		}
	}
	_, err := fmt.Fprint(w, "\n")
	if err != nil {
		return err
	}

	// print Rows
	for _, row := range t.rows {
		for _, cell := range row {
			_, err = fmt.Fprintf(w, "%v\t", cell)
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprint(w, "\n")
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "\n")
	if err != nil {
		return err
	}

	return w.Flush()
}
