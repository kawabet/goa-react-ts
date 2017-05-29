// Package models contains the types for schema 'goa_chat'.
package models

// GENERATED BY XO. DO NOT EDIT.

// NewRoom calls the stored procedure 'goa_chat.new_room(varchar(256), varchar(400)) int(11)' on db.
func NewRoom(db XODB, v0 string, v1 string) (int, error) {
	var err error

	// sql query
	const sqlstr = `SELECT goa_chat.new_room(?, ?)`

	// run query
	var ret int
	XOLog(sqlstr, v0, v1)
	err = db.QueryRow(sqlstr, v0, v1).Scan(&ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}
