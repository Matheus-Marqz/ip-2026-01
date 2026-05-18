package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type paciente struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Fone        string `json:"fone"`
	DataNasc    string `json:"data_nasc"`
	Sexo        string `json:"sexo"`
	NumConvenio string `json:"num_convenio"`
}

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("postgres", "postgresql://postgres:1234@localhost:5433/crud?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database.")
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var p paciente

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.QueryRow(
		"INSERT INTO pacientes (name, fone, data_nasc, sexo, num_convenio) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		p.Name,
		p.Fone,
		p.DataNasc,
		p.Sexo,
		p.NumConvenio,
	).Scan(&p.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id != "" {
		ReadOne(w, id)
		return
	}

	rows, err := db.Query("SELECT id, name, fone, data_nasc, sexo, num_convenio FROM pacientes ORDER BY id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	pacientes := []paciente{}

	for rows.Next() {
		var p paciente

		err = rows.Scan(&p.ID, &p.Name, &p.Fone, &p.DataNasc, &p.Sexo, &p.NumConvenio)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pacientes = append(pacientes, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pacientes)
}

func ReadOne(w http.ResponseWriter, id string) {
	var p paciente

	err := db.QueryRow(
		"SELECT id, name, fone, data_nasc, sexo, num_convenio FROM pacientes WHERE id = $1",
		id,
	).Scan(&p.ID, &p.Name, &p.Fone, &p.DataNasc, &p.Sexo, &p.NumConvenio)

	if err == sql.ErrNoRows {
		http.Error(w, "Paciente nao encontrado", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var p paciente

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"UPDATE pacientes SET name = $1, fone = $2, data_nasc = $3, sexo = $4, num_convenio = $5 WHERE id = $6",
		p.Name,
		p.Fone,
		p.DataNasc,
		p.Sexo,
		p.NumConvenio,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM pacientes WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Paciente deletado com sucesso")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/pacientes/create", Create)
	http.HandleFunc("/pacientes/read", Read)
	http.HandleFunc("/pacientes/update", Update)
	http.HandleFunc("/pacientes/delete", Delete)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
