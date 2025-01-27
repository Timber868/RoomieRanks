package household

import (
	"database/sql"

	"github.com/Timber868/roomieranks/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetHouseholdByID(id int) (*types.Household, error) {
	rows, err := s.db.Query("SELECT * FROM households WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	//Instantiate our household we will modify it later
	h := new(types.Household)

	for rows.Next() {
		h, err = scanRowIntoHousehold(rows)

		//Handle any errors that might have occured
		if err != nil {
			return nil, err
		}
	}

	//If ID is 0 then we know we didnt find anything
	if h.ID == 0 {
		return nil, nil
	}

	return h, nil
}

func (s *Store) CreateHousehold(h types.Household) (int, error) {
	// Insert the household into the database and get the result
	result, err := s.db.Exec("INSERT INTO households (name) VALUES (?)", h.Name)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted household
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func scanRowIntoHousehold(rows *sql.Rows) (*types.Household, error) {
	household := new(types.Household)

	//Directly map the data to our household via pointers
	if err := rows.Scan(
		&household.ID,
		&household.Name,
	); err != nil {
		//Hadnle errors if any were found
		return nil, err
	}

	//Only one houshold per person
	return household, nil
}
