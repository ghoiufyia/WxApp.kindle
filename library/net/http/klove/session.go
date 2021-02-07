package klove

import (
	"net/http"
	"encoding/json"
	"crypto/rand"
	"encoding/hex"
	"time"
	"sync"
	"github.com/ghoiufyia/kindle/library/net/http/util/sessionhandler"
)

//session存储，在内存中保存的格式
type Session struct {
	Sid		string
	lock	sync.RWMutex
	Values	map[string]interface{}
}
//设置值
func (s *Session)Set(key string,value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Values[key] = value
	return nil
}
//读取值
func (s *Session)Get(key string) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if v, ok := s.Values[key]; ok {
		return v
	}
	return nil
}

//session manager配置
type SessionConfig struct {
	SessionIDLength 	int
	CookieLifeTime  	int64
	CookieName			string
	GcTime				int64
	Domain				string
	handlerType			string
	FilePath			string
}

//session管理器，包含配置和handler
type SessionManager struct {
	config		*SessionConfig
	handler		sessionhandler.Handler
}
//默认配置
var defaultConfig = &SessionConfig {
	SessionIDLength 	: 32,
	CookieLifeTime  	: 60*60*24*7,
	CookieName			: "KSID",
	GcTime				: 60*60*24*7,
	Domain				: ".poetnote.com",
	handlerType			: "file",
	FilePath			: "storage/session",
}
//生成新的sessin控制器
func NewSessionManager(c *SessionConfig) (s *SessionManager) {
	if c == nil {
		c = defaultConfig
	}
	var myHandler sessionhandler.Handler
	switch c.handlerType {
	case "file":
		myHandler = &sessionhandler.FileHandler {
			Lifetime:c.CookieLifeTime,
			SavePath: c.FilePath,
		}
	default :
		myHandler = &sessionhandler.FileHandler {
			Lifetime: c.CookieLifeTime,
			SavePath: c.FilePath,
		}
	}
	return &SessionManager{
		config: c,
		handler: myHandler,
	}
}
//开启session，在每次请求开始
func (s *SessionManager)SessionStart(ctx *Context) (sess *Session) {
	if sess,_ = s.readSession(ctx);sess == nil {
		sess = s.newSession()
		ctx.Session = sess
		s.setHttpCookie(ctx)
	}
	return
}
//session写入存储
func (s *SessionManager)SessionFlush(ctx *Context) (err error){
	err = s.writeSession(ctx)
	if err != nil {
		return
	}
	return
}
//设置cookie，sessionid写入cookie
func (s *SessionManager)setHttpCookie(ctx *Context) {
	cookie := &http.Cookie{
		Name:     s.config.CookieName,
		Value:    ctx.Session.Sid,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Domain:   s.config.Domain,
	}
	cookie.MaxAge = int(s.config.CookieLifeTime)
	cookie.Expires = time.Now().Add(time.Duration(s.config.CookieLifeTime) * time.Second)
	http.SetCookie(ctx.ResponseWriter, cookie)
}
//新建一个session结构
func (s *SessionManager) newSession() (sess *Session) {
	b := make([]byte,s.config.SessionIDLength)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return nil
	}
	sess = &Session{
		Sid:	hex.EncodeToString(b),
		Values:	make(map[string]interface{}),
	}
	return
}
//存储中读取session
func (s *SessionManager) readSession(ctx *Context) (sess *Session,err error) {
	cookie, err := ctx.Request.Cookie(s.config.CookieName)
	if err != nil || cookie == nil {
		return
	}
	sid := cookie.Value
	sessionValue,err := s.handler.Read(sid)
	if err != nil {
		return
	}
	var sessionMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(sessionValue),&sessionMap)
	if err != nil {
		return 
	}
	sess = &Session{
		Sid:	sid,
		Values:	sessionMap,
	}
	return
}
//sessionxi9erusession
func (s *SessionManager) writeSession(ctx *Context) (err error) {
	sid := ctx.Session.Sid
	jsonStr,err := json.Marshal(ctx.Session.Values)
	if err != nil {
		return
	}
	err = s.handler.Write(sid,string(jsonStr))
	if err != nil {
		return
	}
	return
}