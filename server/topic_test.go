package main

import (
	"testing"

	"github.com/tinode/chat/server/store/types"
)

type Messages struct {
	messages []interface{}
}

func (s *Session) testWriteLoop(results *Messages, done chan bool) {
	for msg := range s.send {
		results.messages = append(results.messages, msg)
	}
	done <- true
}

func TestReplyGetDescInvalidOpts(t *testing.T) {
	var sess Session
	sess.send = make(chan interface{}, 10)
	writeLoopDone := make(chan bool)
	var responses Messages
	go sess.testWriteLoop(&responses, writeLoopDone)

	topic := Topic{
		cat: types.TopicCatMe,
	}

	msg := ClientComMessage{
		Original: "dummy",
	}
	// Can't specify User in opts.
	if err := topic.replyGetDesc(&sess, 123, &MsgGetOpts{User: "abcdef"}, &msg); err == nil {
		t.Error("replyGetDesc expected to error out.")
	} else if err.Error() != "invalid GetDesc query" {
		t.Errorf("Unexpected error: expected 'invalid GetDesc query', got '%s'", err.Error())
	}
	close(sess.send)
	// Wait for the session's write loop to complete.
	<-writeLoopDone
	if len(responses.messages) != 1 {
		t.Fatalf("`responses` expected to contain 1 element, found %d", len(responses.messages))
	}
	resp := responses.messages[0].(*ServerComMessage)
	if resp.Ctrl == nil {
		t.Fatalf("response expected to contain a Ctrl message")
	}
	if resp.Ctrl.Code != 400 {
		t.Errorf("response code: expected 400, found: %d", resp.Ctrl.Code)
	}
}
