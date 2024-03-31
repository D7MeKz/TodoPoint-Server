// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"todopoint/common/db/ent/member"
	"todopoint/common/db/ent/point"
	"todopoint/common/db/ent/pointinfo"
	"todopoint/common/db/ent/schema"
	"todopoint/common/db/ent/task"
)

// The init function reads all schema descriptors with runtime codes
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescPassword is the schema descriptor for password field.
	memberDescPassword := memberFields[3].Descriptor()
	// member.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	member.PasswordValidator = memberDescPassword.Validators[0].(func(string) error)
	// memberDescCreatedAt is the schema descriptor for created_at field.
	memberDescCreatedAt := memberFields[4].Descriptor()
	// member.DefaultCreatedAt holds the default value on creation for the created_at field.
	member.DefaultCreatedAt = memberDescCreatedAt.Default.(time.Time)
	pointFields := schema.Point{}.Fields()
	_ = pointFields
	// pointDescCreatedAt is the schema descriptor for created_at field.
	pointDescCreatedAt := pointFields[3].Descriptor()
	// point.DefaultCreatedAt holds the default value on creation for the created_at field.
	point.DefaultCreatedAt = pointDescCreatedAt.Default.(time.Time)
	pointinfoFields := schema.PointInfo{}.Fields()
	_ = pointinfoFields
	// pointinfoDescModifiedAt is the schema descriptor for modified_at field.
	pointinfoDescModifiedAt := pointinfoFields[2].Descriptor()
	// pointinfo.DefaultModifiedAt holds the default value on creation for the modified_at field.
	pointinfo.DefaultModifiedAt = pointinfoDescModifiedAt.Default.(time.Time)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[3].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(time.Time)
	// taskDescModifiedAt is the schema descriptor for modified_at field.
	taskDescModifiedAt := taskFields[4].Descriptor()
	// task.DefaultModifiedAt holds the default value on creation for the modified_at field.
	task.DefaultModifiedAt = taskDescModifiedAt.Default.(time.Time)
}
