package collectible

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/Timber868/roomieranks/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCollectibleByID(id int) (*types.Collectible, error) {
	rows, err := s.db.Query("SELECT * FROM collectibles WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	c := new(types.Collectible)

	for rows.Next() {
		c, err = scanRowIntoCollectible(rows)

		if err != nil {
			return nil, err
		}
	}

	if c.ID == 0 {
		return nil, nil
	}

	return c, nil
}

func scanRowIntoCollectible(rows *sql.Rows) (*types.Collectible, error) {
	collectible := new(types.Collectible)

	//Directly map the data to our household via pointers
	if err := rows.Scan(
		&collectible.ID,
		&collectible.Name,
		&collectible.Rarity,
		&collectible.Type,
		&collectible.ImageURL,
		&collectible.UserUsername,
	); err != nil {
		//Hadnle errors if any were found
		return nil, err
	}

	//Only one houshold per person
	return collectible, nil
}

// Adjust fields as needed.
type PokemonTCGResponse struct {
	Data []struct {
		ID     string   `json:"id"`
		Name   string   `json:"name"`
		Rarity string   `json:"rarity"` // <- Add this
		Types  []string `json:"types"`  // <- And this
		Images struct {
			Small string `json:"small"`
			Large string `json:"large"`
		} `json:"images"`
		// You can add more fields as needed
	} `json:"data"`
}

func (s *Store) CreateCollectible(username string) error {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Retrieve the environment variables
	apiKey := os.Getenv("API_KEY")

	// 1. Generate a random Pokédex number in [1..155]
	rand.Seed(time.Now().UnixNano())
	randomDexNum := rand.Intn(155) + 1 // e.g., 57 or 123, etc.

	// 2. Build the query for the nationalPokedexNumbers search
	//    For example, q=nationalPokedexNumbers:123
	q := fmt.Sprintf("nationalPokedexNumbers:%d", randomDexNum)
	baseURL := "https://api.pokemontcg.io/v2/cards"

	// Encode the query so special chars (like spaces) are escaped correctly
	fullURL := fmt.Sprintf("%s?q=%s", baseURL, url.QueryEscape(q))

	// 3. Create the request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add your API key to the header
	req.Header.Set("X-Api-Key", apiKey)

	// 4. Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Unmarshal into our response struct
	var tcgResp PokemonTCGResponse
	err = json.Unmarshal(body, &tcgResp)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// 6. If multiple cards are returned, pick one at random
	if len(tcgResp.Data) == 0 {
		fmt.Printf("No cards found for nationalDex %d\n", randomDexNum)
		return nil
	}

	randomIndex := rand.Intn(len(tcgResp.Data))
	card := tcgResp.Data[randomIndex]

	// 7. Print out details of the random card
	fmt.Printf("Random Dex: %d\n", randomDexNum)
	fmt.Printf("Name:       %s\n", card.Name)
	fmt.Printf("Rarity:  %s\n", card.Rarity)
	fmt.Printf("Type:   %v\n", card.Types)
	fmt.Printf("Small Img:  %s\n", card.Images.Small)

	_, err = s.db.Exec("INSERT INTO collectible (name, rarity, type, image_url, user_username) VALUES (?, ?, ?, ?, ?)", card.Name, card.Rarity, strings.Join(card.Types, ","), card.Images.Small, username)

	if err != nil {
		return err
	}

	return nil
}
