package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const port = ":8080"

func main() {
	temp, err := template.ParseGlob("./template/promo.html")
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
		User := []Etudiant{{"Cyril", "RODRIGUES", 22, true}, {"Kheir-Eddine", "MEDERREG", 22, true}, {"Alan", "PHILIPIERT", 26, true}}
		data := InfoPromo{"Mentor'ac", "Informatique", 5, 3, User}
		temp.ExecuteTemplate(w, "promo", data)
	})

	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
