package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/motorheads/shop_product_service/config"
	"github.com/motorheads/shop_product_service/routes"
)

// test git
var err error

func main() {

	config.DB, err = sql.Open("postgres", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Error while connecting to the database")
		panic(err)
	}

	defer config.DB.Close()

	router := routes.New()
	router.Run()

}
