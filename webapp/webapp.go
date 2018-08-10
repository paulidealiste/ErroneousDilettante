//Package webapp enables the web-app functionality for the dilletante
package webapp

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Erroneus Dilettante</title>
		<link rel="stylesheet" href="https://unpkg.com/picnic">
		<script src="https://unpkg.com/umbrellajs"></script>
  </head>
  <body>
		<div class="flex one center">
			<div style="text-align: center" class="flex one center">
				<h2 class="third">{{.Title}}<span class="label success">{{.Content}}</span></h2>
			</div>
			<div style="font-size: .75em;" class="third">
				<article class="card">
					<header>
						<button id="wildbutton" class="success">Get wild!</button>
					</header>
					<footer>
						<div class="flex one label warning" style="border-bottom: 1px solid #2ecc40;">
							<h3>Public knowledge</h3>
						</div>
						<div id="erroneus" class="flex two" style="max-height: 350px; overflow: auto;"></div>
					</footer>
				</article>
			</div>
		</div>
  </body>
  <script>
	//Elements
	let lbtn = u('#wildbutton');
	let eroc = u('#erroneus');
	//Functions
	let renderResponse = (person) => {
		const namespan = '<span style="padding: .6em;">'+ person.Name +'</span>';
		eroc.append(namespan);
		const surnamespan = '<span span style="padding: .6em;">' + person.Surname + '</span>';
		eroc.append(surnamespan);
	};
	//Handlers
	lbtn.handle('click', async e => {
		const res = await fetch('/erroneus', { method: 'POST' });
		const data = await res.json();
		console.log('Response data: ', data);
		renderResponse(data);
	});
  </script>
</html>
`

func MockStart() {
	http.HandleFunc("/", dilettanteHandler)
	http.HandleFunc("/erroneus", erroneusHandler)
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
	tpl.Execute(w, ce)
}

func erroneusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ce := ErroneusResponse{Name: "Batmacumba!", Surname: "Bethania!"}
	a, _ := json.Marshal(ce)
	w.Write(a)
}
