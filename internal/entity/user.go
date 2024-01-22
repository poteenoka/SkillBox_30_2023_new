package entity

type User struct {
	ID   string `json:"id"       example:"1"`
	Name string `json:"name"       example:"ilya"`
	Age  int    `json:"age"       example:"10"`
}

type Userfriends struct {
	Friends []int
}
