// webchat project main.go
package main

import (
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/objx"
	"github.com/stretchr/signature"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
	fpath    string
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(
			template.ParseFiles(
				filepath.Join(t.fpath, t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":9090", "The address of the web server.")
	var http_proxy = flag.String("http_proxy", "", "The http proxy")
	flag.Parse()
	if *http_proxy != "" {
		log.Println("Access notwork via http proxy ", *http_proxy)
		os.Setenv("HTTP_PROXY", *http_proxy)
	}

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(github.New("8779ebdf5b25ab55e6d2", "eed0cbf3b35887b7093b8a4e76f207ad5b505697",
		"http://localhost:9090/auth/callback/github"))

	r := newRoom(UseGravatar)
	http.Handle("/login", &templateHandler{fpath: "templates", filename: "login.html"})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/chat", MustAuth(&templateHandler{fpath: "templates", filename: "chat.html"}))
	http.Handle("/room", r)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	go r.run()

	log.Println("HTTP listens " + *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Println(err)
		log.Fatal("Unable to start Web server.")
	}
}
