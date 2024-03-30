# Data model

## User
* ID uuid.UUID
* Name string

## Tweet
* ID uuid.UUID
* Content string
* CreationTime uint64

## FeedUser
* Name string
* Feeds Map<string,Map<Map<uuid.UUID, interface{}>>

## Followed
* Name 
* Followers Map<Map<string>>