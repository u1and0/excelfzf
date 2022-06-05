package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

var lines []string

// Excelデータ読み込み
// 時間がかかるとmain関数へ行かない
// つまりブラウザのページ読み込みがストップ
// してしまうので、非同期処理で裏で読み込みをかけておく
// ページを読み込んだ時点でのデータを配信するので、
// 不完全なデータになる場合もある。
// そういうときは、ブラウザ上でリロードかけると
// 最新のデータを配信してくれる。
func init() {
	go func() {
		path := "./data/sample-xlsx-file-for-testing.xlsx"
		f, _ := excelize.OpenFile(path)
		defer f.Close()
		// 行ごとに読み込み
		// 2次元配列で返す
		rows, _ := f.GetRows("Sheet1")
		lines = make([]string, len(rows))
		for i, row := range rows {
			lines[i] = strings.Join(row, " ")
		}
	}()
}

// サーバー立ち上げて
// http://localhost:8080/list で
func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*.tmpl")
	// エントリポイント
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	// list形式のJSONを配信するAPI
	r.GET("/list", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, lines)
	})
	r.Run()
}
