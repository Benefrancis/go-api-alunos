package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Aluno struct {
	codigo      int64
	ra          string
	nome        string
	email       string
	turmaCodigo int64
}

var database *sql.DB

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/alunos")

	if err != nil {
		log.Println("Could not connect!")
	}

	database = db

	// Connect and check the server version
	var version string
	database.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	now := time.Now()

	fmt.Println("PT BR:", now.Format("02/01/2006"))

	http.HandleFunc("/", showAluno)

	http.ListenAndServe(":8999", nil)

}

func showAluno(w http.ResponseWriter, r *http.Request) {

	aluno := Aluno{}

	queryParam := "%" + r.URL.Path[1:] + "%"

	rows, err := database.Query("SELECT codigo, ra, nome, email, turma_Codigo FROM aluno WHERE nome LIKE ?", queryParam)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	html := "<html><head><title>Alunos</title></head><body><h1>Busca por: " + queryParam + "</h1><table border='1'><tr><th>CODIGO</th><th>RA</th><th>NOME</th><th>EMAIL</th></tr>"

	for rows.Next() {
		err := rows.Scan(&aluno.codigo, &aluno.ra, &aluno.nome, &aluno.email, &aluno.turmaCodigo)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d %s %s %s \n", aluno.codigo, aluno.ra, aluno.nome, aluno.email)

		html += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>%s</td></tr>", aluno.codigo, aluno.ra, aluno.nome, aluno.email)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	} else {

		html += "</table></body></html>"

		fmt.Fprintln(w, html)

	}

}
