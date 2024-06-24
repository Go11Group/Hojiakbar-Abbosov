package models

type CreateStation struct {
	Name string `json:"name"`
}

type UpdateStation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Station struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
