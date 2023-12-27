package controllers

import (
    "time"
    "strconv"
    "encoding/json"
    "net/http"
    "new_messageboard/models"
    "github.com/gorilla/mux"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    m := new(models.Message)
    data := m.SelectMessage()

    json.NewEncoder(w).Encode(data)
}

func GetMessageById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    m := new(models.Message)
    m.Id = int64(id)

    data := m.GetMessageById()
    json.NewEncoder(w).Encode(data)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    m := new(models.Message)
    _ = json.NewDecoder(r.Body).Decode(&m)

    m.AddTime = time.Now().Format("2006-01-02 15:04:05")
    if m.Message != "" && m.AddTime != "" {
        m.Id = m.CreateMessage()
    }

    json.NewEncoder(w).Encode(m)
}