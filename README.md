# sn-test

SN Test API Rest.

## Prerequisites:

SN Test may require the following tools to be installed beforehand:

```bash
~$ go get go.mongodb.org/mongo-driver/mongo
~$ go get go.mongodb.org/mongo-driver/mongo/options
~$ go get go.mongodb.org/mongo-driver/bson
~$ go get go.mongodb.org/mongo-driver/bson/primitive
~$ go get golang.org/x/crypto/bcrypt
~$ go get github.com/gorilla/mux
~$ go get github.com/rs/cors
~$ go get github.com/dgrijalva/jwt-go
```

## MongoDB database and collections

SN Test API requires having the following database and collections previously created

### Database
- test

### Collections

- follows
    ```
    _id: ObjectId(string)
    userId: string
    userFollowedID: string
    ```
- likes
    ```
    _id: ObjectId(string)
    userId: string
    likesList: [
        0: string, (message ID)
        1: string, (message ID)
        2: string, (message ID)
        ...
    ]
    ```
- messages
    ```
    _id: ObjectId(string)
    datetime: YYYY-MM-DDThh:mm:ss.SSS+00:00
    userId: string
    message: string
    likesCount: uint64
    ```
- users
    ```
    _id: ObjectId(string)
    name: string
    lastName: string
    dateOfBirth: YYYY-MM-DDThh:mm:ss.SSS+00:00
    email: string
    password: string (encrypted)
    avatar: string
    banner: string
    biography: string
    location: string
    webSite: string
    followingCount: uint64
    followersCount: uint64
    ```
