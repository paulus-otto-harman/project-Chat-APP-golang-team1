package seed

import "project/auth-service/model"

func User() []model.User {
	return []model.User{
		{Email: "satu@mailinator.com"},
		{Email: "dua@mailinator.com"},
		{Email: "tiga@mailinator.com"},
		{Email: "empat@mailinator.com"},
		{Email: "lima@mailinator.com"},
		{Email: "enam@mailinator.com"},
		{Email: "tujuh@mailinator.com"},
		{Email: "delapan@mailinator.com"},
		{Email: "sembilan@mailinator.com"},
		{Email: "sepuluh@mailinator.com"},
	}
}
