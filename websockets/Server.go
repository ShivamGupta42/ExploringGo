package websockets


import (
	"log"
)

type message struct {
	// the json tag means this will serialize as a lowercased field
	Message string `json:"message"`
}
func socket(ws *Conn) {
	for {
		// allocate our container struct
		var m message
		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &amp;m); err != nil {
			log.Println(err)
			break
		}
		log.Println("Received message:", m.Message)
		// send a response
		m2 := message{"Thanks for the message!"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}
	}
}
