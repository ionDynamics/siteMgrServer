package session //import "go.iondynamics.net/siteMgrServer/session"

import (
	"sync"
	"time"

	"go.iondynamics.net/iDhelper/randGen"

	"go.iondynamics.net/siteMgrServer/srv"
)

var m sync.RWMutex
var t map[string]string = make(map[string]string)
var n map[string]*srv.User = make(map[string]*srv.User)

func Start(usr *srv.User) string {
	k := randGen.String(64)
	Set(k, usr)
	Timeout(k, 10*time.Hour)
	return k
}

func Get(key string) *srv.User {
	m.RLock()
	defer m.RUnlock()

	name, ok := t[key]
	if !ok {
		return nil
	}

	return n[name]
}

func Sync(usr *srv.User) {
	usr2 := GetByName(usr.Name)
	if usr2 == nil {
		SetUser(usr)
	} else {
		*usr = *usr2
	}
}

func GetByName(name string) *srv.User {
	m.RLock()
	defer m.RUnlock()

	usr, ok := n[name]
	if !ok {
		return nil
	}

	return usr
}

func GetKeyByName(name string) string {
	m.RLock()
	defer m.RUnlock()

	for key, uname := range t {
		if uname == name {
			return key
		}
	}
	return ""
}

func Set(key string, usr *srv.User) {
	m.Lock()
	defer m.Unlock()

	usr, ok := n[usr.Name]
	if !ok {
		n[usr.Name] = usr
	}
	t[key] = usr.Name

}

func SetUser(usr *srv.User) {
	m.Lock()
	defer m.Unlock()

	n[usr.Name] = usr
}

func Del(key string) {
	m.Lock()
	defer m.Unlock()

	delete(t, key)
}

func Timeout(key string, dur time.Duration) {
	go func(k string, d time.Duration) {
		<-time.After(d)
		Del(k)
	}(key, dur)
}
