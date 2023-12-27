package main
import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "new_messageboard/controllers"
    _"github.com/go-sql-driver/mysql"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/message", controllers.GetMessage).Methods("GET")
    r.HandleFunc("/message/{id}", controllers.GetMessageById).Methods("GET")
    r.HandleFunc("/message", controllers.CreateMessage).Methods("POST")

    fmt.Println("starting server at port:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}