package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

type Config struct {
	server           string
	port             uint64
	clientToken      string
	applicationToken string
}

func (c *Config) parseConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, _ := ioutil.ReadAll(file)
	t := gjson.ParseBytes(b)
	server := t.Get("server")
	if !server.Exists() {
		log.Fatal("No server url be setted. Use ws://localhost")
		c.server = "ws://localhost"
	} else {
		c.server = server.String()
	}

	port := t.Get("port")
	if !port.Exists() {
		log.Println("No server port be setted. Use 80")
		c.port = 80
	} else {
		c.port = port.Uint()
	}

	clientToken := t.Get("clientToken")
	if !clientToken.Exists() {
		log.Fatal("no clientToken")
	} else if clientToken.String() == "Please fill Your clientToken" {
		log.Fatal("Please fill Your clientToken")
	} else {
		c.clientToken = clientToken.String()
	}

	applicationToken := t.Get("applicationToken")
	if !applicationToken.Exists() {
		log.Fatal("no applicationToken")
	} else if applicationToken.String() == "Please fill Your applicationToken" {
		log.Fatal("Please fill Your applicationToken")
	} else {
		c.applicationToken = applicationToken.String()
	}
}

func connectWS(config *Config) error {
	c, _, err := websocket.DefaultDialer.Dial(config., nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("hello ithome30day"))
	if err != nil {
		log.Println(err)
		return
	}
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("receive: %s\n", msg)
}

var config Config

func init() {
	config.parseConfig()
	log.Println(config)

}

func main() {

	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		log.Panic(err)
	}
}
