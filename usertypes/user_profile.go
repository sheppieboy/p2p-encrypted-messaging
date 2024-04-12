package usertypes

type UserProfile struct {
	Name string
	Port string
}

func newUserProfile(name string, port string) * UserProfile{
	return &UserProfile{
		Name: name,
		Port: port,
	}
}