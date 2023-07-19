package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	DB_HOST := "localhost"
	DB_PORT := "5432"
	DB_DATABASE := "gc"
	DB_USERNAME := "postgres"
	DB_PASSWORD := "postgres"

	// Construindo a string de conexão com o banco de dados
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)

	// Conexão com o banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco de dados: %v", err)
	}

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("falha ao testar a conexão com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return db, nil
}
