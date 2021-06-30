package feishu

import (
	"GoodGuy/src/feishu/event"
	"fmt"
	"github.com/buger/jsonparser"
)

func EventHandle(body *[]byte) {
	eventType, err := jsonparser.GetString(*body, "header", "event_type")
	if err != nil {
		fmt.Println(err)
		return
	}
	if eventType == "im.message.receive_v1" {
		event.MessageReceive(body)
	}
}
