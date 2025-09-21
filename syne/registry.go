package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Server struct {
	Username string
	Url      string
	Token    string
}

type ServerRegistry map[string]*Server

const DEFAULT_SYNE_REGISTRY = ".syne.json"
const SYNE_REGISTRY_ENV = "SYNE_REGISTRY"

func GetRegistryPath() (string, error) {
	reg, ok := os.LookupEnv(SYNE_REGISTRY_ENV)
	if ok {
		return reg, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, DEFAULT_SYNE_REGISTRY), nil
}

func ParseRegistry() (ServerRegistry, error) {
	path, err := GetRegistryPath()
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		fmt.Println("Syne registry not found at", path)
		return make(map[string]*Server), nil
	}
	if err != nil {
		return nil, err
	}

	var registry ServerRegistry
	json.Unmarshal(bytes, &registry)

	return registry, nil
}

func (r *ServerRegistry) Write() error {
	path, err := GetRegistryPath()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, 0644)
}
