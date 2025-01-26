package user

import (
	"database/sql"
	"fmt"

	"github.com/Timber868/roomieranks/types"
)

//This is our file to handle repositories

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoChoreInstance(rows *sql.Rows) (*types.User, error) {
	choreInstance := new(types.choreInstance)

	//Directly map the data to our user via pointers
	err := rows.Scan(
		&choreInstance.ID,
		&choreInstance.Username,
		&choreInstance.ChoreID,
		&choreInstance.Completed,
		&choreInstance.DueDate,
	)

	if err != nil {
		return nil, err
	}

	return choreInstance, nil
}

func (s *Store) GetChoreInstanceByUsername(username string) ([]*ChoreInstance, error) {
	//Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM chore_instance")

	if err != nil {
		return nil, err
	}

	c := []new(types.ChoreInstance)

	for rows.Next() {
		c, err = append(c, rows)

		//Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}

	var choreInstances []ChoreInstance
	for index, value := range c {
        if value.Username == username {
			choreInstances = append(choreInstances, value)
		}
    }

	return &choreInstances, nil
}


func (s *Store) GetChoreInstanceByChoreID(choreId int) ([]*ChoreInstance, error) {
	//Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM chore_instance")
	c := []new(types.ChoreInstance)

	for rows.Next() {
		c, err = append(c, rows)

		//Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}
	
	var choreInstances []ChoreInstance
	for index, value := range c {
        if value.ChoreID == choreId {
			choreInstances = append(choreInstances, value)
		}
    }

	return &choreInstances, nil
}

func (s *Store) CreateChoreInstance(choreInstance types.ChoreInstance) error {
	_, err := s.db.Exec("INSERT INTO chore_instance (username, choreId, completed, dueDate) VALUES (?, ?, ?, ?)", choreInstance.Username, choreInstance.ChoreID, false, choreInstance.DueDate)

	return err
}

func (s *Store) GetChoreInstanceByID(choreInstanceId int) (ChoreInstance, error) {
	//Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM chore_instance WHERE id = ?", choreInstanceId)

	if err != nil {
		return nil, err
	}

	c := new(types.ChoreInstance)

	for rows.Next() {
		c, err = scanRowIntoChoreInstance(rows)

		//Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (s *Store) AssignChoreInstance(choreInstanceId int, username string) error {
	updateQuery := fmt.Sprintf("UPDATE chore_instance SET username = username WHERE id = choreInstanceId", tableName, columnName)

	_, err = db.Exec(updateQuery, newValue, id)
	if err != nil {
		return fmt.Errorf("failed to update row: %w", err)
	}

	return nil
}

GetChoreXP(choreInstanceId int) (int, error) {
	//Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM chore_instance WHERE id = ?", choreInstanceId)

	if err != nil {
		return nil, err
	}

	c := new(types.ChoreInstance)

	for rows.Next() {
		c, err = scanRowIntoChoreInstance(rows)

		//Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}

	//Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM chore WHERE id = ?", c.ChoreID)

	if err != nil {
		return nil, err
	}

	// YoU NEED TO UPDATE THIS GARRETT AND CALL UR METHOD TO SCAN INTO A CHORE
	c := new(types.Chore)

	for rows.Next() {
		c, err = scanRowIntoChore(rows)

		//Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}
}

// TO BE COMPLETED
CompleteChore(choreInstanceId int) error {

}