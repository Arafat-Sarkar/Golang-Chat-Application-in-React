package main

import (
	coustomWebsocket "chatApplication/websocket"
	"fmt"
	"net/http"
)

func webSocket( pool *coustomWebsocket.Pool, w http.ResponseWriter , r *http.Request){
	
     conn , err:=coustomWebsocket.Upgrade(w,r)
	 if err != nil{
		fmt.Println(err)
		return
	 }
	 client := &coustomWebsocket.Client{
		Conn: conn,
		Pool: pool,
	 }
	 pool.Register <- client
	 client.Read()
}

func setupRutes(){
	fmt.Println("This is Working")
	pool:=coustomWebsocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		webSocket(pool, w,r)
	})
}


func main() {
	setupRutes()
 http.ListenAndServe(":8080",nil)
}