package main

import (
	"github.com/Nobuo-Hirai/golang-study/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//helloHandler という変数でハンドラを定義するだけでは、それが使われることはありません。
	//サーバーでハンドラを使うためには、「このハンドラを使ってください」という登録作業を行う必要があります。
	//今回この作業は、net/http パッケージ内で定義されているHandleFunc 関数4で行っています。
	//HandleFunc 関数の第二引数に先ほど作成した自作ハンドラhelloHandler を渡すことで、サーバーでこのハンドラが使われるようになります。
	//http.HandleFunc("/", helloHandler)
	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDataHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceArticleHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	//サーバー起動時のログを出力
	log.Println("server start at port 8080")

	//今回は第一引数に:8080 を指定しているため、サーバー起動場所はlocalhost:8080
	//http.ListenAndServe 関数の返り値error は、サーバー起動時に起こったエラーがあったときに返ってきます。
	//もしそれが起こった場合にエラーの内容を拾ってログとして出力するために、返り値のエラーをlog.Fatal 関数に渡しています。
	log.Fatal(http.ListenAndServe(":8080", r))

	// 二行で書いたパターン
	//err := http.ListenAndServe(":8080", nil)
	//log.Fatal(err)
}
