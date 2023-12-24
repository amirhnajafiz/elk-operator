package query

import (
	"fmt"
	"strings"

	"github.com/amirhnajafiz/elk-operator/pkg/crypto"
)

type UsersQuery struct {
	Username string
	Password string
	Roles    []string
	Clusters []string
}

func (u *UsersQuery) Table() string {
	return "users"
}

func (u *UsersQuery) queryInsert() string {
	return fmt.Sprintf(
		"INSERT INTO %s (username, password, roles, clusters) VALUES ('%s', '%s', '%s', '%s');",
		u.Table(),
		u.Username,
		crypto.Hash(u.Password),
		strings.Join(u.Roles, ";"),
		strings.Join(u.Clusters, ";"),
	)
}

func (u *UsersQuery) queryUpdate(old *UsersQuery) string {
	return fmt.Sprintf(
		"UPDATE %s SET username='%s', roles='%s', clusters='%s' WHERE username='%s';",
		u.Table(),
		u.Username,
		strings.Join(u.Roles, ";"),
		strings.Join(u.Clusters, ";"),
		old.Username,
	)
}

func (u *UsersQuery) queryDelete() string {
	return fmt.Sprintf("DELETE FROM %s WHERE username='%s'", u.Table(), u.Username)
}

func (u *UsersQuery) querySelect() string {
	return fmt.Sprintf("SELECT * FROM %s WHERE username='%s'", u.Table(), u.Username)
}
