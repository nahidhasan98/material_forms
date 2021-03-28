package model

type UserTable struct {
	ID                   string `json:"aid"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	DateOfBirth          int64  `json:"dateOfBirth"`
	Phone                string `json:"phone"`
	CreatedAt            int64  `json:"createdAt"`
	IsVerified           int    `json:"isVerified"`
	AccVerifyToken       string `json:"accVerifyToken"`
	AccVerifyTokenSentAt int64  `json:"accVerifyTokenSentAt"`
	PassResetToken       string `json:"passResetToken"`
	PassResetTokenSentAt int64  `json:"passResetTokenSentAt"`
	Type                 string `json:"type"`
	Status               int    `json:"status"`
}
