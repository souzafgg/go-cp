package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere usuario no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}
	var u usuario
	if err = json.Unmarshal(corpoRequisicao, &u); err != nil {
		w.Write([]byte("Erro ao converter para struct"))
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco"))
		return
	}
	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(u.Nome, u.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}
	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso -> %v", idInserido)))
}
