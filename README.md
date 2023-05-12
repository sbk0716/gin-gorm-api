# 1. Project Overview
* This is a project for an app called `gin-gorm-api`.

## (1)App Features
* This app is able to use below function.

### User Story
* Registration with a username and password.
* Login with a username and password.
* Create a new diary entry.
* Retrieve all diary entries.

## (2)Project Structure
```sh
.
├── README.md
├── controller
│   ├── authentication.go
│   └── entry.go
├── database
│   └── database.go
├── diary_api
├── docker-compose.yml
├── go.mod
├── go.sum
├── helper
│   └── jwt.go
├── main.go
├── middleware
│   └── jwtAuth.go
└── model
    ├── authenticationInput.go
    ├── entry.go
    └── user.go
```

# 2. Usage
## 2.1. Start a postgres server and create a database.

```shell
% ls | grep docker-compose.yml
docker-compose.yml
% docker-compose up -d
% docker ps
CONTAINER ID   IMAGE         COMMAND                  CREATED          STATUS          PORTS                    NAMES
1a4b66aaadc4   postgres:14   "docker-entrypoint.s…"   15 minutes ago   Up 15 minutes   0.0.0.0:5432->5432/tcp   diary_pg
% 
% docker exec -it diary_pg /bin/bash
root@1a4b66aaadc4:/# printenv | grep PASS
POSTGRES_PASSWORD=*******
root@1a4b66aaadc4:/# psql --username super-user --dbname diary_app
psql (14.7 (Debian 14.7-1.pgdg110+1))
Type "help" for help.

diary_app=# \l
                                      List of databases
   Name    |   Owner    | Encoding |  Collate   |   Ctype    |       Access privileges       
-----------+------------+----------+------------+------------+-------------------------------
 diary_app | super-user | UTF8     | en_US.utf8 | en_US.utf8 | 
 postgres  | super-user | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0 | super-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"super-user"              +
           |            |          |            |            | "super-user"=CTc/"super-user"
 template1 | super-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"super-user"              +
           |            |          |            |            | "super-user"=CTc/"super-user"
(4 rows)

diary_app=# \dt
           List of relations
 Schema |  Name   | Type  |   Owner    
--------+---------+-------+------------
 public | entries | table | super-user
 public | users   | table | super-user
(2 rows)

diary_app=# \q
root@1a4b66aaadc4:/# exit
exit
% 
```

## 2.2. Run the app

```shell
% go get .
% go run main.go
Successfully connected to the database
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /auth/register            --> diary_api/controller.Register (3 handlers)
[GIN-debug] POST   /auth/login               --> diary_api/controller.Login (3 handlers)
[GIN-debug] POST   /api/entry                --> diary_api/controller.AddEntry (4 handlers)
[GIN-debug] GET    /api/entry                --> diary_api/controller.GetAllEntries (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
...
...
```


# 3. Check the operation of each API endpoint
## 3.1. `POST /auth/register`
* Registration with a username and password.

```shell
% curl -s -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"testuser01", "password":"sfasf"}' \
    http://localhost:8000/auth/register | jq -r '.'

{
  "user": {
    "ID": 3,
    "CreatedAt": "2023-05-07T14:10:16.243371+09:00",
    "UpdatedAt": "2023-05-07T14:10:16.243371+09:00",
    "DeletedAt": null,
    "username": "testuser01",
    "Entries": null
  }
}
%
```

## 2.2. `POST /auth/login`
* Login with a username and password.

```sh
% curl -s -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"testuser01", "password":"sfasf"}' \
    http://localhost:8000/auth/login | jq -r '.'
{
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODM2NDE4ODksImlhdCI6MTY4MzYzOTg4OSwiaWQiOjN9.vIA1y3pFYU4b29wwtYkVO0mcwLI-gW0sqMjcxbIEHlg"
}
% 
```

## 2.3. `POST /api/entry`
* Create a new diary entry.

```sh
% curl -s -d '{"content":"test content"}' \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODM5MTcwMDcsImlhdCI6MTY4MzkxNTAwNywiaWQiOjN9.D-Y9B-scstj7Y0MsSRoZ1dFodmpWG8mh4UgfPPYEUIQ" \
    -X POST http://localhost:8000/api/entry | jq -r '.'
{
  "data": {
    "ID": 5,
    "CreatedAt": "2023-05-13T03:15:17.798931+09:00",
    "UpdatedAt": "2023-05-13T03:15:17.798931+09:00",
    "DeletedAt": null,
    "content": "test content",
    "UserID": 3
  }
}
% 
```


## 2.4. `GET /api/entry`
* Retrieve all diary entries.

```sh
% curl -s -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODM5MTgwOTcsImlhdCI6MTY4MzkxNjA5NywiaWQiOjN9.lX1TJ3GQOR8l-LB7pHcCgj2-4ZJ6H-WzSc9sl8DePpo" \
    -X GET http://localhost:8000/api/entry | jq -r '.'
{
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2023-05-13T03:11:09.04402+09:00",
      "UpdatedAt": "2023-05-13T03:11:09.04402+09:00",
      "DeletedAt": null,
      "content": "A sample content1",
      "UserID": 3
    },
    {
      "ID": 2,
      "CreatedAt": "2023-05-13T03:11:33.260782+09:00",
      "UpdatedAt": "2023-05-13T03:11:33.260782+09:00",
      "DeletedAt": null,
      "content": "A sample content2",
      "UserID": 3
    },
    {
      "ID": 3,
      "CreatedAt": "2023-05-13T03:12:32.35037+09:00",
      "UpdatedAt": "2023-05-13T03:12:32.35037+09:00",
      "DeletedAt": null,
      "content": "A sample content3",
      "UserID": 3
    },
    {
      "ID": 4,
      "CreatedAt": "2023-05-13T03:13:03.822377+09:00",
      "UpdatedAt": "2023-05-13T03:13:03.822377+09:00",
      "DeletedAt": null,
      "content": "A sample content4",
      "UserID": 3
    },
    {
      "ID": 5,
      "CreatedAt": "2023-05-13T03:15:17.798931+09:00",
      "UpdatedAt": "2023-05-13T03:15:17.798931+09:00",
      "DeletedAt": null,
      "content": "A sample content5",
      "UserID": 3
    }
  ]
}
% 
```