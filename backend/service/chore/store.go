package chore

import (
	"database/sql"
	"fmt"

	"github.com/Timber868/roomieranks/types"
)

// Store is the struct that will hold the database connection
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetChoreByID will return a chore by its ID
func (s *Store) GetChoreByID(id int) (*types.Chore, error) {
	// Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM chore WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	c := new(types.Chore)

	for rows.Next() {
		c, err = scanRowIntoChore(rows)

		// Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}

	// If the id is an empty string the chore was not found
	if c.ID == 0 {
		return nil, fmt.Errorf("chore not found")
	}

	return c, nil
}

func scanRowIntoChore(rows *sql.Rows) (*types.Chore, error) {
	chore := new(types.Chore)

	// Directly map the data to our chore via pointers
	err := rows.Scan(
		&chore.ID,
		&chore.Name,
		&chore.Difficulty,
		&chore.TimeEstimate,
		&chore.CompletitonTime,
		&chore.HouseholdID,
	)

	return chore, err
}

// CreateChore will create a new chore
func (s *Store) CreateChore(c types.Chore) error {
	_, err := s.db.Exec("INSERT INTO chore (name, difficulty, timeEstimate, completionTime, householdID) VALUES (?, ?, ?, ?, ?)", c.Name, c.Difficulty, c.TimeEstimate, c.CompletitonTime, c.HouseholdID)

	if err != nil {
		return err
	}

	return nil
}

// GetChoreByHouseholdID will return all chore for a household
func (s *Store) GetChoreByHouseholdID(id int) ([]*types.Chore, error) {
	// Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM chore WHERE householdID = ?", id)

	if err != nil {
		return nil, err
	}

	var chore []*types.Chore

	for rows.Next() {
		c, err := scanRowIntoChore(rows)

		// Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}

		chore = append(chore, c)
	}

	return chore, nil
}

// UpdateChore will update a chore
func (s *Store) UpdateChore(c types.Chore) error {
	_, err := s.db.Exec("UPDATE chore SET name = ?, difficulty = ?, timeEstimate = ?, completionTime = ?, householdID = ? WHERE id = ?", c.Name, c.Difficulty, c.TimeEstimate, c.CompletitonTime, c.HouseholdID, c.ID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteChore will delete a chore
func (s *Store) DeleteChore(id int) error {
	_, err := s.db.Exec("DELETE FROM chore WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
