package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/kshmatov/dashboard/types"
)

var dbConnection *sql.DB

func Init(config types.Database) error {
	dbC, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			config.Username, config.Password,config.Host, config.Port, config.Schema))
	if err != nil {
		dbConnection = nil
		return err
	}
	dbConnection = dbC
	return nil
}
