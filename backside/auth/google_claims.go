package auth

type GoogleClaims struct {
	Id            string
	Email         string
	EmailVerified bool
	FirstName     string
	LastName      string
}
