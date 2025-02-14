package models

type Community struct {
	ID   string `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}
