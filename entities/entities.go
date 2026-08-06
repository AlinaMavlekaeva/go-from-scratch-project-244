package entities

type Tree struct {
	Key      string `json:"key"`
	Status   string `json:"status"`
	OldValue any    `json:"oldValue"`
	Value    any    `json:"value"`
	Deep     int    `json:"deep"`
}
