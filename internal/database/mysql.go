package database

// internal/database/database.go

import (
	"fmt"
	"log"
	"time"

	"internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	for {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
		)

		db, err = sqlx.Connect("mysql", dsn)
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
