package main

import (
	"fmt"
	"os"

	"github.com/Necoo33/neormgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Env'ler yüklenirken Şu Hata Alındı: %s", err)
	}

	db := neormgo.Neorm{}

	db_conn_url := os.Getenv("DB_CONN_URL")

	db, err = db.Connect(db_conn_url)

	if err != nil {
		fmt.Printf("Daniş Mahfazasına bağlanırken şu hata alındı: %s", err)
	}

	schemaName := os.Getenv("SCHEMA_NAME")

	db.CreateSchema(schemaName)
	db.IfNotExist()
	db.Finish()
	_, err = db.Execute()

	if err != nil {
		fmt.Printf("Şema kurarken şu hata alındı: %s", err)
	}

	db.Use(schemaName)

	db.CreateTable("users").IfNotExist()

	db.AddColumn("id")
	db.Type("INT")
	db.PrimaryKey()
	db.AutoIncrement()

	db.AddColumn("name")
	db.Type("VARCHAR(30)")
	db.NotNull()

	db.AddColumn("email")
	db.Type("VARCHAR(60)")
	db.Unique()
	db.NotNull()

	db.AddColumn("password")
	db.Type("VARCHAR(30)")
	db.NotNull()

	db.AddColumn("role")
	db.Type("VARCHAR(15)")
	db.NotNull()
	db.Default("user")

	db.AddColumn("pic")
	db.Type("JSON")
	db.Null()

	db.Finish()
	err = db.QueryDrop()

	if err != nil {
		fmt.Printf("\nAn error occured when we creating users table: %s\n", err)
	}

	db.CreateTable("articles").IfNotExist()

	db.AddColumn("aid")
	db.Type("INT")
	db.PrimaryKey()
	db.AutoIncrement()

	db.AddColumn("name")
	db.Type("VARCHAR(30)")
	db.NotNull()

	db.AddColumn("description")
	db.Type("TEXT")
	db.NotNull()

	db.AddColumn("pic")
	db.Type("JSON")
	db.Null()

	db.AddColumn("created_at")
	db.Type("DATETIME")
	db.Default("CURRENT_TIMESTAMP")

	db.AddColumn("published")
	db.Type("INT")
	db.NotNull()

	db.AddColumn("id")
	db.Type("INT")

	db.Finish()
	err = db.QueryDrop()

	if err != nil {
		fmt.Printf("\nAn error occured when we creating users table: %s\n", err)
	}
}
