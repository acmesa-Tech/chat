package main

import (
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tinode/chat/server/db/mock_adapter"
	"github.com/tinode/chat/server/store"
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

func (h *Hub) testHubLoop(done chan bool) {
	for msg := range h.route {
		log.Printf("hub.route: %+v", msg)
	}
	done <- true
}

func TestHandleBroadcastP2P(t *testing.T) {
	uid1 := types.Uid(1)
	uid2 := types.Uid(2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	a := mock_adapter.NewMockAdapter(ctrl)

	a.EXPECT().GetName().Return("dummy")
	a.EXPECT().IsOpen().Return(false)
	a.EXPECT().Open(gomock.Any())
	a.EXPECT().SetMaxResults(gomock.Any())
	a.EXPECT().CheckDbVersion()
	a.EXPECT().TopicUpdateOnMessage("p2p-test", gomock.Any()).Return(nil)
	a.EXPECT().MessageSave(gomock.Any()).Return(nil)
	a.EXPECT().SubsUpdate("p2p-test", uid1, gomock.Any())

	store.RegisterAdapter(a)
	storeConfig := []byte(`{"use_config": "dummy",
                          "uid_key": "la6YsO+bNX/+XIkOqc5Svw==" }`)
	if err := store.Open(0, storeConfig); err != nil {
		log.Fatal("Failed to initialize data store:", err)
	}

	ss := make([]*Session, 2)
	done := make([]chan bool, 2)
	messages := make([]*Messages, 2)
	for i := range ss {
		ss[i] = &Session{sid: "sid" + string(i)}
		ss[i].send = make(chan interface{}, 1)
		done[i] = make(chan bool)
		messages[i] = &Messages{}
		go ss[i].testWriteLoop(messages[i], done[i])
	}

	h := &Hub{
		route: make(chan *ServerComMessage, 10),
	}
	globals.hub = h
	hubDone := make(chan bool)
	go h.testHubLoop(hubDone)

	topic := Topic{
		name:   "p2p-test",
		cat:    types.TopicCatP2P,
		status: topicStatusLoaded,
		perUser: map[types.Uid]perUserData{
			uid1: perUserData{
				modeWant:  types.ModeCP2P,
				modeGiven: types.ModeCP2P,
			},
			uid2: perUserData{
				modeWant:  types.ModeCP2P,
				modeGiven: types.ModeCP2P,
			},
		},
		isProxy: false,
		sessions: map[*Session]perSessionData{
			ss[0]: perSessionData{uid: uid1},
			ss[1]: perSessionData{uid: uid2},
		},
	}
	msg := &ServerComMessage{
		AsUser: uid1.UserId(),
		//Original: "dummy",
		Data: &MsgServerData{
			Topic:   "p2p",
			From:    uid1.UserId(),
			Content: "test",
		},
		sess:    ss[0],
		SkipSid: ss[0].sid,
	}
	log.Println("as user = ", msg.AsUser)
	topic.handleBroadcast(msg)
	for _, s := range ss {
		close(s.send)
	}
	for i, _ := range ss {
		// Wait for the session's write loop to complete.
		<-done[i]
	}
	// Hub loop.
	close(h.route)
	<-hubDone
	// Message uid1 -> uid2.
	for i, m := range messages {
		if i == 0 {
			if len(m.messages) != 0 {
				t.Fatalf("Uid1: expected 0 messages, got %d", len(m.messages))
			}
		} else {
			if len(m.messages) != 1 {
				t.Fatalf("Uid2: expected 1 messages, got %d", len(m.messages))
			}
			r := m.messages[0].(*ServerComMessage)
			if r.Data == nil {
				t.Fatalf("Response[0] must have a ctrl message")
			}
			if r.Data.Content.(string) != "test" {
				t.Errorf("Response[0] content: expected 'test', got '%s'", r.Data.Content.(string))
			}
		}
	}
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
