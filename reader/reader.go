//Package reader provides reading and parsing of the .csv files for database fields
package reader

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/paulidealiste/ErroneusDilettante/models"
)

//Reader implements .csv reading methods and data aggregation
type Reader struct {
	primbuck models.PrimaryBucket
}

//FillNames reads in and structures the csv with names
func (r *Reader) FillNames(filepath string) error {
	readread, err := innerReader(filepath)
	r.primbuck.Names = readread
	return err
}

//FillSurnames reads in and structures the csv with surnames
func (r *Reader) FillSurnames(filepath string) error {
	readread, err := innerReader(filepath)
	r.primbuck.Surnames = readread
	return err
}

//FillReviews reads in and structures the csv with reviews
func (r *Reader) FillReviews(filepath string) error {
	readread, err := innerReader(filepath)
	r.primbuck.Reviews = readread
	return err
}

func innerReader(filepath string) ([]string, error) {
	var readread []string
	csvFile, err := os.Open(filepath)
	if err != nil {
		return readread, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return readread, err
		}
		readread = append(readread, line[0])
	}
	return readread, nil
}
