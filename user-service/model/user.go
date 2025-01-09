package model

type User struct {
	Email     string `gorm:"not null;unique"`
	FirstName string
	LastName  string
	IsOnline  bool `gorm:"default:false"`
}

func (User) TableName() string {
	return "accounts"
}

func Seed() []User {
	return []User{
		{Email: "satu@mailinator.com", FirstName: "Satu", LastName: "Satu"},
		{Email: "dua@mailinator.com", FirstName: "Dua", LastName: "Dua"},
		{Email: "tiga@mailinator.com", FirstName: "Tiga", LastName: "Tiga"},
		{Email: "empat@mailinator.com", FirstName: "Empat", LastName: "Empat"},
		{Email: "lima@mailinator.com", FirstName: "Lima", LastName: "Lima"},
		{Email: "enam@mailinator.com", FirstName: "Enam", LastName: "Enam"},
		{Email: "tujuh@mailinator.com", FirstName: "Tujuh", LastName: "Tujuh", IsOnline: true},
		{Email: "delapan@mailinator.com", FirstName: "Delapan", LastName: "Delapan", IsOnline: true},
	}
}
