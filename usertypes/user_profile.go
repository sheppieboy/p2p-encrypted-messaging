package usertypes

type UserProfile struct {
	Name string
	Port string
}

func NewUserProfile(name string, port string) * UserProfile{
	return &UserProfile{
		Name: name,
		Port: port,
	}
}