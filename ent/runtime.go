// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ent-demo/ent/pet"
	"ent-demo/ent/player"
	"ent-demo/ent/schema"
	"ent-demo/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescCreatedAt is the schema descriptor for created_at field.
	petDescCreatedAt := petFields[2].Descriptor()
	// pet.DefaultCreatedAt holds the default value on creation for the created_at field.
	pet.DefaultCreatedAt = petDescCreatedAt.Default.(func() time.Time)
	playerFields := schema.Player{}.Fields()
	_ = playerFields
	// playerDescCreatedAt is the schema descriptor for created_at field.
	playerDescCreatedAt := playerFields[1].Descriptor()
	// player.DefaultCreatedAt holds the default value on creation for the created_at field.
	player.DefaultCreatedAt = playerDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
