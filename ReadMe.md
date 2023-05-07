# reference
 - docker & devcontainer : https://qiita.com/Chika110/items/e76978548bb2019a6c49

# mysql 
 - enter in docker container
   `% docker exec -it f1f /bin/sh`
 - connct database
   `%mysql -h localhost -u dev_user -ppassword test1`
     -> mysql -h localhost -u <username> -ppassword <database name>

# gin run 


## dbmate

https://github.com/amacneil/dbmate
### linux -> 一応dockerにも記述した
```
$ sudo curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
$ sudo chmod +x /usr/local/bin/dbmate
```

### dbmateでdocker-dbに接続できるようにする
- .envファイルを作成
- .envにDATABASE_URLを書く
- 接続確認
  `dbmate --wait-timeout=5s wait`
- migration fileの作成
  `dbmate new create_todos_table`
- migration run (migrate upに記載してある箇所を実行)
  `$ dbmate up`
- migration run (rollback downに記載してある箇所を実行)
  `$ dbmate down`

## API起動
` go run main.go`

### curlで確認
 - get
`curl http://localhost:8080/todo `
 - post
 `curl -X POST -H "Content-Type: application/json" -d '{"name":"shinnosuke yayama", "description":"hajimeteno go gin"}' http://localhost:8080/todo`
 - delete
 `curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/todo/2`
