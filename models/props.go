package models

type IndexRouteProps struct {
	InitialCount int `json:"initialCount"`
}

type AboutRouteProps struct {
	Content string `json:"content"`
}

type PostRouteProps struct {
	Content string `json:"content"`
}