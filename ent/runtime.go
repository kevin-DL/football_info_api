// Code generated by ent, DO NOT EDIT.

package ent

import (
	"football_api/ent/profile"
	"football_api/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescCreatedAt is the schema descriptor for created_at field.
	profileDescCreatedAt := profileFields[3].Descriptor()
	// profile.DefaultCreatedAt holds the default value on creation for the created_at field.
	profile.DefaultCreatedAt = profileDescCreatedAt.Default.(func() time.Time)
	// profileDescUpdatedAt is the schema descriptor for updated_at field.
	profileDescUpdatedAt := profileFields[4].Descriptor()
	// profile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profile.DefaultUpdatedAt = profileDescUpdatedAt.Default.(func() time.Time)
	// profileDescID is the schema descriptor for id field.
	profileDescID := profileFields[0].Descriptor()
	// profile.DefaultID holds the default value on creation for the id field.
	profile.DefaultID = profileDescID.Default.(func() uuid.UUID)
}