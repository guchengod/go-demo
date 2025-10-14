package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 1. 定义 User 结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 2. 创建一个内存 "数据库"
type UserStore struct {
	// sync.RWMutex 是一个读写锁，允许多个读操作同时进行，但写操作是独占的
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

// 3. 创建一个作为我们 API 处理器的结构体
// 将 UserStore 作为其一个字段，这样处理器就可以访问数据库了
type UserHandler struct {
	store *UserStore
}

// 4. 实现 http.Handler 接口
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/users":
		h.getUsers(w, r)
	// --- 新增的 case ---
	case r.Method == http.MethodPost && r.URL.Path == "/users":
		h.createUser(w, r)
	default:
		http.NotFound(w, r)
	}
}

// getUsers 处理 GET /users 请求
// 它是 *UserHandler 的一个方法
func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	// 1. 加读锁，因为我们要读取 users map。当函数返回时，自动解锁。
	h.store.mu.RLock()
	defer h.store.mu.RUnlock()

	// 2. 将 map 转换为 slice，因为 JSON 更适合表示数组
	var userList []User
	for _, user := range h.store.users {
		userList = append(userList, user)
	}

	// 3. 设置响应头的 Content-Type 为 application/json
	w.Header().Set("Content-Type", "application/json")

	// 4. 使用 json.NewEncoder 将 userList 编码为 JSON 并直接写入 ResponseWriter
	// 这比先 Marshal 到内存再写入更高效
	err := json.NewEncoder(w).Encode(userList)
	if err != nil {
		// 如果编码失败，向客户端返回一个服务器内部错误
		http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
	}
}

// createUser 处理 POST /users 请求
func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	// 1. 解码请求体中的 JSON
	var user User
	// json.NewDecoder(r.Body) 创建一个解码器
	// .Decode(&user) 将 JSON 解码到 user 变量的地址中
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// 如果 JSON 格式错误，返回 400 Bad Request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. 简单的验证
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// 3. 加写锁，准备修改共享数据
	h.store.mu.Lock()
	defer h.store.mu.Unlock()

	// 4. 分配新 ID 并存储
	user.ID = h.store.nextID
	h.store.users[user.ID] = user
	h.store.nextID++

	// 5. 设置响应头和状态码，并将创建好的用户返回
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 设置状态码为 201
	json.NewEncoder(w).Encode(user)
}

func main() {
	// 初始化我们的数据库
	store := &UserStore{
		users:  make(map[int]User),
		nextID: 1,
	}
	// 添加一些初始数据
	store.users[store.nextID] = User{ID: store.nextID, Name: "Alice", Age: 30}
	store.nextID++
	store.users[store.nextID] = User{ID: store.nextID, Name: "Bob", Age: 25}
	store.nextID++

	// 初始化我们的处理器
	userHandler := &UserHandler{
		store: store,
	}

	fmt.Println("API 服务器即将启动，请访问 http://localhost:8080/users")

	// 启动服务器，将所有请求都交给我们的 userHandler 处理
	err := http.ListenAndServe(":8080", userHandler)
	if err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
