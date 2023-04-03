package database

import (
	"database/sql"
	"time"
	"user-management-api/internal/utils"
)

func PrepareDB(db *sql.DB) error {
	var err error
	for i := 0; i < 60; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		return err
	}

	if _, err = db.Exec("CREATE TABLE IF NOT EXISTS `user` ( `id` int NOT NULL AUTO_INCREMENT, `username` varchar(50) NOT NULL, `email` varchar(100) NOT NULL, `password` varchar(255) NOT NULL, `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`id`), UNIQUE KEY `username_UNIQUE` (`username`), UNIQUE KEY `email_UNIQUE` (`email`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"); err != nil {
		return err
	}

	var count int
	if err = db.QueryRow("SELECT COUNT(id) FROM user").Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		if err = PrepareDefaultUser(db); err != nil {
			return err
		}
	}

	return nil
}

func PrepareDefaultUser(db *sql.DB) error {
	username := "administrator"
	password := "admin"
	email := "admin@management.api"

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	password = hashedPassword

	_, err = db.Exec("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", username, email, password)
	if err != nil {
		return err
	}
	return nil
}
