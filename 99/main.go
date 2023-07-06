package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 定义WebSocket连接的结构体
type WebSocketConnection struct {
	Conn *websocket.Conn
}

// 定义房间连接池的结构体
type RoomConnectionPool struct {
	connections map[string][]*WebSocketConnection
	mutex       sync.Mutex
}

// 创建房间连接池，分配内存空间
func NewRoomConnectionPool() *RoomConnectionPool {
	return &RoomConnectionPool{
		connections: make(map[string][]*WebSocketConnection),
	}
}

// 从连接池中获取可用的WebSocket连接
func (pool *RoomConnectionPool) GetConnection(roomID string) (*WebSocketConnection, bool) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	conns := pool.connections[roomID]
	if len(conns) > 0 {
		conn := conns[0]
		pool.connections[roomID] = conns[1:]
		return conn, true
	}

	return nil, false
}

// 将WebSocket连接释放回连接池
func (pool *RoomConnectionPool) ReleaseConnection(roomID string, conn *WebSocketConnection) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	pool.connections[roomID] = append(pool.connections[roomID], conn)
}

// 处理WebSocket请求
func handleWebSocket(pool *RoomConnectionPool, w http.ResponseWriter, r *http.Request) {
	// 升级HTTP请求为WebSocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	// 获取房间ID
	roomID := r.URL.Query().Get("roomID")

	// 从连接池中获取可用的WebSocket连接
	wsConn, ok := pool.GetConnection(roomID)
	if !ok {
		// 如果没有可用的连接，则创建一个新的连接并添加到连接池中
		wsConn = &WebSocketConnection{Conn: conn}
	} else {
		// 如果有可用的连接，则重用现有的连接
		wsConn.Conn = conn
	}

	// 在连接关闭时将其释放回连接池
	defer func() {
		pool.ReleaseConnection(roomID, wsConn)
	}()

	// 保持WebSocket连接打开
	for {
		messageType, p, err := conn.ReadMessage() //接收客户端传输的数据
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println("WebSocket read error:", err)
			} else {
				log.Println("WebSocket connection closed:", err)
			}
			break
		}
		fmt.Println("Received message:", string(p))
		switch messageType {
		case websocket.TextMessage:
			// 处理接收到的文本消息 p
			var data map[string]interface{}
			json.Unmarshal(p, &data)
			data["createAt"] = time.Now().Format("2006-01-02 15:04:05")
			bytes, _ := json.Marshal(data)
			sendToRoom(pool, roomID, bytes)
		case websocket.BinaryMessage:
			// 处理接收到的二进制消息 p
		case websocket.CloseMessage:
			// 处理关闭连接的逻辑
		case websocket.PingMessage:
			// 处理接收到的 Ping 消息
		case websocket.PongMessage:
			// 处理接收到的 Pong 消息
		default:
			// 处理未知的消息类型
		}

	}
}

// 向指定房间中的所有连接发送消息
func sendToRoom(pool *RoomConnectionPool, roomID string, message []byte) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	connections := pool.connections[roomID]
	for _, conn := range connections {
		if err := conn.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("WebSocket write error:", err)
			continue
		}
	}
}

func main() {
	pool := NewRoomConnectionPool()

	// 处理WebSocket请求的路由
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(pool, w, r)
	})

	// 广播消息的路由
	http.HandleFunc("/broadcast", func(w http.ResponseWriter, r *http.Request) {
		roomID := r.URL.Query().Get("roomID")
		message := []byte(r.URL.Query().Get("message"))
		sendToRoom(pool, roomID, message)
	})

	// 启动服务器
	log.Println("Starting WebSocket server...")
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
