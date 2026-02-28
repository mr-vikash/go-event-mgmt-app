package database

import (
	"database/sql"
	"fmt"
	"log"

	"go-event-mgmt-app/models"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:3311)/event_mgmt_db"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB Open error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB Connection failed:", err)
	}

	fmt.Println("✅ MySQL Connected Successfully")
}

func InsertEvent(name string, location string) int {
	query := "insert into events(name, location) values(?, ?)"
	result, err := DB.Exec(query, name, location)
	if err != nil {
		log.Fatal(err)
	}

	eventID, _ := result.LastInsertId()
	fmt.Println("Inserted ID:", eventID)
	return int(eventID)
}

func UpdateEvent(name string, location string, id int) (models.Event, error) {
	query := "update events set name = ? location = ?  where id = ?"
	result, err := DB.Exec(query, name, location, id)

	if err != nil {
		return models.Event{}, err
	}

	rowAffected, _ := result.RowsAffected()
	if rowAffected == 0 {
		return models.Event{}, fmt.Errorf("Event not updated")
	}

	var event models.Event
	queryRecord := "select id, name, location where id = ?"
	err = DB.QueryRow(queryRecord).Scan(&event.ID, &event.Name, &event.Location)

	if err != nil {
		return models.Event{}, err
	}

	return event, nil

}
