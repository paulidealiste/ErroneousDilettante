//Package models provides all thhe app-wide data structures and database models
package models

//PrimaryBucket provides first-stop station for the data from the .csv file
type PrimaryBucket struct {
	Names    []string
	Surnames []string
	Reviews  []string
}

//TargetBucket is the target container for the entities sent to database
type TargetBucket int

//Possible values of the target container
const (
	Names TargetBucket = iota + 1
	Surnames
	Reviews
)

//EntityModel provides direct database model for all of the database fields
type EntityModel struct {
	ID      string `json:"ID"`
	Content string `json:"Content"`
}
