package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

//App struct
type App struct {
	//Router is gorilla *rawr*
	Router *mux.Router
}

//NewApp create instance of app
func NewApp() *App {
	return &App{}
}

//Initialize : initializing the routing, db connection .etc
func (app *App) Initialize() {
	router := mux.NewRouter()

	router.HandleFunc("/register", app.Register)
	router.HandleFunc("/check-username", app.CheckUsername)
	router.HandleFunc("/login", app.Login)

	app.Router = router
}

//Run blablabla
func (app *App) Run(addr string) {
	server := http.Server{
		Handler: app.Router,
		Addr:    addr,
	}

	server.ListenAndServe()
}
