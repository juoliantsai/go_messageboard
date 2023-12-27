package controllers

import (
	"fmt"
	"messageboard/models"
	"net/http"
	"time"
	"github.com/alecthomas/template"
)

func Message(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("./views/message.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	m := new(models.Message)
	switch r.Method {
		case "GET":
			data := m.SelectMessage()
			t.Execute(w, data)
			return
		case "POST":
			m.AddTime  = time.Now().Format("2006-01-02 15:04:05")
			m.Message  = r.FormValue("message")
			m.Nickname = r.FormValue("nickname")
			if m.Message != "" && m.AddTime != "" {
				_ = m.AddMessage()
			}
			http.Redirect(w, r, "/message", http.StatusFound)
			return
	}
}