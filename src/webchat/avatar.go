package main

import (
	"errors"
)

var ErrorNoAvatarURL = errors.New("chat : not able to get avatar URL.")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UserAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if url_str, ok := url.(string); ok {
			return url_str, nil
		}
	}
	return "", ErrorNoAvatarURL
}

type GravatarAavatar struct{}

var UseGravatar GravatarAavatar

func (_ GravatarAavatar) GetAvatarURL(c *client) (string, error) {
	if usrid, ok := c.userData["usrid"]; ok {
		if usrid_str, ok := usrid.(string); ok {

			return "avatars/" + usrid_str + ".jpg", nil
		}
	}
	return "", ErrorNoAvatarURL
}
