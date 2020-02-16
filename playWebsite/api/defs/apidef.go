package defs

import (
	"time"
)

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// data model

type VideoInfo struct {
	VideoId string
	AuthorId int
	Name string
	DisplayCtime string
	CreateTime time.Time
}

type Comment struct {
	CommentId string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string
	TTL int64
}