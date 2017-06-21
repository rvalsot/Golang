package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/satori/uuid"
)

var templat *template.Template

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	port := ":8070"
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(port, nil)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		sessionID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, cookie)
	}

	return cookie
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fileName string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fileName) {
		s += "|" + fileName
	}
	c.Value = s

	http.SetCookie(w, c)
	return c
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)

	// process sumbission
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("new_file")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// Copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// add filename to this user's cookie
		cookie = appendValue(w, cookie, fname)
	}
	xs := strings.Split(cookie.Value, "|")

	templat.ExecuteTemplate(w, "index.gohtml", xs)
}
