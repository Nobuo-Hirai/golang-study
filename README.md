# ローカル起動、動作確認方法
## 起動
- cd golang-study
- go run main.go
## 動作確認
- http://localhost:8080/hello へアクセス  
  または
- curl http//:localhost:8080/hello
### http status確認
- curl http://localhost:8080/hello -X GET -w '%{http_code}\n'
- curl http://localhost:8080/hello -X POST -w '%{http_code}\n'