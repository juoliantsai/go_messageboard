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

    r.HandleFunc("/movies", controllers.GetMessage).Methods("GET")
    r.HandleFunc("/movie/{id}", controllers.GetMessageById).Methods("GET")
    r.HandleFunc("/movie", controllers.CreateMessage).Methods("POST")
    r.HandleFunc("/movie/{id}", controllers.UpdateMessage).Methods("PUT")

    fmt.Println("starting server at port:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}