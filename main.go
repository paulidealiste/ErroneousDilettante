package main

import (
	"flag"
	"fmt"

	"github.com/paulidealiste/ErroneusDilettante/cmdos"
	"github.com/paulidealiste/ErroneusDilettante/database"
	"github.com/paulidealiste/ErroneusDilettante/reader"
)

func main() {
	flgs := cmdos.Initialize()
	flag.Parse()
	rawreader := reader.Reader{}
	rawdbaser := database.Store{}
	if *flgs.Names != flag.Lookup("names").DefValue {
		nerr := rawreader.FillNames(*flgs.Names)
		if nerr != nil {
			fmt.Println("Names could not be loaded or not provided.", nerr)
		}
	}
	if *flgs.Surnames != flag.Lookup("surnames").DefValue {
		serr := rawreader.FillSurnames(*flgs.Surnames)
		if serr != nil {
			fmt.Println("Surnames could not be loaded or not provided.", serr)
		}
	}
	if *flgs.Reviews != flag.Lookup("reviews").DefValue {
		rerr := rawreader.FillReviews(*flgs.Reviews)
		if rerr != nil {
			fmt.Println("Reviews could not be loaded or not provided.", rerr)
		}
	}
	if *flgs.BasePath != flag.Lookup("basepath").DefValue {
		rawdbaser.KickstartDB(*flgs.BasePath)
	}

}
