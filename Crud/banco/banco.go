package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

//Conectar abre a conexão com o DB
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	fmt.Println("Conexão Aberta!")

	return db, nil
}
