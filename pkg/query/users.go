package query

import (
	"fmt"
	"strings"
)

type UsersQuery struct {
	Username string
	Password string
	Cluster  string
	Roles    []string
}

func (u *UsersQuery) Table() string {
	return "users"
}

func (u *UsersQuery) QueryInsert() string {
	return fmt.Sprintf(
		"INSERT INTO %s (username, password, roles, cluster) VALUES ('%s', '%s', '%s', '%s');",
		u.Table(),
		u.Username,
		u.Password,
		strings.Join(u.Roles, ";"),
		u.Cluster,
	)
}

func (u *UsersQuery) QueryDelete() string {
	return fmt.Sprintf("UPDATE %s SET deleted_at = now() WHERE username='%s'", u.Table(), u.Username)
}
