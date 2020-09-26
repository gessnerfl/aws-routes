package storage

import "os/user"

//GetUserHome returns the home directory of the user
func GetUserHome() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}
