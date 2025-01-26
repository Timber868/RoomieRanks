package types

import "time"

// -- User types

// This is the type i am going to be using to read json data sent and parse payload
type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=130"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
	Name     string `json:"name" validate:"required"`
}

type LoginUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
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

// Interfaces are super easy to test so thats why
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
	Difficulty      int    `json:"difficulty"`
	TimeEstimate    int    `json:"timeEstimate"`
	CompletitonTime int    `json:"completionTime"`
	HouseholdID     int    `json:"householdID"`
}

type ChoreStore interface {
	GetChoreByID(id int) (*Chore, error)
	CreateChore(Chore) error
	GetChoreByHouseholdID(id int) ([]*Chore, error)
	UpdateChore(Chore) error
	DeleteChore(id int) error
}

type RegisterChorePayload struct {
	Name            string `json:"name" validate:"required"`
	Difficulty      int    `json:"difficulty" validate:"required"`
	TimeEstimate    int    `json:"timeEstimate" validate:"required"`
	CompletitonTime int    `json:"completionTime" validate:"required"`
	HouseholdID     int    `json:"householdID" validate:"required"`
}

// Chore Instance type
type ChoreInstance struct {
	ID        int       `json:"ID"`
	Username  string    `json:"username"`
	ChoreID   int       `json:"choreID"`
	Completed bool      `json:"completed"`
	DueDate   time.Time `json:"dueDate"`
}

type ChoreInstanceStore interface {
	GetChoreInstanceByUsername(username string) ([]*ChoreInstance, error)
	GetChoreInstanceByChoreID(choreId int) ([]*ChoreInstance, error)
	CreateChoreInstance(ChoreInstance) error
	GetChoreInstanceByID(choreInstanceId int) (*ChoreInstance, error)
	AssignChoreInstance(choreInstanceId int, username string) error
	GetChoreXP(choreInstanceId int) (int, error)
	CompleteChore(choreInstanceId int) error
}

type RegisterChoreInstancePayload struct {
	Username string `json:"username" validate:"required"`
	ChoreID  int    `json:"choreID" validate:"required"`
}

// -- Bill Types

type Bill struct {
	ID            int       `json:"ID"`
	PurchaseDate  time.Time `json:"purchaseDate"`
	PayerUsername string    `json:"payerUsername"`
}

type BillStore interface {
	GetBillByID(id int) (*Bill, error)
	CreateBill(Bill) error
}

type RegisterBillPayload struct {
	PurchaseDate  time.Time `json:"purchaseDate" validate:"required"`
	PayerUsername string    `json:"payerUsername" validate:"required"`
}

// -- Debt Types

type Debt struct {
	BillID       int     `json:"billID"`
	UserUsername string  `json:"userUsername"`
	Amount       float64 `json:"amount"`
}

type DebtStore interface {
	GetDebtByBillID(id int) ([]*Debt, error)
	GetDebtByUserUsername(username string) ([]*Debt, error)
	CreateDebt(Debt) error
}

type RegisterDebtPayload struct {
	BillID       int     `json:"billID" validate:"required"`
	UserUsername string  `json:"userUsername" validate:"required"`
	Amount       float64 `json:"amount" validate:"required"`
}

// -- Collectible Types

type Collectible struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	IsLegendary  bool   `json:"isLegendary"`
	UserUsername string `json:"userUsername"`
	ImageURL     string `json:"imageURL"`
	Evolution    int    `json:"evolution"`
}

type CollectibleStore interface {
	GetCollectibleByID(id int) (*Collectible, error)
	CreateCollectible(Collectible) error
}

type RegisterCollectiblePayload struct {
	Name         string `json:"name" validate:"required"`
	IsLegendary  bool   `json:"isLegendary" validate:"required"`
	UserUsername string `json:"userUsername" validate:"required"`
	ImageURL     string `json:"imageURL" validate:"required"`
	Evolution    int    `json:"evolution" validate:"required"`
}

// -- TradeRequest Types

type TradeRequest struct {
	ID int `json:"ID"`
}

type TradeRequestStore interface {
	GetTradeRequestByID(id int) (*TradeRequest, error)
	CreateTradeRequest(TradeRequest) error
}

type RegisterTradeRequestPayload struct {
}

// -- CollectibleTradeItem Types

type CollectibleTradeItem struct {
	TradeRequestID int    `json:"tradeRequestID"`
	CollectibleID  int    `json:"collectibleID"`
	SenderUsername string `json:"senderUsername"`
}

type CollectibleTradeItemStore interface {
	GetCollectibleTradeItemByTradeRequestID(id int) ([]*CollectibleTradeItem, error)
	GetCollectibleTradeItemByCollectibleID(id int) ([]*CollectibleTradeItem, error)
	CreateCollectibleTradeItem(CollectibleTradeItem) error
}

type RegisterCollectibleTradeItemPayload struct {
	TradeRequestID int    `json:"tradeRequestID" validate:"required"`
	CollectibleID  int    `json:"collectibleID" validate:"required"`
	SenderUsername string `json:"senderUsername" validate:"required"`
}

// -- ChoreTradeItem Types

type ChoreTradeItem struct {
	TradeRequestID int    `json:"tradeRequestID"`
	ChoreID        int    `json:"choreID"`
	SenderUsername string `json:"senderUsername"`
}

type ChoreTradeItemStore interface {
	GetChoreTradeItemByTradeRequestID(id int) ([]*ChoreTradeItem, error)
	GetChoreTradeItemByChoreID(id int) ([]*ChoreTradeItem, error)
	CreateChoreTradeItem(ChoreTradeItem) error
}

type RegisterChoreTradeItemPayload struct {
	TradeRequestID int    `json:"tradeRequestID" validate:"required"`
	ChoreID        int    `json:"choreID" validate:"required"`
	SenderUsername string `json:"senderUsername" validate:"required"`
}
