package main

import (
	_ "embed"
	"encoding/json"
	"os"
)

//go:embed authTemplate.json
var templateDatabase string

type AuthDatabase struct {
	admins   []string
	users    map[string]string
	sessions map[string]struct {
		username   string
		last_login int
	}
	shared_files map[string]struct {
		files       []string
		accesses    int
		time_shared int
		lifetime    int
	}
	permissions map[string][]string
}

func NewAuthDatabase(file string) (*AuthDatabase, error) {
	content, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var payload AuthDatabase
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func CreateAuthDatabase(file string) (*AuthDatabase, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	_, err = f.WriteString(templateDatabase)
	if err != nil {
		return nil, err
	}

	return NewAuthDatabase(file)
}

func (d *AuthDatabase) CheckAuth(username string, password string) bool {
	saved_password, ok := d.users[username]
	if !ok {
		return false
	}
	return saved_password == password
}

func (d *AuthDatabase) GetSessionToken(username string) string {

}
