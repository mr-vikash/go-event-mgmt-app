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

func UpdateEvent(id int, name, location string) (*models.Event, error) {

	query := "UPDATE events SET name = ?, location = ? WHERE id = ?"
	result, err := database.DB.Exec(query, name, location, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("event not found")
	}

	var event models.Event
	queryRecord := "SELECT id, name, location FROM events WHERE id = ?"

	err = database.DB.QueryRow(queryRecord, id).
		Scan(&event.ID, &event.Name, &event.Location)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func DeleteEvent(id int) (string, error) {
	Query := "delete from events where id = ?"
	result, err := database.DB.Exec(Query, id)

	if err != nil {
		return "Some went wrong", err
	}

	rowAffected, err := result.RowsAffected()

	if rowAffected == 0 {
		return "Record not deleted", err
	}

	return "record deleted", nil
}

func GetAllEvents() (*[]models.Event, error) {
	query := "Select * from events"
	events := []models.Event{}
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Location)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return &events, nil
}
