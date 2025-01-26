package chore_instance

import (
	"database/sql"
	"fmt"

	"github.com/Timber868/roomieranks/types"

	//Import Chore/store.go
	"github.com/Timber868/roomieranks/service/chore"
)

//This is our file to handle repositories

type Store struct {
	db         *sql.DB
	choreStore *chore.Store
}

func NewStore(db *sql.DB, choreStore *chore.Store) *Store {
	return &Store{
		db:         db,
		choreStore: choreStore,
	}
}

func scanRowIntoChoreInstance(rows *sql.Rows) (*types.ChoreInstance, error) {
	choreInstance := new(types.ChoreInstance)

	//Directly map the data to our user via pointers
	err := rows.Scan(
		&choreInstance.ID,
		&choreInstance.ChoreID,
		&choreInstance.Username,
		&choreInstance.DueDate,
		&choreInstance.Completed,
	)

	if err != nil {
		return nil, err
	}

	return choreInstance, nil
}

func (s *Store) GetChoreInstanceByUsername(username string) ([]*types.ChoreInstance, error) {
	// Query the database
	rows, err := s.db.Query("SELECT * FROM chore_instance")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var choreInstances []*types.ChoreInstance

	for rows.Next() {
		// Use the existing scanRowIntoChoreInstance function
		choreInstance, err := scanRowIntoChoreInstance(rows)
		if err != nil {
			return nil, err
		}

		// Filter by username
		if choreInstance.Username == username {
			choreInstances = append(choreInstances, choreInstance)
		}
	}

	return choreInstances, nil
}

func (s *Store) GetChoreInstanceByChoreID(choreId int) ([]*types.ChoreInstance, error) {
	// Query the database
	rows, err := s.db.Query("SELECT * FROM chore_instance")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var choreInstances []*types.ChoreInstance

	for rows.Next() {
		// Use the existing scanRowIntoChoreInstance function
		choreInstance, err := scanRowIntoChoreInstance(rows)
		if err != nil {
			return nil, err
		}

		// Filter by username
		if choreInstance.ChoreID == choreId {
			choreInstances = append(choreInstances, choreInstance)
		}
	}

	return choreInstances, nil
}

func (s *Store) CreateChoreInstance(choreInstance types.ChoreInstance) error {
	_, err := s.db.Exec("INSERT INTO chore_instance (user_username, chore_Id, completed, due_date) VALUES (?, ?, ?, ?)", choreInstance.Username, choreInstance.ChoreID, false, choreInstance.DueDate)

	return err
}

func (s *Store) GetChoreInstanceByID(choreInstanceId int) (*types.ChoreInstance, error) {
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
	updateQuery := fmt.Sprintf("UPDATE chore_instance SET username = %s WHERE id = %d", username, choreInstanceId)

	_, err := s.db.Exec(updateQuery)
	if err != nil {
		return fmt.Errorf("failed to update row: %w", err)
	}

	return nil
}

func (s *Store) GetChoreXP(choreInstanceId int) (int, error) {
	// Compute the XP from the chore instance based on difficulty and time estimate (Max XP = 20) (Higher difficulty and higher time estimate = more XP)
	choreInstance, err := s.GetChoreInstanceByID(choreInstanceId)
	if err != nil {
		return 0, err
	}
	chore, err := s.choreStore.GetChoreByID(choreInstance.ChoreID)
	if err != nil {
		return 0, err
	}
	xp := chore.Difficulty * chore.TimeEstimate / 10

	if xp > 20 {
		xp = 20
	}

	return xp, nil
}

func (s *Store) CompleteChore(choreInstanceId int) error {
	updateQuery := fmt.Sprintf("UPDATE chore_instance SET completed = %t WHERE id = %d", true, choreInstanceId)

	_, err := s.db.Exec(updateQuery)
	if err != nil {
		return fmt.Errorf("failed to update row: %w", err)
	}
	//TODO - Add XP to user
	return nil
}
