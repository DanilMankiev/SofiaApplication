package entity

type User struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type RegiterInput struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}
