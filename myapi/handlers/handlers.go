package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// /handler のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
		// 何がリクエストされても"Hello, World!”の文字列を返す
		io.WriteString(w, "Hello, World!\n")
}

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "Posting Article...\n")
}

// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request)  {
			io.WriteString(w, "Article List\n")
}

// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request)  {
			articleID :=1
			resString :=fmt.Sprintf("Article No.%d\n", articleID)
			io.WriteString(w, resString)
}

// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request)  {
			io.WriteString(w, "Posting Nice...\n")
}

// /article/comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request)  {
			io.WriteString(w, "Posting Comment...\n")
}
