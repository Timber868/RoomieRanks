package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Timber868/roomieranks/service/household"
	"github.com/Timber868/roomieranks/types"
	"github.com/gorilla/mux"
)

func TestGetUserWithCollectibles(t *testing.T) {
	// Create a new userstore that has been mocked
	userStore := &mockUserStore{}
	// Create a nil household store since we don't need it for this test
	var householdStore *household.Store
	handler := NewHandler(userStore, householdStore)

	t.Run("Should return user with collectibles", func(t *testing.T) {
		// Create a request to get user with username "testuser"
		req, err := http.NewRequest(http.MethodGet, "/user/testuser", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Set up router with the handler
		router := mux.NewRouter()
		router.HandleFunc("/user/{username}", handler.handleGetUser)

		// Create a response recorder
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Check that the request was successful
		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}

		// Parse the response to check if collectibles are included
		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}

		// Check that collectibles field exists and is an array
		if collectibles, exists := response["collectibles"]; !exists {
			t.Error("expected 'collectibles' field in response")
		} else {
			collectiblesArray, ok := collectibles.([]interface{})
			if !ok {
				t.Error("expected 'collectibles' to be an array")
			}

			// Check that we have the expected collectibles
			if len(collectiblesArray) != 2 {
				t.Errorf("expected 2 collectibles, got %d", len(collectiblesArray))
			}
		}
	})
}

type mockUserStore struct{}

func (s *mockUserStore) GetUserByUsername(username string) (*types.User, error) {
	// Return a mock user with collectibles
	return &types.User{
		Username:    "testuser",
		Name:        "Test User",
		Email:       "test@example.com",
		Password:    "hashedpassword",
		HouseholdID: 1,
		Title:       "Pokemon Master",
		Level:       5,
		XP:          150,
		Collectibles: []types.Collectible{
			{
				ID:           1,
				Name:         "Pikachu",
				Rarity:       "Common",
				Type:         "Electric",
				ImageURL:     "https://example.com/pikachu.jpg",
				UserUsername: "testuser",
			},
			{
				ID:           2,
				Name:         "Charizard",
				Rarity:       "Rare",
				Type:         "Fire",
				ImageURL:     "https://example.com/charizard.jpg",
				UserUsername: "testuser",
			},
		},
	}, nil
}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func (s *mockUserStore) ChangeTitle(username string, title string) error {
	return nil
}

func (s *mockUserStore) AddXP(username string, xp int) error {
	return nil
}

func (s *mockUserStore) ChangeHousingID(username string, householdID int) error {
	return nil
}

func (s *mockUserStore) GetCollectiblesByUsername(username string) ([]types.Collectible, error) {
	return []types.Collectible{
		{
			ID:           1,
			Name:         "Pikachu",
			Rarity:       "Common",
			Type:         "Electric",
			ImageURL:     "https://example.com/pikachu.jpg",
			UserUsername: "testuser",
		},
		{
			ID:           2,
			Name:         "Charizard",
			Rarity:       "Rare",
			Type:         "Fire",
			ImageURL:     "https://example.com/charizard.jpg",
			UserUsername: "testuser",
		},
	}, nil
}

type mockHouseholdStore struct{}

func (s *mockHouseholdStore) GetHouseholdByID(id int) (*types.Household, error) {
	return &types.Household{
		ID:   id,
		Name: "Test Household",
	}, nil
}

func (s *mockHouseholdStore) CreateHousehold(household types.Household) (int, error) {
	return 1, nil
}

// Integration test to verify the getUser endpoint works with real database
func TestGetUserWithCollectiblesIntegration(t *testing.T) {
	// This test requires a database connection
	// Skip if not in integration test mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// TODO: Set up test database connection
	// TODO: Create test user with collectibles
	// TODO: Test the actual API endpoint
	t.Log("Integration test placeholder - requires database setup")
}
