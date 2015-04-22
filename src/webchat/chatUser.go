package main

import (
	gomniauthcomm "github.com/stretchr/gomniauth/common"
)

type ChatUser interface {
	AvatarURL() string
	UniqueID() string
}

type chatUser struct {
	gomniauthcomm.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}
