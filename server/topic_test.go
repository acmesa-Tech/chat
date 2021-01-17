package main

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/tinode/chat/server/auth"
	"github.com/tinode/chat/server/store"
	"github.com/tinode/chat/server/store/types"
)

// Dummy DB adapter.
type dummyAdapter struct {
	maxResults int
}

func (a *dummyAdapter) GetName() string {
	return "dummy"
}

func (a *dummyAdapter) Open(config json.RawMessage) error                             { return nil }
func (a *dummyAdapter) Close() error                                                  { return nil }
func (a *dummyAdapter) IsOpen() bool                                                  { return false }
func (a *dummyAdapter) GetDbVersion() (int, error)                                    { return 1, nil }
func (a *dummyAdapter) CheckDbVersion() error                                         { return nil }
func (a *dummyAdapter) SetMaxResults(val int) error                                   { return nil }
func (a *dummyAdapter) CreateDb(reset bool) error                                     { return nil }
func (a *dummyAdapter) UpgradeDb() error                                              { return nil }
func (a *dummyAdapter) Version() int                                                  { return 1 }
func (a *dummyAdapter) UserCreate(user *types.User) error                             { return nil }
func (a *dummyAdapter) UserGet(uid types.Uid) (*types.User, error)                    { return nil, nil }
func (a *dummyAdapter) UserGetAll(ids ...types.Uid) ([]types.User, error)             { return []types.User{}, nil }
func (a *dummyAdapter) UserDelete(uid types.Uid, hard bool) error                     { return nil }
func (a *dummyAdapter) UserUpdate(uid types.Uid, update map[string]interface{}) error { return nil }
func (a *dummyAdapter) UserUpdateTags(uid types.Uid, add, remove, reset []string) ([]string, error) {
	return []string{}, nil
}
func (a *dummyAdapter) UserGetByCred(method, value string) (types.Uid, error) { return 1, nil }
func (a *dummyAdapter) UserUnreadCount(uid types.Uid) (int, error)            { return 1, nil }
func (a *dummyAdapter) CredUpsert(cred *types.Credential) (bool, error)       { return true, nil }
func (a *dummyAdapter) CredGetActive(uid types.Uid, method string) (*types.Credential, error) {
	return nil, nil
}
func (a *dummyAdapter) CredGetAll(uid types.Uid, method string, validatedOnly bool) ([]types.Credential, error) {
	return []types.Credential{}, nil
}
func (a *dummyAdapter) CredDel(uid types.Uid, method, value string) error { return nil }
func (a *dummyAdapter) CredConfirm(uid types.Uid, method string) error    { return nil }
func (a *dummyAdapter) CredFail(uid types.Uid, method string) error       { return nil }
func (a *dummyAdapter) AuthGetUniqueRecord(unique string) (types.Uid, auth.Level, []byte, time.Time, error) {
	return 1, auth.LevelAuth, []byte{}, time.Now(), nil
}
func (a *dummyAdapter) AuthGetRecord(user types.Uid, scheme string) (string, auth.Level, []byte, time.Time, error) {
	return "", auth.LevelAuth, []byte{}, time.Now(), nil
}
func (a *dummyAdapter) AuthAddRecord(user types.Uid, scheme, unique string, authLvl auth.Level, secret []byte, expires time.Time) error {
	return nil
}
func (a *dummyAdapter) AuthDelScheme(user types.Uid, scheme string) error { return nil }
func (a *dummyAdapter) AuthDelAllRecords(uid types.Uid) (int, error)      { return 0, nil }
func (a *dummyAdapter) AuthUpdRecord(user types.Uid, scheme, unique string, authLvl auth.Level, secret []byte, expires time.Time) error {
	return nil
}
func (a *dummyAdapter) TopicCreate(topic *types.Topic) error                        { return nil }
func (a *dummyAdapter) TopicCreateP2P(initiator, invited *types.Subscription) error { return nil }
func (a *dummyAdapter) TopicGet(topic string) (*types.Topic, error)                 { return nil, nil }
func (a *dummyAdapter) TopicsForUser(uid types.Uid, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	return []types.Subscription{}, nil
}
func (a *dummyAdapter) UsersForTopic(topic string, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	return []types.Subscription{}, nil
}
func (a *dummyAdapter) OwnTopics(uid types.Uid) ([]string, error)                     { return []string{}, nil }
func (a *dummyAdapter) TopicShare(subs []*types.Subscription) error                   { return nil }
func (a *dummyAdapter) TopicDelete(topic string, hard bool) error                     { return nil }
func (a *dummyAdapter) TopicUpdateOnMessage(topic string, msg *types.Message) error   { return nil }
func (a *dummyAdapter) TopicUpdate(topic string, update map[string]interface{}) error { return nil }
func (a *dummyAdapter) TopicOwnerChange(topic string, newOwner types.Uid) error       { return nil }
func (a *dummyAdapter) SubscriptionGet(topic string, user types.Uid) (*types.Subscription, error) {
	return nil, nil
}
func (a *dummyAdapter) SubsForUser(user types.Uid, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	return []types.Subscription{}, nil
}
func (a *dummyAdapter) SubsForTopic(topic string, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	return []types.Subscription{}, nil
}
func (a *dummyAdapter) SubsUpdate(topic string, user types.Uid, update map[string]interface{}) error {
	return nil
}
func (a *dummyAdapter) SubsDelete(topic string, user types.Uid) error  { return nil }
func (a *dummyAdapter) SubsDelForTopic(topic string, hard bool) error  { return nil }
func (a *dummyAdapter) SubsDelForUser(user types.Uid, hard bool) error { return nil }
func (a *dummyAdapter) FindUsers(user types.Uid, req [][]string, opt []string) ([]types.Subscription, error) {
	return []types.Subscription{}, nil
}
func (a *dummyAdapter) FindTopics(req [][]string, opt []string) ([]types.Subscription, error) {
	return []types.Subscription{}, nil
}
func (a *dummyAdapter) MessageSave(msg *types.Message) error { return nil }
func (a *dummyAdapter) MessageGetAll(topic string, forUser types.Uid, opts *types.QueryOpt) ([]types.Message, error) {
	return []types.Message{}, nil
}
func (a *dummyAdapter) MessageDeleteList(topic string, toDel *types.DelMessage) error { return nil }
func (a *dummyAdapter) MessageGetDeleted(topic string, forUser types.Uid, opts *types.QueryOpt) ([]types.DelMessage, error) {
	return []types.DelMessage{}, nil
}
func (a *dummyAdapter) MessageAttachments(msgId types.Uid, fids []string) error { return nil }
func (a *dummyAdapter) DeviceUpsert(uid types.Uid, dev *types.DeviceDef) error  { return nil }
func (a *dummyAdapter) DeviceGetAll(uid ...types.Uid) (map[types.Uid][]types.DeviceDef, int, error) {
	return map[types.Uid][]types.DeviceDef{}, 0, nil
}
func (a *dummyAdapter) DeviceDelete(uid types.Uid, deviceID string) error { return nil }
func (a *dummyAdapter) FileStartUpload(fd *types.FileDef) error           { return nil }
func (a *dummyAdapter) FileFinishUpload(fid string, status int, size int64) (*types.FileDef, error) {
	return nil, nil
}
func (a *dummyAdapter) FileGet(fid string) (*types.FileDef, error) { return nil, nil }
func (a *dummyAdapter) FileDeleteUnused(olderThan time.Time, limit int) ([]string, error) {
	return []string{}, nil
}

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
	uid1 := types.Uid(1)
	uid2 := types.Uid(2)

	topic := Topic{
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

func TestMain(m *testing.M) {
	store.RegisterAdapter(&dummyAdapter{maxResults: 100})
	storeConfig := []byte(`{"use_config": "dummy",
                          "uid_key": "la6YsO+bNX/+XIkOqc5Svw==" }`)
	if err := store.Open(0, storeConfig); err != nil {
		log.Fatal("Failed to initialize data store:", err)
	}

	m.Run()
}
