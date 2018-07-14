// Package cmdos defines basic command-line flags
package cmdos

import "flag"

// Flags are the available fill-fields
type Flags struct {
	Names       *string
	Surnames    *string
	Reviews     *string
	BasePath    *string
	RandomPath  *string
	RandomQuant *string
}

// Initialize defines all the possible fill-fields
func Initialize() Flags {
	var flgs Flags
	flgs.Names = flag.String("names", "-", "path to .csv file with all names")
	flgs.Surnames = flag.String("surnames", "-", "path to .csv file with all surnames")
	flgs.Reviews = flag.String("reviews", "-", "path to .csv file with reviews")
	flgs.BasePath = flag.String("basepath", "-", "filename and path for a new or existing database")
	flgs.RandomPath = flag.String("randompath", "-", "filename and path for writing the results into")
	flgs.RandomQuant = flag.String("randomquant", "1", "number of randomized results in the output")
	return flgs
}
