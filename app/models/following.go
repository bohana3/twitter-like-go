package models

// Followed return a list of followers for a specific user
type Followed struct {
	Name      string                 // following Name
	Followers map[string]interface{} // set of followers
}

func NewFollowing(name string) Followed {
	return Followed{
		Name:      name,
		Followers: make(map[string]interface{}, 0),
	}
}
