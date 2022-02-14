package dtos

type UsernameInput struct {
	Username string `json:"username" binding:"required,username"`
}

type AgeInput struct {
	Age int `json:"age" binding:"min=1,max=5"`
}
