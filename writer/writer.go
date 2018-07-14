// Package writer provides basic file-based output functionality
package writer

import (
	"fmt"
	"os"
	"time"
)

//Writer provides all the methods for writing into .txt output files
type Writer struct {
	output    string
	OutputSet bool
}

//SetOutput sets the output path and filename for a writer
func (w *Writer) SetOutput(path string) {
	w.output = path
	w.OutputSet = true
}

//WriteToPath outputs new-line style of the supplied content to .txt file
func (w *Writer) WriteToPath(content []string) error {
	file, err := os.Create(w.output)
	if err != nil {
		return err
	}
	defer file.Close()
	t := time.Now()
	fmt.Fprintf(file, "ErroneusDilletante output for %s\n", t.Format(time.RFC1123))
	fmt.Fprintln(file, " ")
	for _, core := range content {
		fmt.Fprintln(file, core)
	}
	return nil
}
