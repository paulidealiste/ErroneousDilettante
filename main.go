package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/paulidealiste/ErroneusDilettante/cmdos"
	"github.com/paulidealiste/ErroneusDilettante/database"
	"github.com/paulidealiste/ErroneusDilettante/models"
	"github.com/paulidealiste/ErroneusDilettante/reader"
	"github.com/paulidealiste/ErroneusDilettante/webapp"
	"github.com/paulidealiste/ErroneusDilettante/writer"
)

func main() {
	flgs := cmdos.Initialize()
	flag.Parse()
	rawreader := reader.Reader{}
	rawdbaser := database.Store{}
	rawwriter := writer.Writer{}
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
		databaseFiller(&rawreader, &rawdbaser)
	} else {
		log.Fatal("Database path is mandatory.")
	}
	if *flgs.RandomPath != flag.Lookup("randompath").DefValue {
		rawwriter.SetOutput(*flgs.RandomPath)
	}
	var repeats int
	if *flgs.RandomQuant == flag.Lookup("randomquant").DefValue {
		repeats = 1
	} else {
		rpts, err := strconv.Atoi(*flgs.RandomQuant)
		if err != nil {
			log.Fatal("Quantity probably not an integer number!")
		}
		repeats = rpts
	}
	if *flgs.PerformCrunch == true {
		loopCrunchPrinter(repeats, &rawwriter, &rawdbaser)
	}
	if *flgs.StartWebserver == true {
		webapp.MockStart()
	}
}

func databaseFiller(rreader *reader.Reader, dbaser *database.Store) {
	if len(rreader.Primbuck.Names) > 0 {
		nerr := dbaser.HoopEntities(rreader.Primbuck.Names, models.Names)
		fmt.Println(nerr)
	}
	if len(rreader.Primbuck.Surnames) > 0 {
		serr := dbaser.HoopEntities(rreader.Primbuck.Surnames, models.Surnames)
		fmt.Println(serr)
	}
	if len(rreader.Primbuck.Reviews) > 0 {
		rerr := dbaser.HoopEntities(rreader.Primbuck.Reviews, models.Reviews)
		fmt.Println(rerr)
	}
}

func loopCrunchPrinter(repeats int, wwriter *writer.Writer, dbaser *database.Store) {
	var loopcrunched []string
	for j := 0; j < repeats; j++ {
		strng, err := dbaser.CrunchEntities()
		if err != nil {
			panic(err)
		}
		loopcrunched = append(loopcrunched, strng)
	}
	if wwriter.OutputSet == true {
		wwriter.WriteToPath(loopcrunched)
	} else {
		for _, lpc := range loopcrunched {
			fmt.Println(lpc)
		}
	}
}
