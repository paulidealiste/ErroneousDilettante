package database

import (
	"fmt"
	"testing"

	"github.com/paulidealiste/ErroneusDilettante/models"
)

func TestKickstartDB(t *testing.T) {
	teststore := Store{}
	teststore.KickstartDB("test.db")
}

func TestHoopEntities(t *testing.T) {
	teststore := Store{}
	teststore.KickstartDB("test.db")
	names := []string{"cevin", "skinny", "wiltus", "debbox"}
	err := teststore.HoopEntities(names, models.Names)
	fmt.Println(err)
	surnames := []string{"compton", "puppy", "mennya", "doota"}
	err = teststore.HoopEntities(surnames, models.Surnames)
	fmt.Println(err)
	reviews := []string{"stinky", "dinky", "blinky", "pinky"}
	err = teststore.HoopEntities(reviews, models.Reviews)
	fmt.Println(err)
	fmt.Println("Start reading")
	teststore.CheerEntities(models.Names)
	crucn, _ := teststore.CrunchEntities()
	fmt.Println(crucn)
}
