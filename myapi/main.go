package main

import (
	"io"
	"log"
	"net/http"
)

func main()  {
		//ハンドラを定義
		helloHandler := func(w http.ResponseWriter, req *http.Request) {
			// 何がリクエストされても"Hello, World!”の文字列を返す
			io.WriteString(w, "Hello, World!\n")
		}

		// 定義したhelloHandlerを使うように登録
		http.HandleFunc("/", helloHandler)

		// サーバー起動時のログを出力
		log.Println("server start at port 8080")

		// ListenAndServe関数にて、サーバーを起動
		log.Fatal(http.ListenAndServe(":8080", nil))
}