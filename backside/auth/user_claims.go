package auth

type UserClaims struct {
	Id            string
	Email         string
	EmailVerified bool
	FirstName     string
	LastName      string
}
