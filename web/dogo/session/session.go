package session

import (
	// "github.com/gorilla/sessions"
	"sync"
)

type Manager struct {
	cookieName	string
	lock		sync.Mutex
	lifetime	int64
	Handler		handler
}
var lock sync.Mutex
var DefaultManager = Manager{
	cookieName:"dogo_sid",
	lifetime: 60*60*24,
	Handler:	&FileSession {
		savePath: "storage/session",
	},
}

func NewManager() *Manager {
	return &Manager{
		cookieName: "dogo_sid",
		lifetime: 60*60*24,
		Handler:	&FileSession {
			savePath: "storage/session",
		},
	}
}

type Session interface {
	sessionId() (string)
}



// SetSessionService(r *http.Request, w http.ResponseWriter)
