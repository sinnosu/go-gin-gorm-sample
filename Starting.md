# how to create project

## reference
https://medium.com/@frederich/todo-list-node-js-like-architecture-rest-api-with-golang-gin-gorm-mysql-68eea2409e57


## go mod init
```
$ cd backend
# go mod init <github-link>
$ go mod init github/sinnosu/go-gin-gorm-sample
```

## go package install
```
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get gorm.io/driver/mysql
go get github.com/joho/godotenv
```

## migration tools dbmate
# gormにmigration機能あるけどdbmateの方が優秀なので(特にmacでは)それを使う
 - install dbmate on mac
 `brew install dbmate`

 - create migration file -> db/migrations/~
 `dbmate new todos`

 - edit migration file
 backend/db/migrations/~_<table name>.sql

 - 

# gin sample source 
- create files
`$ go mod tidy`
`$ go run main.go
```


