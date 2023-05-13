# 1. Project Overview
* This is a project for an app called `gin-gorm-api`.

## (1)App Features
* This app is able to use below function.

### User Story
* Registration with a username and password.
* Login with a username and password.
* Create a new diary entry.
* Retrieve all your entries.
* Retrieve any entry of yourself.

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
CONTAINER ID   IMAGE                    COMMAND                  CREATED          STATUS          PORTS                    NAMES
9420238acadd   gin-gorm-api-diary_api   "go run main.go"         49 seconds ago   Up 48 seconds   0.0.0.0:8000->8000/tcp   diary_api
93716eee56a1   postgres:14              "docker-entrypoint.s…"   49 seconds ago   Up 48 seconds   0.0.0.0:5432->5432/tcp   diary_pg
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

## 2.2. Run the app(local)

```shell
% make up
docker-compose up
[+] Running 2/0
 ⠿ Container diary_pg   Created                             0.0s
 ⠿ Container diary_api  Created                             0.0s
Attaching to diary_api, diary_pg
diary_pg   | 
diary_pg   | PostgreSQL Database directory appears to contain a database; Skipping initialization
diary_pg   | 
diary_pg   | 2023-05-13 03:58:12.477 UTC [1] LOG:  starting PostgreSQL 14.8 (Debian 14.8-1.pgdg110+1) on aarch64-unknown-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
diary_pg   | 2023-05-13 03:58:12.477 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
diary_pg   | 2023-05-13 03:58:12.477 UTC [1] LOG:  listening on IPv6 address "::", port 5432
diary_pg   | 2023-05-13 03:58:12.479 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
diary_pg   | 2023-05-13 03:58:12.483 UTC [27] LOG:  database system was shut down at 2023-05-13 03:58:11 UTC
diary_pg   | 2023-05-13 03:58:12.487 UTC [1] LOG:  database system is ready to accept connections
diary_api  | filePath: ".env.local"
diary_api  | Successfully connected to the database..
...
...
```

## 2.3. Run the app(production)

```shell
% make up/prod
docker-compose -f docker-compose.production.yaml up
[+] Running 2/0
 ⠿ Container diary_pg   Created                                                                       0.0s
 ⠿ Container diary_api  Created                                                                       0.0s
Attaching to diary_api, diary_pg
diary_pg   | 
diary_pg   | PostgreSQL Database directory appears to contain a database; Skipping initialization
diary_pg   | 
diary_pg   | 2023-05-13 04:08:00.220 UTC [1] LOG:  starting PostgreSQL 14.8 (Debian 14.8-1.pgdg110+1) on aarch64-unknown-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
diary_pg   | 2023-05-13 04:08:00.220 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
diary_pg   | 2023-05-13 04:08:00.220 UTC [1] LOG:  listening on IPv6 address "::", port 5432
diary_pg   | 2023-05-13 04:08:00.224 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
diary_pg   | 2023-05-13 04:08:00.228 UTC [27] LOG:  database system was shut down at 2023-05-13 04:07:57 UTC
diary_pg   | 2023-05-13 04:08:00.231 UTC [1] LOG:  database system is ready to accept connections
diary_api  | filePath: ".env.production"
diary_api  | Successfully connected to the database
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
* Retrieve all your entries.

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

## 2.5. `GET /api/entry/:id`
* Retrieve any entry of yourself.

```sh
% curl -s -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODM5MTgwOTcsImlhdCI6MTY4MzkxNjA5NywiaWQiOjN9.lX1TJ3GQOR8l-LB7pHcCgj2-4ZJ6H-WzSc9sl8DePpo" \
    -X GET http://localhost:8000/api/entry/1 | jq -r '.'
{
  "error": "The target entryId does not exist. [entryId: 1]"
}
%
% curl -s -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODM5MTgwOTcsImlhdCI6MTY4MzkxNjA5NywiaWQiOjN9.lX1TJ3GQOR8l-LB7pHcCgj2-4ZJ6H-WzSc9sl8DePpo" \
    -X GET http://localhost:8000/api/entry/2 | jq -r '.'
{
  "entry": {
    "ID": 2,
    "CreatedAt": "2023-05-13T03:11:33.260782+09:00",
    "UpdatedAt": "2023-05-13T03:11:33.260782+09:00",
    "DeletedAt": null,
    "content": "A sample content2",
    "UserID": 3
  }
}
% 
```