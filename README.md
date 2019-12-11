go-sessions 
==============

[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/go-sessions)](https://goreportcard.com/report/github.com/Ulbora/go-sessions)

A Go wrapper for gorilla sessions
http://www.gorillatoolkit.org/pkg/sessions

# Installation

```
$ go get github.com/gorilla/sessions
$ go get github.com/Ulbora/go-sessions

```

# Usage

## Initialize and use
```
    import usess github.com/Ulbora/go-sessions
    var s usess.Session
    func main() {
	    s.MaxAge = 5 * 60
	    s.Name = "user-session-test"
	    s.SessionKey = "554dfgdffdd11dfgf1ff1f"
    }

    func handleSomething(r http.ResponseWriter, w *http.Request) {
	    s.InitSessionStore()
        session, err := s.GetSession(r)
	    if err != nil {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
	    }

        //get already set session value        
	    user := session.Values["username"]

        //sets new session values
        session.Values["someVal"] = "someValue"
        session.Values["someOtherVal"] = 55
        //saves new session values
        session.Save(r, w)

        //do something
    }
```

# Important Note:
If you aren't using gorilla/mux, you need to wrap your handlers with context.ClearHandler or else you will leak memory! An easy way to do this is to wrap the top-level mux when calling http.ListenAndServe: