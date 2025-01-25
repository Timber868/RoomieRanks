package auth

import "golang.org/x/crypto/bcrypt"

//Necessary for security to hash passwords
func HashPassword(password string) (string, error) {
	//Industry standard library for hashing passwords
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
