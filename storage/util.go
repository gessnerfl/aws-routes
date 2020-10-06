package storage

import (
	"os/user"

	log "github.com/sirupsen/logrus"
)

//GetUserHome returns the home directory of the user
func GetUserHome() string {
	u, err := user.Current()
	if err != nil {
		log.Warnf("Cannot determine user home directory. Fall back to /tmp; message = '%s'", err.Error())
		return "/tmp"
	}
	return u.HomeDir
}
