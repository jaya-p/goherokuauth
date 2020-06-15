package goherokuauth

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	var sb strings.Builder
	sb.Grow(n)

	l := len(letters)
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(l)])
	}
	return sb.String()
}

// getAccountID provides accountID on correct username and passwordHash
func getAccountID(username string, passwordHash string) (int, error) {
	if username == "" || passwordHash == "" {
		return -1, errors.New("username and/or passwordHash are empty")
	}

	db, errO := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errO != nil {
		log.Print(errO)
		return -1, errO
	}

	var id int
	row := db.QueryRow("SELECT id FROM account WHERE username=$1 AND passwordHash=$2", username, passwordHash)
	switch errR := row.Scan(&id); errR {
	case nil:
		return id, nil
	default: // including sql.ErrNoRows
		log.Print(errR)
		return -1, errR
	}
}

// createToken on current implementation. 1 user = 1 token. no expiration.
func createToken(id int) (string, error) {
	if id == 0 {
		return "", errors.New("invalid id. id = 0")
	}

	db, errO := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errO != nil {
		log.Print(errO)
		return "", errO
	}

	tokenLength := 32
	token := randomString(tokenLength)

	_, errE := db.Exec("INSERT INTO token(user_id, token) "+
		" VALUES($1, $2) "+
		" ON CONFLICT (user_id) "+
		" DO UPDATE "+
		" SET user_id=$1, token=$2", id, token)
	if errE != nil {
		log.Print(errE)
		return "", errE
	}

	return token, nil
}

// GetToken will get value of a field
func GetToken(username string, passwordHash string) (string, error) {
	userId, errC := getAccountID(username, passwordHash)
	if errC != nil {
		log.Print(errC)
		return "", errC
	}

	token, errG := createToken(userId)
	if errG != nil {
		log.Print(errG)
		return "", errG
	}

	return token, nil
}
