package defs

import (
	"time"
)


// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// response
type SignedUp struct {
	Success bool `json: "success"`
	SessionId string `json: "session_id"`
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