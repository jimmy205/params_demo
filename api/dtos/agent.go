package dtos

type UsernameInput struct {
	Username string `json:"username" binding:"required"`
}

type AgeInput struct {
	Age int `json:"age" binding:"min=1,max=5"`
}

type GenderInput struct {
	Gender string `json:"gender" binding:"oneof=male female"`
}
