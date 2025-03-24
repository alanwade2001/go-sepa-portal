package main

import (
	"log"

	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-portal/data"

	utils "github.com/alanwade2001/go-sepa-utils"
)

func main() {
	persist := db.NewPersist()
	schemaName := utils.Getenv("DB_SCHEMA", "public")
	if err := persist.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + schemaName).Error; err != nil {
		log.Fatalf("cannot create schema: [%s], error:[%v]", schemaName, err)
		return
	}

	persist.DB.AutoMigrate(&data.Initiation{})
	log.Printf("created table: [%s]", "initiations")

}
