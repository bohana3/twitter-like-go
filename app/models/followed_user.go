package models

// FollowedUser return a list of followers for a specific user
type FollowedUser struct {
	Name      string                 // following Name
	Followers map[string]interface{} // set of followers
}

func NewFollowing(name string) *FollowedUser {
	return &FollowedUser{
		Name:      name,
		Followers: make(map[string]interface{}, 0),
	}
}
