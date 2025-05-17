package cons

import "errors"

var (
	NO_ROWS_AFFECTED error = errors.New("sql: no rows affected")
)

const (
	DEV  = "development"
	STAG = "staging"
	PROD = "production"
	TEST = "test"

	API = "/api/v1"

	EMPTY           = ""
	Nil             = iota
	InvalidUUID     = "00000000-0000-0000-0000-000000000000"
	DEFAULT_ERR_MSG = "API is busy please try again later!"

	DATE_TIME_FORMAT = "2006-01-02 15:04:05"
)

const (
	POSTGRES = "postgres"
)
