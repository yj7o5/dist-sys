package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Echo struct {
	Type      string `json:"type,omitempty"`
	MsgID     int    `json:"msg_id,omitempty"`
	InReplyTo int    `json:"in_reply_to,omitempty"`
	Echo      string `json:"echo,omit_empty"`
}

func main() {
	node := maelstrom.NewNode()
	node.Handle("echo", func(msg maelstrom.Message) error {
		return handleEchoMessage(node, msg)
	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}

func handleEchoMessage(node *maelstrom.Node, msg maelstrom.Message) error {
	req := Echo{}
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	body := Echo{
		Type:      "echo_ok",
		MsgID:     req.MsgID,
		InReplyTo: req.InReplyTo,
		Echo:      req.Echo,
	}
	return node.Reply(msg, body)
}
