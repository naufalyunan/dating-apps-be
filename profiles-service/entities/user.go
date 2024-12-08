package entities

type User struct {
	ID         uint
	Username   string
	Email      string
	IsPremium  bool
	IsVerified bool
}
