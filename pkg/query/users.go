package query

import (
	"fmt"
	"strings"
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

func (u *UsersQuery) QueryInsert() string {
	return fmt.Sprintf(
		"INSERT INTO %s (username, password, roles, clusters) VALUES ('%s', '%s', '%s', '%s');",
		u.Table(),
		u.Username,
		u.Password,
		strings.Join(u.Roles, ";"),
		strings.Join(u.Clusters, ";"),
	)
}

func (u *UsersQuery) QueryDelete() string {
	return fmt.Sprintf("DELETE FROM %s WHERE username='%s'", u.Table(), u.Username)
}
