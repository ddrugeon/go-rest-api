package model

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

func NewDroid(id string, name string, droid_type string, company string, class string, model string, height string, vehicles string) Droid {
	return Droid{
		ID:       id,
		Name:     name,
		Type:     droid_type,
		Company:  company,
		Class:    class,
		Model:    model,
		Height:   height,
		Vehicles: vehicles,
	}
}
