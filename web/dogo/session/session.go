package session

import (
	// "github.com/gorilla/sessions"
	"sync"
	"net/http"
	"net/url"
	"github.com/RichardKnop/uuid"
	"github.com/ghoiufyia/WxApp.kindle/web/until/config"
)

type Config struct {
	cookieName	string
	lifetime	int64
	handlerType	string
	filePath	string
}

type Manager struct {
	config		*Config
	lock		sync.Mutex
	handle		handler
}
var lock sync.Mutex

var defaultConfig = Config {
	cookieName:"sid_",
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
		myHandler = &FileSession {
			savePath: c.filePath,
		}
	default :
	myHandler = &FileSession {
			savePath: "storage/session",
		}
	}

	return &Manager{
		config: &c,
		handle: myHandler,
	}
}

func (m *Manager)getSid(r *http.Request) (string,error) {
	cookie, err := r.Cookie(m.config.cookieName)
	if err != nil || cookie.Value == "" {
		return "", nil
	}
	return url.QueryUnescape(cookie.Value)
}

func (m *Manager)SessionStart(w http.ResponseWriter, r *http.Request) (*Store,error) {
	sid,err := m.getSid(r)
	if err != nil {
		return nil,err
	}
	if sid == "" {
		sid = m.SeesionID()
	}
	 m.handle.Read(sid)
	 

	return nil,err
}

func (m *Manager) SeesionID() (string) {
	return uuid.New()
}




// SetSessionService(r *http.Request, w http.ResponseWriter)
