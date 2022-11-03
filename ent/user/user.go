// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgePlayers holds the string denoting the players edge name in mutations.
	EdgePlayers = "players"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "master_name"
	// PlayersTable is the table that holds the players relation/edge.
	PlayersTable = "players"
	// PlayersInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayersInverseTable = "players"
	// PlayersColumn is the table column denoting the players relation/edge.
	PlayersColumn = "user_players"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldUsername,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)