package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var (
	dsn string
)

func TestMain(m *testing.M) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3306"
	}
	dsn = fmt.Sprintf("root:@tcp(127.1:%s)/", port)

	code := m.Run()
	os.Exit(code)
}

func BenchmarkGetISStats(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		getISStats(db)
		db.Close()
	}
}

func BenchmarkShowTableStatus(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		getTableStatus(db)
		db.Close()
	}
}
