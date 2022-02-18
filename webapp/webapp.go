//Package webapp enables the web-app functionality for the dilletante
package webapp

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/paulidealiste/ErroneusDilletante/database"
)

var dbhand dbaseHandler

func init() {
	fs := http.FileServer(http.Dir("./webapp/content"))
	http.Handle("/home/", http.StripPrefix("/home/", fs))
}

type dbaseHandler struct {
	dbs *database.Store
}

func (dbh *dbaseHandler) respond() ErroneusResponse {
	var resp ErroneusResponse
	if dbh.dbs == nil {
		return resp
	}
	respstring, err := dbh.dbs.CrunchEntities()
	fmt.Println(respstring)
	if err != nil {
		return resp
	}
	resplist := strings.Split(respstring, " ")
	resp.Name = resplist[0]
	resp.Surname = resplist[1]
	return resp
}

//MockStart provides simplified means of starting the webapp -temporary
func MockStart(dbs *database.Store) {
	dbhand = dbaseHandler{dbs: dbs}

	http.HandleFunc("/", dilettanteHandler)
	http.HandleFunc("/erroneus", erroneusHandler)
	log.Println("Listening on localhost:8080/home")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//DilletanteWelcome embodies all the initial page structural fields
type DilletanteWelcome struct {
	Title   string
	Content string
}

//ErroneusResponse embodies one particular instance of response
type ErroneusResponse struct {
	Name    string
	Surname string
}

func dilettanteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ce := DilletanteWelcome{Title: "Ravager", Content: "Who awaits the incantation?"}
	ip := filepath.Join("index.html")
	tpl := template.Must(template.ParseFiles(ip))
	tpl.Execute(w, ce)
}

func erroneusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ce := dbhand.respond()
	a, _ := json.Marshal(ce)
	w.Write(a)
}
