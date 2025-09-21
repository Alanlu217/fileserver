package main

import (
	_ "embed"
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
)

//go:embed authTemplate.json
var templateDatabase string

// Keep session alive for a week of inactivity
const SESSION_LIFETIME = 604800

type Session struct {
	Username   string
	Last_login int
}

type SharedFile struct {
	Files       []string
	Accesses    int
	Time_shared int
	Lifetime    int
}

type UserPermission map[string]int

type AuthDatabase struct {
	Filename     string
	Admin        []string                  `json:"admin"`
	Users        map[string]string         `json:"users"`
	Sessions     map[string]Session        `json:"sessions"`
	Shared_files map[string]SharedFile     `json:"shared_files"`
	Permissions  map[string]UserPermission `json:"permissions"`
}

func LoadAuthDatabase(file string) (*AuthDatabase, error) {
	databaseFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer databaseFile.Close()

	content, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var payload AuthDatabase
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}

	payload.Filename = file

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

	return LoadAuthDatabase(file)
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
