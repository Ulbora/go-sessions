package gosession

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
)

var s GoSession
var store *sessions.CookieStore

func TestSession_CreateSessionStore(t *testing.T) {
	s.MaxAge = 5 * 60
	s.Name = "user-session-test"
	s.SessionKey = "554dfgdffdd11dfgf1ff1f"
	store = s.InitSessionStore()
	if store == nil {
		t.Fail()
	}
}
func TestSession(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	session, err := store.Get(r, s.Name)
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
	store2 := so.InitSessionStore()
	if store2 == nil || store2.Options.MaxAge != 3600 || so.Name != "user-session" {
		t.Fail()
	}
}

//go mod init github.com/Ulbora/go-sessions
