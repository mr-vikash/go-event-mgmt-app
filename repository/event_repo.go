package repository

import (
	"fmt"
	"log"

	"go-event-mgmt-app/database"
	"go-event-mgmt-app/models"
)

func InsertEvent(name string, location string) int {
	query := "insert into events(name, location) values(?, ?)"
	result, err := database.DB.Exec(query, name, location)
	if err != nil {
		log.Fatal(err)
	}

	eventID, _ := result.LastInsertId()
	fmt.Println("Inserted ID:", eventID)
	return int(eventID)
}

func UpdateEvent(id int, name, location string) (models.Event, error) {
	query := "UPDATE events SET name = ?, location = ? WHERE id = ?"
	result, err := database.DB.Exec(query, name, location, id)

	if err != nil {
		return models.Event{}, err
	}

	rowAffected, _ := result.RowsAffected()
	if rowAffected == 0 {
		return models.Event{}, fmt.Errorf("Event not updated")
	}

	var event models.Event
	queryRecord := "select id, name, location where id = ?"
	err = database.DB.QueryRow(queryRecord).Scan(&event.ID, &event.Name, &event.Location)

	if err != nil {
		return models.Event{}, err
	}

	return event, nil

}
