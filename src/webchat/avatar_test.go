// avatar_test.go
package main

import (
	"strings"
	"testing"
)

func TestAuthAvator(t *testing.T) {

	var authAvatar AuthAvatar

	client := &client{}

	testUrl := "http://url-to-gravatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}

	if url, err := authAvatar.GetAvatarURL(client); err != nil {
		t.Error("Avatar.GetAvatarURL shoudl return no error.")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURL should return correct URL")
		}
	}
}

func TestGravatar(t *testing.T) {
	var grAvatar GravatarAavatar

	client := &client{}

	client.userData = map[string]interface{}{"usrid": "89603fd30c1eb5cd0fe3c52f527746df"}
	url, err := grAvatar.GetAvatarURL(client)

	if err != nil {
		t.Error("GravatarAvitar.GetAvatarURL should not return an error")
	}

	if !strings.HasSuffix("//www.gravatar.com/avatars/89603fd30c1eb5cd0fe3c52f527746df.jpg", "/"+url) {
		t.Errorf("GravatarAvitar.GetAvatarURL wrongly returned %s", url)
	}
}
