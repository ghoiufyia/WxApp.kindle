package session

import (
	// "github.com/gorilla/sessions"
	"sync"
	"net/http"
	"net/url"
	// "bytes"
	"encoding/json"
	"github.com/RichardKnop/uuid"
	"github.com/ghoiufyia/WxApp.kindle/web/until/config"
	// "fmt"
)

var (
	Session		*Manager
)

func init() {
	Session = NewManager("config/session.json")
}

type Config struct {
	cookieName	string
	lifetime	int64
	handlerType	string
	filePath	string
}

type Manager struct {
	config		*Config
	lock		sync.Mutex
	handler		handler
}

var defaultConfig = Config {
	cookieName:"sid",
	lifetime: 60*60*24,
	handlerType: "file",
	filePath: "storage/session",
}

func NewManager(cf string) *Manager {
	var c Config
	err := config.ParseFile(&c,cf)
	if err != nil {
		c = defaultConfig
	}
	var myHandler handler
	switch c.handlerType {
	case "file":
		myHandler = &FileHandler {
			savePath: c.filePath,
		}
	default :
		myHandler = &FileHandler {
			savePath: c.filePath,
		}
	}

	return &Manager{
		config: &c,
		handler: myHandler,
	}
}

func (m *Manager)getSid(r *http.Request) (string,error) {
	cookie, err := r.Cookie(m.config.cookieName)
	if err != nil || cookie.Value == "" {
		return "", nil
	}
	return url.QueryUnescape(cookie.Value)
}

func (m *Manager)SessionStart(w http.ResponseWriter, r *http.Request) (session *Store,err error) {
	sid,err := m.getSid(r)
	if err != nil {
		return nil,err
	}
	if sid == "" {
		sid = m.SeesionID()
	}
	sessionValue,err := m.handler.Read(sid)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     m.config.cookieName,
		Value:    url.QueryEscape(sid),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Domain:   ".go.poetnoe.com",
	}
	// if m.config.EnableSetCookie {
		http.SetCookie(w, cookie)
	// }
	r.AddCookie(cookie)

	var sessionMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(sessionValue),&sessionMap)
	return &Store{
		sid: sid,
		values: sessionMap,
		handle: m.handler,
	},nil

	return nil,err
}

func (m *Manager) SeesionID() (string) {
	return uuid.New()
}




// SetSessionService(r *http.Request, w http.ResponseWriter)
