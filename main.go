package main


import (
    "fmt"
    "log"
	"flag"
	"embed"
	"io/fs"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	uc "github.com/kindercare2022/kindercare-server/usecases"
)

//go:embed client/web/*
var static embed.FS

func htmlWebsite(w http.ResponseWriter, r *http.Request) {
	website, _ := fs.Sub(static, "client")
    handler := http.FileServer(http.FS(website))
    handler.ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct{Success string}{Success: "API home"})
}

func resourceNotFound(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct{Success string}{Success: "The API doesn't have what you are looking for !"})
}

func getRouter() *mux.Router {
	website, _ := fs.Sub(static, "client/web")
	router := mux.NewRouter()
	router.HandleFunc("/user/login", uc.UserLogin).Methods("POST")
	router.HandleFunc("/register/pupil", uc.RegisterPupil).Methods("POST")
	router.HandleFunc("/register/teacher", uc.RegisterTeacher).Methods("POST")
	
	//Home
	router.HandleFunc("/", index ).Methods("POST")
	router.PathPrefix("/").Handler( http.FileServer(http.FS(website)) ).Methods("GET")
	
	//Not found
	router.NotFoundHandler = http.HandlerFunc(resourceNotFound)
	
	return router
}

func main() {
    //++++| os.Args |+++++
    wsEndPoint := ":6200" 
    addr := flag.String("addr", wsEndPoint, "KinderCare API service address") 
    flag.Parse()
    //++++++++++++++++++++
    uc.Init()
    
    fmt.Println("Server listening on port: "+(strings.Split(wsEndPoint,":")[1])) 
    log.Fatal(http.ListenAndServe(*addr, getRouter()))
}








