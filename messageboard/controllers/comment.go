package controllers

import (
    "fmt"
    "messageboard/models"
    "net/http"
    "strconv"
    "time"
    "github.com/alecthomas/template"
)

func Comment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    t, err := template.ParseFiles("./views/comment.html")
    if err != nil {
        fmt.Println(err)
        return
    }

    c := new (models.Comment)
    message_id, _ :=strconv.Atoi(r.FormValue("id"))
    c.MessageId = int64(message_id)
    switch r.Method {
        case "GET":
            m := new(models.Message)
            m.Id = int64(message_id)
            messageInfo := m.GetMessageById()
            commentList := c.SelectComment()

            type Data struct {
                MessageInfo models.Message
                CommentList []models.Comment
            }

            data := Data{
                MessageInfo: messageInfo,
                CommentList: commentList,
            }
            t.Execute(w, data)
            return
        case "POST":
            c.AddTime  = time.Now().Format("2006-01-02 15:04:05")
            c.Comment  = r.FormValue("comment")
            c.Nickname = r.FormValue("nickname")
            if c.Comment != "" && c.AddTime != "" && c.MessageId > 0 {
                _ = c.AddComment()
            }

            http.Redirect(w, r, fmt.Sprintf("/comment?id=?", c.MessageId), http.StatusFound)
            return
    }
}