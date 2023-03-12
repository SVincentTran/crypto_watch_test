package helpers

import (
	"cryptowatch/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDBConnection(cfg *config.PostgresqlConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}

func GetUser(db *sql.DB, userId string) (bool, error) {
	rows, err := db.Query("SELECT user_id FROM users WHERE user_id = $1", userId)
	defer rows.Close()
	if err != nil {
		log.Printf("Error while querying for user %v", err)
		return false, err
	}

	var queriedId string
	var userIds []string
	for rows.Next() {
		if err := rows.Scan(&queriedId); err != nil {
			log.Printf("Error while scanning user %v", err)
			return false, err
		}
		userIds = append(userIds, userId)
	}

	if len(userIds) > 0 {
		return true, nil
	}

	return false, nil
}

func AddNewUser(db *sql.DB, userId string) error {
	_, err := db.Exec("INSERT INTO users(user_id) VALUES($1)", userId)

	if err != nil {
		log.Printf("Error while inserting new record into databse %v", err)
		return err
	}

	return nil
}
