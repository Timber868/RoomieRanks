package types

// -- User types

//This is the type i am going to be using to read json data sent and parse payload
type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=130"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
	Name     string `json:"name" validate:"required"`
}

type User struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	HouseholdID int    `json:"household_id"`
	Title       string `json:"title"`
	Level       int    `json:"level"`
}

//Interfaces are super easy to test so thats why
type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
	ModifyUser(User) error
	ChangeTitle(username string, title string) error
	LevelUp(username string) error
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
