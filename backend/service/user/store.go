package user

import (
	"database/sql"
	"fmt"

	"github.com/Timber868/roomieranks/service/collectible"
	"github.com/Timber868/roomieranks/types"
)

// This is our file to handle repositories
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	//Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM users WHERE username = ?", username)

	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowIntoUser(rows)

		//Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}

	//If the username is an empty string the user was not found
	if u.Username == "" {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	//Directly map the data to our user via pointers
	err := rows.Scan(
		&user.Username,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.HouseholdID,
		&user.Title,
		&user.Level,
		&user.XP,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Methods that my user store interface calls

func (s *Store) CreateUser(user types.User) error {
	//Insert the user into the database
	_, err := s.db.Exec("INSERT INTO users (username, name, email, password, household_id, title, level, xp) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Username, user.Name, user.Email, user.Password, user.HouseholdID, user.Title, user.Level, user.XP)

	return err
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	//Just use basic sql to communicate with the database
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowIntoUser(rows)

		//Make sure everything is good for data validation
		if err != nil {
			return nil, err
		}
	}

	//If id equals 0 the user was not found
	if u.Email == "" {
		return nil, fmt.Errorf("user with email %s not found", email)
	}

	return u, nil
}

func (s *Store) ModifyUser(user types.User) error {
	//To be implemented
	return nil
}

func (s *Store) ChangeTitle(username string, title string) error {
	u, err := s.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	//Update the title
	u.Title = title

	//Update the user in the database
	_, err = s.db.Exec("UPDATE users SET title = ? WHERE username = ?", title, username)
	if err != nil {
		return err
	}

	return nil
}

// Constant that defines the amount of XP needed to level up
const LevelUpXP = 20

func (s *Store) AddXP(username string, xp int) error {
	u, err := s.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	//Increment the level
	u.XP += xp
	var levelUps int

	//Check if the user has enough XP to level up
	if u.XP >= LevelUpXP {
		//Could level up multiple times
		levelUps = u.XP / LevelUpXP
		//Level up the user
		u.Level += levelUps
		u.XP = u.XP - LevelUpXP*levelUps
	}

	fmt.Println(levelUps)
	if levelUps > 0 {
		s.handleLevelUp(*u)
	}

	//Update the user XP in the database
	_, err = s.db.Exec("UPDATE users SET xp = ?, level = ? WHERE username = ?", u.XP, u.Level, username)
	if err != nil {
		return err
	}

	//Update the user level in the database
	_, err = s.db.Exec("UPDATE users SET level = ? WHERE username = ?", u.Level, username)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ChangeHousingID(username string, householdID int) error {
	u, err := s.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	//Update the household id
	u.HouseholdID = householdID

	//Update the user in the database
	_, err = s.db.Exec("UPDATE users SET household_id = ? WHERE username = ?", householdID, username)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) handleLevelUp(u types.User) {
	//To be implemented
	fmt.Println("User leveled up!")
	collectibleStore := collectible.NewStore(s.db)
	collectibleStore.CreateCollectible(u.Username)
}

func (s *Store) GetCollectiblesByUsername(username string) ([]types.Collectible, error) {
	rows, err := s.db.Query("SELECT * FROM collectible WHERE user_username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collectibles []types.Collectible

	for rows.Next() {
		collectible := new(types.Collectible)
		err := rows.Scan(
			&collectible.ID,
			&collectible.Name,
			&collectible.Rarity,
			&collectible.Type,
			&collectible.ImageURL,
			&collectible.UserUsername,
		)
		if err != nil {
			return nil, err
		}
		collectibles = append(collectibles, *collectible)
	}

	return collectibles, nil
}
