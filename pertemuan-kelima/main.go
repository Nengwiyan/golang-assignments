package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Member struct {
	Email     string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var PORT = ":8000"

var members = []Member{
	{Email: "nengwiyan@tes.com", Nama: "Neneng", Alamat: "Jl. ABC", Pekerjaan: "Backend", Alasan: "Alasan Neneng"},
	{Email: "wiyanti@tes.com", Nama: "Wiyanti", Alamat: "Jl.XYZ", Pekerjaan: "Backend", Alasan: "Alasan Wiyanti"},
	{Email: "rika@tes.com", Nama: "Rika", Alamat: "Jl.Lorem", Pekerjaan: "Backend", Alasan: "Alasan Rika"},
	{Email: "karin@tes.com", Nama: "Riana", Alamat: "Jl.Ipsum", Pekerjaan: "Backend", Alasan: "Alasan Riana"},
	{Email: "yani@tes.com", Nama: "Yani", Alamat: "Jl.BCA", Pekerjaan: "Backend", Alasan: "Alasan Yani"},
}

func main() {
	http.HandleFunc("/", routeGet)
	http.HandleFunc("/process", routePost)

	fmt.Println("server started at", PORT)
	http.ListenAndServe(PORT, nil)
}

func routeGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var template = template.Must(template.New("form").ParseFiles("templatelogin.html"))
		var err = template.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func routePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var email = r.FormValue("email")
		var mmb Member

		for _, user := range members {
			if user.Email == email {
				mmb = user
				break
			}
		}

		if mmb.Email == "" {
			notFound(w)
		} else {
			renderData(w, mmb)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func renderData(w http.ResponseWriter, data Member) {
	template := template.Must(template.New("result").ParseFiles("templatedata.html"))

	if err := template.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func notFound(w http.ResponseWriter) {
	template := template.Must(template.New("fail").ParseFiles("notfound.html"))

	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
