package database

// internal/database/database.go

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"sewapoint/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
		)

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Gagal koneksi ke database: %v", err)
			sleepDuration := 1 // Detik
			log.Printf("Mencoba kembali dalam %d detik...", sleepDuration)
			time.Sleep(time.Duration(sleepDuration) * time.Second)
			continue
		} else {
			log.Println("Berhasil terkoneksi ke database")
			break
		}
	}

	return db, nil
}
