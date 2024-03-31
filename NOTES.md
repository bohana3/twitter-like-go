# Design patterns used:
- repository pattern: each data is represented by a repository that is specialised in one data type (decoupling of data)\
    -> simple and single logic in each repository (single responsibility principle)\
    -> during data modeling you need to have storage concerns (system design problematic)\
in this case the tweeter content is the most heavy data therefore need to separate it from other data\
We could think of using for feed_repository an elastic storage like Redis (easy scalable) and for all other repo we can use an SQL database
- each class depends on the abstraction (dependency inversion principle)\
    -> you encapsulate data structures (cf. maps below) in the data access layer (infra directory)
    -> it allows easily to mock class and data if needed.\
    I did not create it but in EDR we are using mockery (https://vektra.github.io/mockery/latest/) that auto-generate mocks based on interface. Therefore it's important to create interfaces. 

- code organization: Domain Driven Design (DDD)\
    -> each handler is using one or more several repositories and contains the business logic (business layer)\
    -> each repository is like a table in a DB and exposes basic data manipulation (data access layer)\
    -> each handler see only interfaces of repositories (decoupling between business and data access layer)\
    -> you can read about DDD here https://sarathsp06.medium.com/domain-driven-design-with-go-be3066ae213c


# Data structures used:

All relational data is stored in 3 maps - cf. [DataModel](dataModel.md)

## Users
`Map<String,User>`\
"Users" stores the list of users (followed users or followers)

## FeedUser
`Map<string,Map<Map<uuid.UUID, interface{}>>`

"FeedUser" stores for each user a set of tweetIDs (we use map in Goland when we need a set, meaning a collection of unique items)

## FollowedUser
`Map<string,Map<string,nil>>`

"FollowedUser" stores a list of "followed user".
For each "followed user" we store a set of "followers" (we used map for the same reasons)
