# 1. Usage
## 1.1. Start a postgres server and create a database.
Start a postgres server and create a database.

```shell
% ls | grep docker-compose.yml
docker-compose.yml
% docker-compose up -d
% docker ps 
CONTAINER ID   IMAGE         COMMAND                  CREATED          STATUS          PORTS                    NAMES
1a4b66aaadc4   postgres:14   "docker-entrypoint.s…"   15 minutes ago   Up 15 minutes   0.0.0.0:5432->5432/tcp   diary_pg
% 
```

## 1.2. Run the app

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


# 2. Check the operation of each API endpoint
## 2.1. `POST /auth/register`

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

```sh
curl -s -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"testuser01", "password":"sfasf"}' \
    http://localhost:8000/auth/login | jq -r '.'
{
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODM2NDE4ODksImlhdCI6MTY4MzYzOTg4OSwiaWQiOjN9.vIA1y3pFYU4b29wwtYkVO0mcwLI-gW0sqMjcxbIEHlg"
}
admin@gw-mac gin-gorm-api % 
```
