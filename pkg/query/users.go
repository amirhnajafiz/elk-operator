package query

import "fmt"

const UsersTable = "users"

func queryInsertUser(username, password, roles, clusters string) string {
	return fmt.Sprintf(
		"INSERT INTO %s (username, password, roles, clusters) VALUES ('%s', '%s', '%s', '%s');",
		UsersTable,
		username,
		password,
		roles,
		clusters,
	)
}

func queryUpdateUser(old, username, roles, clusters string) string {
	return fmt.Sprintf(
		"UPDATE %s SET username='%s', roles='%s', clusters='%s' WHERE username='%s';",
		UsersTable,
		username,
		roles,
		clusters,
		old,
	)
}

func queryDeleteUser(username string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE username='%s'", UsersTable, username)
}
