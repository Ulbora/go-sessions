package gosession

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var s GoSession

func TestSession_CreateSessionStore(t *testing.T) {
	s.MaxAge = 5 * 60
	s.Name = "user-session-test"
	s.SessionKey = "554dfgdffdd11dfgf1ff1f"
	s.InitSessionStore()
	if s.store == nil {
		t.Fail()
	}
}
func TestSession(t *testing.T) {
	//var r = new(http.Request)
	r, _ := http.NewRequest("GET", "/test", nil)
	//var w http.ResponseWriter
	w := httptest.NewRecorder()
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.Values["test"] = "test_2"
	session.Values["othertest"] = 55
	session.Save(r, w)
	test1 := session.Values["test"]
	othertest := session.Values["othertest"]
	fmt.Println(test1)
	fmt.Println(othertest)
	if test1 != "test_2" || othertest != 55 {
		t.Fail()
	}
}

func TestSession_CreateSessionStoreOptions(t *testing.T) {
	var so GoSession
	so.InitSessionStore()
	if so.store == nil || so.store.Options.MaxAge != 3600 || so.Name != "user-session" {
		t.Fail()
	}
}

//go mod init github.com/Ulbora/go-sessions
