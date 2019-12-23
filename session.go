package gosession

import (
	"log"

	gs "github.com/gorilla/sessions"
)

//GoSession a GoSession
type GoSession struct {
	Path       string
	MaxAge     int
	HTTPOnly   bool
	Secure     bool
	Name       string
	SessionKey string
}

//InitSessionStore initialize session store
func (s *GoSession) InitSessionStore() *gs.CookieStore {
	if s.Path == "" {
		log.Println("using defalut path of /")
		s.Path = "/"
	}
	if s.MaxAge == 0 {
		log.Println("using defalut max age")
		s.MaxAge = 3600 //default 3600 seconds --  1 hour
	}
	if s.Name == "" {
		log.Println("using defalut name")
		s.Name = "user-session"
	}
	if s.SessionKey == "" {
		log.Println("using defalut sesstion key")
		s.SessionKey = "554dfgdffdd11dfgf1ff1f" // default key
	}
	store := s.createSessionStore()
	return store
}

// CreateSessionStore creates a sesstion
func (s *GoSession) createSessionStore() *gs.CookieStore {
	store := gs.NewCookieStore([]byte(s.SessionKey))
	store.Options = &gs.Options{
		Path:     s.Path,
		MaxAge:   s.MaxAge,
		HttpOnly: s.HTTPOnly,
		Secure:   s.Secure,
	}
	return store
}
