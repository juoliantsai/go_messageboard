package models

import (
    "fmt"
)

type Message struct {
    Id         int64  `db:"id"`
    PraiseNum  int64  `db:"praise_num"`
    CommentNum int64  `db:"comment_num"`
    Nickname   string `db:"nickname"`
    Message    string `db:"message"`
    AddTime    string `db:"add_time"`
}

func (m *Message) SelectMessage() (message []Message) {
    err := Db.Select(&message, "SELECT id, nickname, message, add_time, praise_num, comment_num FROM message ORDER BY id DESC")
    if err != nil {
        fmt.Println(err)
    }

    return
}

func (m *Message) GetMessageById() (message Message) {
    err := Db.Get(&message, "SELECT id, nickname, message, add_time, praise_num, comment_num FROM message WHERE id = ?", m.Id)
    fmt.Println(m.Id)
    if err != nil {
        fmt.Println(err)
    }
    return
}

func (m *Message) CreateMessage() int64 {
    r, err := Db.Exec("INSERT INTO message(message, nickname, add_time) VALUES(?, ?, ?)", m.Message, m.Nickname, m.AddTime)

    if err != nil {
        fmt.Println(err)
        return 0
    }

    id, err := r.LastInsertId()
    if err != nil {
        fmt.Println(err)
        return 0
    }

    return id
}

func (m *Message) UpdateMessage() int64 {
    _, err := Db.Exec("UPDATE message SET message = ?, nickname = ? WHERE id = ?", m.Message, m.Nickname, m.Id)

    if err != nil {
        fmt.Println(err)
    }

    return 0
}

func (m *Message) DeleteMessage() string {
    _, err := Db.Exec("DELETE FROM message WHERE id = ?", m.Id)

    if err != nil {
        return "fail"
    }

    return "success"
}
