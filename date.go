package main

import (
    "fmt"
    "net/http"
    "log"
    "html/template"
    "time"
)

type TemplateData struct {
    Title string
    Datetime string
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパース
    tmpl := template.Must(template.ParseFiles("views/index.html", "views/body.html"))

    // タイトル
    title := "現在の時刻"

    // bodyに表示するコンテンツ
    // 表示レイアウトを設定
    const layout = "Now, Monday Jan 02 15:04:05 JST 2006"
    
    datetime := fmt.Sprint(time.Now().Format(layout))

    // テンプレートを実行して出力
    templatedata := TemplateData{title, datetime}
    if err := tmpl.ExecuteTemplate(w, "base", templatedata); err != nil {
        fmt.Println(err)

    }
}

func main() {

    http.HandleFunc("/date", HelloServer)
    log.Printf("Start Go HTTP Server")
    err := http.ListenAndServe(":4000", nil)
    if err != nil {
       log.Fatal("ListenAndServe: ", err)
    
    }
}
