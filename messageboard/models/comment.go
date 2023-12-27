package models

import (
	"fmt"
)

type Comment struct {
	Id        int64  `db:"id"`
	MessageId int64  `db:"message_id"`
	Nickname  string `db:"nickname"`
	Comment   string `db:"comment"`
	AddTime   string `db:"add_time"`
}

func (c *Comment) AddComment() int64 {
	conn, err := Db.Begin()
	r, err := Db.Exec("INSERT INTO comment(message_id, nickname, comment, add_time) VALUES(?, ?, ?, ?)", c.MessageId, c.Nickname, c.Comment, c.AddTime)

	if err != nil {
		fmt.Println(err)
		conn.Rollback()
		return 0
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println(err)
		conn.Rollback()
		return 0
	}

	res, err := Db.Exec("UPDATE message SET comment_num = comment_num+1 WHERE id = ?", c.MessageId)
	if err != nil {
		fmt.Println(err)
		conn.Rollback()
		return 0
	}

	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		conn.Rollback()
	}
	conn.Commit()

	return id
}

func (c *Comment) SelectComment() (comment []Comment) {
	err := Db.Select(&comment, "SELECT id, message_id, nickname, comment, add_time FROM comment WHERE message_id = ? ORDER BY id DESC", c.MessageId)
	if err != nil {
		fmt.Println(err)

		return nil
	}

	return
}