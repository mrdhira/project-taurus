package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrdhira/project-taurus/pkg/sqlExt"
)

type Config struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type mysqlExt struct {
	dbConn         *sqlx.DB
	dbReadOnlyConn *sqlx.DB
}

func New(config Config) (sqlExt.ISqlExt, error) {
	var (
		dbConn         *sqlx.DB
		dbReadOnlyConn *sqlx.DB
		err            error
	)

	dbConn, err = sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		),
	)
	if err != nil {
		return nil, err
	}

	// Check connection
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	// Temporary use the same connection
	dbReadOnlyConn = dbConn

	return &mysqlExt{
		dbConn:         dbConn,
		dbReadOnlyConn: dbReadOnlyConn,
	}, nil
}
