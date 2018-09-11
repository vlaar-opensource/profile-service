package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vlaar-opensource/profile-service/db"
)

//App struct
type App struct {
	//Router is gorilla *rawr*
	Router *mux.Router
	DB     db.IfaceDB
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
	db, err := db.NewDBPostresImpl("localhost", 5432, "postgres", "mydb", "secret")
	app.DB = db

	if err != nil {
		os.Exit(22)
	}
}

//Run blablabla
func (app *App) Run(addr string) {
	server := http.Server{
		Handler: app.Router,
		Addr:    addr,
	}
	fmt.Println("> Running on :: ", addr)
	server.ListenAndServe()
}
