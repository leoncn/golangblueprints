// upload.go
package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

func uploaderHandler(w http.ResponseWriter, r *http.Request) {
	usrId := r.FormValue("usrid")

	file, header, err := r.FormFile("avatarFile")

	if err == nil {
		data, err := ioutil.ReadAll(file)

		if err == nil {
			filename := path.Join("avatars", usrId+path.Ext(header.Filename))
			err = ioutil.WriteFile(filename, data, 0777)

			if err != nil {
				io.WriteString(w, err.Error())
				return
			} else {
				io.WriteString(w, "Successful")
			}
		} else {
			io.WriteString(w, err.Error())
			return
		}

	} else {
		io.WriteString(w, err.Error())
		return
	}
}
