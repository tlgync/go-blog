package user

type Model struct {
	ID       uint   `json:"id" gorm:"primary_key" gorm:"AUTO_INCREMENT"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
