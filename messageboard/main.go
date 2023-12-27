package main
import(
    "fmt"
    "messageboard/controllers"
    "net/http"
    _"github.com/go-sql-driver/mysql"
)

func main() {
    routers()
}
func routers() {
    http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/message", controllers.Message)
    http.HandleFunc("/comment", controllers.Comment)
    if err := http.ListenAndServe(":8080", nil); err != nil{
        fmt.Println(err)
    }
}
