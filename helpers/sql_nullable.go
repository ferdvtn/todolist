package helpers

import "database/sql"

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  len(s) > 0,
	}
}
