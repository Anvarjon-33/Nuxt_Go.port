package main

type User struct {
	Name     string `gorm:"name" form:"name"`
	LastName string `gorm:"last_name" form:"last_name"`
	Email    string `gorm:"email" form:"email"`
	Password string `gorm:"password" form:"password"`
}

func main() {

}
