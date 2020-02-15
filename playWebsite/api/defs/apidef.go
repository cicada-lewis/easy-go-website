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

