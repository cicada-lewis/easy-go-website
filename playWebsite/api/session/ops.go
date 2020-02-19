package session


import (
	"github.com/Catelemmon/easy-go-website/playWebsite/api/dbops"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/defs"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/utils"
	"sync"
	"time"
)


var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	res, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	res.Range(func(k, v interface{}) bool {
		ss, ok := v.(*defs.SimpleSession)
		if !ok {
			return false
		}
		sessionMap.Store(k, ss)
		return true
	})

}

func GenerateNewSessionId(uname string) string {
	sid, _ := utils.NewUUID()
	ct := time.Now().UnixNano() / 100000
	ttl := ct + 30 * 60 * 1000// service session expired: 30min

	ss := &defs.SimpleSession{
		Username:uname,
		TTL: ttl,
	}
	sessionMap.Store(sid, ss)
	dbops.InsertSession(sid, ttl, uname)
	return sid
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano() / 100000
		if ss.(*defs.SimpleSession) .TTL <= ct {
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}

func deleteExpiredSession(sid string) error {
	sessionMap.Delete(sid)
	err := dbops.DeleteSession(sid)
	if err != nil {
		return err
	}
	return nil
}

