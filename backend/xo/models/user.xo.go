// Package models contains the types for schema 'demo'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// User represents a row from 'demo.users'.
type User struct {
	ID          int       `json:"id"`           // id
	Email       string    `json:"email"`        // email
	DisplayName string    `json:"display_name"` // display_name
	CreatedAt   time.Time `json:"created_at"`   // created_at
	UpdatedAt   time.Time `json:"updated_at"`   // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO demo.users (` +
		`email, display_name, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.Email, u.DisplayName, u.CreatedAt, u.UpdatedAt)
	res, err := db.Exec(sqlstr, u.Email, u.DisplayName, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = int(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE demo.users SET ` +
		`email = ?, display_name = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.Email, u.DisplayName, u.CreatedAt, u.UpdatedAt, u.ID)
	_, err = db.Exec(sqlstr, u.Email, u.DisplayName, u.CreatedAt, u.UpdatedAt, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM demo.users WHERE id = ?`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserByEmail retrieves a row from 'demo.users' as a User.
//
// Generated from index 'email'.
func UserByEmail(db XODB, email string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, email, display_name, created_at, updated_at ` +
		`FROM demo.users ` +
		`WHERE email = ?`

	// run query
	XOLog(sqlstr, email)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, email).Scan(&u.ID, &u.Email, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserByID retrieves a row from 'demo.users' as a User.
//
// Generated from index 'users_id_pkey'.
func UserByID(db XODB, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, email, display_name, created_at, updated_at ` +
		`FROM demo.users ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Email, &u.DisplayName, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
