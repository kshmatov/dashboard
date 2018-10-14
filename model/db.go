package model

import (
	"database/sql"
	"errors"

	"github.com/kshmatov/dashboard/types"
)

var dbConnection *sql.Conn

func Init(config types.Database) error {
	return errors.New("Not implemented yet")
}