# 💬 Live Chat Application

A real-time chat application built with **Go (Golang)** for the backend and **React.js** for the frontend using **WebSockets**.

## 🚀 Features

- 🔗 Real-time bi-directional messaging using WebSocket
- 📦 Organized backend with Gorilla WebSocket
- ⚛️ Responsive React frontend
- 🔒 Graceful connection handling (connect/disconnect)
- 💬 Broadcast messages to all connected clients

---

## 🛠️ Technologies Used

### 🔧 Backend (Golang)
- Go (Golang)
- Gorilla WebSocket
- Custom WebSocket Pool & Client handling

### 🎨 Frontend (React)
- React.js 
- WebSocket

---

## 📂 Project Structure
/backend
└── main.go
└── coustomWebsocket/
├── client.go
├── pool.go
└── ...

/frontend
└── src/
├── App.js
