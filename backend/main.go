package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

// 我们需要定义一个 Upgrader
// 它需要定义 ReadBufferSize 和 WriteBufferSize
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,

    // 可以用来检查连接的来源
    // 这将允许从我们的 React 服务向这里发出请求。
    // 现在，我们可以不需要检查并运行任何连接
    CheckOrigin: func(r *http.Request) bool { return true },
}

// 定义一个 reader 用来监听往 WS 发送的新消息
func reader(conn *websocket.Conn) {
    for {
        // 读消息
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        // 打印消息
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}

// 定义 WebSocket 服务
func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)

    // 将连接更新为 WebSocket 连接
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    // 一直监听 WebSocket 连接上传来的新消息
    reader(ws)
}

func setupRoutes() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
    })

    // 将 `/ws` 端点交给 `serveWs` 函数处理
    http.HandleFunc("/ws", serveWs)
}

func main() {
    fmt.Println("Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":8080", nil)
}