# Data model

## User
* ID uuid.UUID
* Name string
* TweetIds []uuid.UUID

## Feed
* ID uuid.UUID
* Content string

## Follower
* UserId uuid.UUID
* FollowerUserIds []uuid.UUID