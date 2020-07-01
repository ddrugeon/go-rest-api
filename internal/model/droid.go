package model

import "errors"

type Droid struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Company  string `json:"company,omitempty"`
	Class    string `json:"class"`
	Model    string `json:"model"`
	Height   string `json:"height"`
	Vehicles string `json:"vehicles,omitempty"`
}

var (
	ErrNoRecord        = errors.New("models: no matching record found")
	ErrDuplicateRecord = errors.New("models: duplicate record")
)
