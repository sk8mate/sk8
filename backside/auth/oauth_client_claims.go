package auth

type OAuthClientClaims struct {
	Id            string
	Email         string
	EmailVerified bool
	FirstName     string
	LastName      string
}
