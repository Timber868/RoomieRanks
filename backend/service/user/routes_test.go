package user

// func TestUserServiceHandlers(t *testing.T) {
// 	//Create a new userstore that has been mocked
// 	userStore := &mockUserStore{}
// 	handler := NewHandler(userStore)

// 	t.Run("Invalid payload", func(t *testing.T) {
// 		/*
// 			Test with an invalid email section of the payload. Should fail with a bad request
// 		*/

// 		payload := types.RegisterUserPayload{
// 			Password: "asd",
// 			Email:    "sad",
// 		}

// 		//Youve got to marshal your struct into a json
// 		marshalled, _ := json.Marshal(payload)

// 		//Also need to transform it into a new byte buffer
// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

// 		if err != nil {
// 			//Automatically fails a test
// 			t.Fatal(err)
// 		}

// 		//Actually make the request to our handler server
// 		requestAgainstHandler := httptest.NewRecorder()
// 		router := mux.NewRouter()
// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(requestAgainstHandler, req)

// 		//Our test shoulve failed so check that
// 		if requestAgainstHandler.Code != http.StatusBadRequest {
// 			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, requestAgainstHandler.Code)
// 		}
// 	})

// 	t.Run("Should correctly register the user", func(t *testing.T) {
// 		/*
// 			Test with a valid payload. Should successfully create the user
// 		*/

// 		payload := types.RegisterUserPayload{
// 			Username: "user",
// 			Name:     "hi",
// 			Password: "invalid",
// 			Email:    "sad@gmail.com",
// 		}

// 		//Youve got to marshal your struct into a json
// 		marshalled, _ := json.Marshal(payload)

// 		//Also need to transform it into a new byte buffer
// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

// 		if err != nil {
// 			//Automatically fails a test
// 			t.Fatal(err)
// 		}

// 		//Actually make the request to our handler server
// 		requestAgainstHandler := httptest.NewRecorder()
// 		router := mux.NewRouter()
// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(requestAgainstHandler, req)

// 		//Our test shoulve failed so check that
// 		if requestAgainstHandler.Code != http.StatusCreated {
// 			t.Errorf("expected status code %d, got %d", http.StatusCreated, requestAgainstHandler.Code)
// 		}
// 	})
// }

// type mockUserStore struct{}

// // Methods needed for a user store need to be mocked
// func (s *mockUserStore) GetUserByUsername(email string) (*types.User, error) {
// 	return nil, fmt.Errorf("username not found") //This needs to be mocked to ignore the case where user already exists
// }

// func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
// 	return nil, nil
// }

// func (s *mockUserStore) CreateUser(user types.User) error {
// 	return nil
// }

// func (s *mockUserStore) ModifyUser(user types.User) error {
// 	return nil
// }

// func (s *mockUserStore) ChangeTitle(username string, title string) error {
// 	return nil
// }

// func (s *mockUserStore) LevelUp(username string) error {
// 	return nil
// }

// func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
// 	return nil, nil
// }
