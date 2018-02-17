package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	tableSchemaQuery = `
		SELECT SQL_NO_CACHE
		    TABLE_SCHEMA,
		    TABLE_NAME,
		    TABLE_TYPE,
		    ifnull(ENGINE, 'NONE') as ENGINE,
		    ifnull(VERSION, '0') as VERSION,
		    ifnull(ROW_FORMAT, 'NONE') as ROW_FORMAT,
		    ifnull(TABLE_ROWS, '0') as TABLE_ROWS,
		    ifnull(DATA_LENGTH, '0') as DATA_LENGTH,
		    ifnull(INDEX_LENGTH, '0') as INDEX_LENGTH,
		    ifnull(DATA_FREE, '0') as DATA_FREE,
		    ifnull(CREATE_OPTIONS, 'NONE') as CREATE_OPTIONS
		  FROM information_schema.tables
		  WHERE TABLE_SCHEMA = '%s'
`
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func main() {
	fmt.Println("vim-go")
	db, err := sql.Open("mysql", "root:@tcp(127.1:3306)/")
	if err != nil {
		panic(err)
	}
	getISStats(db)
}

func getISStats(db *sql.DB) {
	dbList := showDatabases(db)

	for _, database := range dbList {
		tableSchemaRows, err := db.Query(fmt.Sprintf(tableSchemaQuery, database))
		if err != nil {
			continue
		}
		defer tableSchemaRows.Close()

		var (
			tableSchema   string
			tableName     string
			tableType     string
			engine        string
			version       uint64
			rowFormat     string
			tableRows     uint64
			dataLength    uint64
			indexLength   uint64
			dataFree      uint64
			createOptions string
		)

		for tableSchemaRows.Next() {
			err = tableSchemaRows.Scan(
				&tableSchema,
				&tableName,
				&tableType,
				&engine,
				&version,
				&rowFormat,
				&tableRows,
				&dataLength,
				&indexLength,
				&dataFree,
				&createOptions,
			)
		}
	}

}

func showDatabases(db *sql.DB) []string {
	databases := []string{}
	rows, err := db.Query(fmt.Sprintf("SHOW DATABASES"))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var database string
		err := rows.Scan(&database)
		if err != nil {
			panic(err)
		}
		databases = append(databases, database)
	}
	return databases
}

func getTableStatus(db *sql.DB) {
	dbList := showDatabases(db)
	for _, dbname := range dbList {
		var (
			name            sql.NullString
			engine          sql.NullString
			version         sql.NullFloat64
			row_format      sql.NullString
			rows            sql.NullInt64
			avg_row_length  sql.NullInt64
			data_length     sql.NullInt64
			max_data_length sql.NullInt64
			index_length    sql.NullInt64
			data_free       sql.NullInt64
			auto_increment  sql.NullInt64
			create_time     NullTime
			update_time     NullTime
			check_time      NullTime
			collation       sql.NullString
			checksum        sql.NullInt64
			create_options  sql.NullString
			comment         sql.NullString
		)
		srows, err := db.Query(fmt.Sprintf("SHOW TABLE STATUS FROM %s LIKE '%%'", dbname))
		if err != nil {
			panic(err)
		}
		for srows.Next() {
			srows.Scan(
				&name,
				&engine,
				&version,
				&row_format,
				&rows,
				&avg_row_length,
				&data_length,
				&max_data_length,
				&index_length,
				&data_free,
				&auto_increment,
				&create_time,
				&update_time,
				&check_time,
				&collation,
				&checksum,
				&create_options,
				&comment,
			)
		}

	}
}
