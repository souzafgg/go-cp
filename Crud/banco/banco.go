package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// conectar abre a con com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/szadb?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
