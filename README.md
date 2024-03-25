# twitter-like-go

manage a Twitter-like app.
The code must be written in go, managing state in any way you prefer (simpler is better)
When done, upload the code to github.
The functions you should implement:
1. CRUD a user:
    a. CreateUser(name string) error\
    b. GetUser(name string) (*User, error)\
    c. UpdateUser(oldName string, newName string) error\
    d. DeleteUser(name string) error
2. Users can follow each other
    a. FollowUser(followerName string, followingName string) error\
    b. UnfollowUser(followerName string, followingName string) error
3. Users can post tweets\
    a. PostTweet(userName string, tweetContent string) error
4. Each user have their own feed (feed == tweets by other users that this user follows, sorted by most recent first)\
    a. GetUserFeed(userName string) ([]Tweet, error)
5. Return mutual followers – users that follow both userName1 and userName2
    a. GetMutualFollowers(userName1 string, userName2 string) ([]*User, error)
6. Identify the top 'n' users with the highest number of followers.
    a. GetTopInfluencers(n int) ([]User, error)\

Notes:\
    ● Unit tests are not required\
    Success criteria, by priority:\
    ● Make it work\
    ● Model the data correctly\
    ● Code simple and readable\
    ● Don’t have bugs (consider edge cases)