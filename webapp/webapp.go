//Package webapp enables the web-app functionality for the dilletante
package webapp

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/paulidealiste/ErroneusDilettante/database"
)

var tpl *template.Template
var dbhand dbaseHandler

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <title>Erroneus Dilettante</title>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/uikit/3.0.0-rc.10/css/uikit.min.css" />
  <script src="https://cdnjs.cloudflare.com/ajax/libs/uikit/3.0.0-rc.10/js/uikit.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/uikit/3.0.0-rc.10/js/uikit-icons.min.js"></script>
  <script src="https://unpkg.com/umbrellajs"></script>
</head>

<body class="uk-overflow-hidden">
  <div class="uk-flex uk-width-1-1 uk-height-viewport uk-flex-stretch uk-flex-column">
    <div class="uk-active uk-background-primary">
      <nav class="uk-navbar-container uk-margin uk-margin-small-bottom" uk-navbar>
        <div class="uk-navbar-center">
          <a href="" class="uk-navbar-item uk-logo">
            <div>
              REPROBATE
              <div class="uk-navbar-subtitle">Pleasure Dome!</div>
            </div>
          </a>
        </div>
      </nav>
    </div>
    <div class="uk-flex uk-flex uk-width-1-1 uk-height-1-1 uk-flex-1 uk-background-muted uk-padding-small uk-text-small">
      <div class="uk-card uk-card-default uk-width-1-4 uk-padding-remove">
        <div class="uk-card-header">
          <button id="wildbutton" class="uk-button uk-button-default uk-width-1-1">Call erroneus</button>
        </div>
        <div class="uk-flex uk-width-1-1 uk-card-body uk-overflow-auto uk-height-max-small uk-padding-remove">
          <div id="DilettanteName" class="uk-width-1-1">
          </div>
        </div>
      </div>
      <div class="uk-card uk-card-default uk-flex uk-flex-column uk-flex-1 uk-margin-small-left">
        <div class="uk-card-header uk-flex">
          <div class="uk-flex-1">
            <h4 id="chosenvictim" class="uk-margin-remove uk-text-uppercase">&nbsp;</h4>
          </div>
          <div class="uk-margin-small-left">
            <button id="rantbutton" class="uk-button uk-button-default uk-width-1-1">Save this rant</button>
          </div>
        </div>
        <div class="uk-card-body uk-padding-small uk-flex uk-flex-1">
          <form class="uk-height-1-1 uk-width-1-1">
            <fieldset class="uk-fieldset uk-height-1-1 uk-width-1-1">
              <div class="uk-margin uk-height-1-1 uk-width-1-1">
                <textarea class="uk-textarea uk-height-1-1 uk-width-1-1" placeholder="Rant if not any!" style="resize: none"></textarea>
              </div>
            </fieldset>
          </form>
        </div>
      </div>
    </div>
  </div>
</body>
<script>
  //Elements
  let lbtn = u('#wildbutton');
  let rbtn = u('#rantbutton');
  let eron = u('#DilettanteName');
  let victim = u('#chosenvictim');
  //Interactions
  let chooseNameli = (e) => {
    let elm = u(e.target);
    u('div.reprobationRow').removeClass('uk-background-secondary uk-light');
    elm.parent().addClass('uk-background-secondary uk-light');
    victim.text(elm.text());
  };
  //Functions
  let renderResponse = (person) => {
    const nameli = '<div class="reprobationRow uk-padding-small" style="border-bottom: 1px solid grey"><span style="cursor: pointer;" class="uk-width-1-1">' + person.Name + ' ' + person.Surname + '</span></div>';
    eron.append(nameli)
      .find('span')
      .off('click')
      .handle('click', chooseNameli);
  };
  //Handlers
  lbtn.handle('click', async e => {
    const res = await fetch('/erroneus', {
      method: 'POST'
    });
    const data = await res.json();
    renderResponse(data);
  });
</script>

</html>
`

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
	fmt.Println(resplist)
	resp.Name = resplist[0]
	resp.Surname = resplist[1]
	return resp
}

func MockStart(dbs *database.Store) {
	dbhand = dbaseHandler{dbs: dbs}
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
	ce := dbhand.respond()
	a, _ := json.Marshal(ce)
	w.Write(a)
}
