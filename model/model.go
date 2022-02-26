package model

type Data struct {
	Name   string `json:"name"`
	Dept   string `json:"dept"`
	Scores []int  `json:"score"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
