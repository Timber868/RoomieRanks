package types

import "time"

// -- User types

//This is the type i am going to be using to read json data sent and parse payload
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type User struct {
	ID          int       `json:"ID"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	CreatedAt   time.Time `json:"createdAT"`
	HouseholdID int       `json:"householdID"`
}

//Interfaces are super easy to test so thats why
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

// -- Household types
type Household struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type HouseholdStore interface {
	GetHouseholdByID(id int) (*Household, error)
	CreateHousehold(Household) error
}

type RegisterHouseholdPayload struct {
	Name string `json:"name" validate:"required"`
}

// -- Chore types
type Chore struct {
	ID              int    `json:"ID"`
	Name            string `json:"name"`
	XP              int    `json:"xp"`
	Difficulty      int    `json:"difficulty"`
	TimeEstimate    int    `json:"timeEstimate"`
	CompletitonTime int    `json:"completionTime"`
}

type ChoreStore interface {
	GetChoreByID(id int) (*Chore, error)
	CreateChore(Chore) error
}

type RegisterChorePayload struct {
	Name            string `json:"name" validate:"required"`
	XP              int    `json:"xp" validate:"required"`
	Difficulty      int    `json:"difficulty" validate:"required"`
	TimeEstimate    int    `json:"timeEstimate" validate:"required"`
	CompletitonTime int    `json:"completionTime" validate:"required"`
}
