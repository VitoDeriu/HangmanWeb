package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

const port = ":8080"

var conteur int

type User struct {
	Lastname  string
	Firstname string
	Birthday  string
	Sexe      string
}

var data = User{}

func main() {
	temp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR !>", err.Error()))
		return
	}

	type Etudiant struct {
		Prenom string
		Nom    string
		Age    int
		Sexe   bool
	}
	type InfoPromo struct {
		Nom     string
		Filiere string
		Niveau  int
		Nombre  int
		Edu     []Etudiant
	}
	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		User := []Etudiant{{"Cyril", "RODRIGUES", 22, true}, {"Kheir-Eddine", "MEDERREG", 22, false}, {"Alan", "PHILIPIERT", 26, true}}
		data := InfoPromo{"Mentor'ac", "Informatique", 5, 3, User}
		temp.ExecuteTemplate(w, "promo", data)
	})

	type Display struct {
		Cpt  int
		Pair bool
	}

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		conteur++
		var ispair bool
		if conteur%2 == 0 {
			ispair = true
		} else {
			ispair = false
		}
		data := Display{conteur, ispair}
		temp.ExecuteTemplate(w, "change", data)
	})

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
        temp.ExecuteTemplate(w, "init", nil)
    })

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		data = User{
			r.FormValue("nom"),
			r.FormValue("prenom"),
			r.FormValue("date"),
			r.FormValue("sexe")}
        http.Redirect(w, r, "/user/display", http.StatusMovedPermanently)
    })

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
        temp.ExecuteTemplate(w, "display", data)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
