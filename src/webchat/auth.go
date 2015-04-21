package main

import (
	"crypto/md5"
	"fmt"
	"github.com/stretchr/gomniauth"
	"io"

	"github.com/stretchr/objx"
	"log"
	"net/http"
	"strings"
)

type authHandler struct {
	next http.Handler
}

func (h authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		//no cookie here, redirect to login auth
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		//log err
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func MustAuth(hh http.Handler) http.Handler {
	return &authHandler{next: hh}
}

//login uri pattern is auth
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")

	action := segs[2]
	provider := segs[3]

	switch action {
	case "login":
		log.Println("Handle login request by ", provider)

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Err when trying with provider ", provider, "-", err)
		}

		loginUrl, err := provider.GetBeginAuthURL(nil, objx.New(map[string]interface{}{
			"scope": "user",
		}))

		if err != nil {
			log.Fatalln("Err when trying to GetBeginAuthURL ", provider, "-", err)
		}

		log.Println("Redirect to login url ", loginUrl)
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Err when trying to get provider ", provider, "-", err)
		}

		log.Println(r.URL.RawQuery)

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))

		if err != nil {
			log.Fatalln("Err when trying to complte auth with provider ", provider, "-", err)
		}

		user, err := provider.GetUser(creds)

		if err != nil {
			log.Fatalln("Err when trying to get authed user ", provider, "-", err)
		}

		for k, v := range user.Data() {
			log.Println(k, ":", v)
		}

		usrName := user.Data()["login"]

		m := md5.New()
		io.WriteString(m, strings.ToLower(usrName.(string)))
		usrId := fmt.Sprintf("%x", m.Sum(nil))

		log.Println(usrId)
		authCookieValue := objx.New(map[string]interface{}{
			"usrid":      usrId,
			"name":       usrName,
			"avatar_url": user.Data()["avatar_url"],
			"email":      "test@mail.com",
		}).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})

		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		fmt.Fprintf(w, "Unknown request %s request %s", action, provider)
	}
}
