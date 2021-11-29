package domain

type User struct {
	Id        int       `bson:"_id,omitempty"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	OAuthId   string    `bson:"oauth_id"`
}
