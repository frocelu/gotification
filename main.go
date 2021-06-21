package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func (c *Config) genWsUri() string {
	wsUri := ""
	if c.applicationToken != "" {
		wsUri = fmt.Sprintf("%v:%v/stream", config.server, config.port)
	} else {
		wsUri = fmt.Sprintf("%v:%v/stream", config.server, config.port)
	}
	fmt.Println(wsUri)
	return wsUri
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
		c.applicationToken = ""
	} else if applicationToken.String() == "Please fill Your applicationToken, or remove this key to get all messages" {
		c.applicationToken = ""
	} else {
		c.applicationToken = applicationToken.String()
	}
}

func connectWS(config *Config) error {
	header := http.Header{}
	header.Set("X-Gotify-Key", config.clientToken)
	c, _, err := websocket.DefaultDialer.Dial(config.genWsUri(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return err
	}
	log.Printf("receive: %s\n", msg)
	return nil
}

var config Config

func init() {
	config.parseConfig()
	log.Println(config)

}

func main() {
	connectWS(&config)
	// err := beeep.Notify("gotification", "Title", "測試", "assets/information.png")
	// if err != nil {
	// 	log.Panic(err)
	// }
	err := beeep.Alert("gotification", "Title", "測試", "assets/information.png")
	if err != nil {
		log.Panic(err)
	}
}
