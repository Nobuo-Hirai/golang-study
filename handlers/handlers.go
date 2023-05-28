package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// サーバーで使用するハンドラの定義
// ハンドラとは「HTTPリクエストを受け取ってそれに対するHTTPレスポンスの内容をコネクションに書き込む」関数
// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		// ハンドラの処理内容:
		// 何がきても、"Hello, World!"の文字列を返す
		io.WriteString(w, "Hello, world!\n")
	} else {
		// 405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		// 405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article Lst...\n")
	} else {
		// 405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/1 のハンドラ
func ArticleDataHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		// 405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/nice のハンドラ
func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		// 405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		// 405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
