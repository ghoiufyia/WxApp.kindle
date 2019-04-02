package session

import (
	// "github.com/gorilla/sessions"
)

type SessionManager struct {
	Handler		handler
}

func NewManager() *SessionManager {
	return &SessionManager{
		Handler:	&FileSession {
			lifetime: 60*60*24,
			savePath: "storage/session",
		},
	}
}


// SetSessionService(r *http.Request, w http.ResponseWriter)
