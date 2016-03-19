package backup //import "go.iondynamics.net/siteMgrServer/backup"

import (
	"go.iondynamics.net/siteMgr"
	"go.iondynamics.net/siteMgrServer/srv"
)

var (
	Version = "0.2.0"
)

type Data struct {
	Sites       []siteMgr.Site
	Credentials []siteMgr.Credentials
	Version     string
}

func Create(usr *srv.User) Data {
	return Data{
		Sites:       usr.GetSites(),
		Credentials: usr.GetAllCredentials(),
		Version:     Version,
	}
}

func Restore(usr *srv.User, d *Data) {
	if siteMgr.AtLeast("0.2.0", d.Version) {
		for _, v := range d.Sites {
			usr.SetSite(v)
		}
		for _, v := range d.Credentials {
			usr.SetCredentials(v)
		}
	}
}
