// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/dmalykh/tagservice/repository/entgo/ent/category"
	"github.com/dmalykh/tagservice/repository/entgo/ent/namespace"
	"github.com/dmalykh/tagservice/repository/entgo/ent/relation"
	"github.com/dmalykh/tagservice/repository/entgo/ent/schema"
	"github.com/dmalykh/tagservice/repository/entgo/ent/tag"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	namespaceFields := schema.Namespace{}.Fields()
	_ = namespaceFields
	// namespaceDescName is the schema descriptor for name field.
	namespaceDescName := namespaceFields[0].Descriptor()
	// namespace.NameValidator is a validator for the "name" field. It is called by the builders before save.
	namespace.NameValidator = namespaceDescName.Validators[0].(func(string) error)
	relationFields := schema.Relation{}.Fields()
	_ = relationFields
	// relationDescTagID is the schema descriptor for tag_id field.
	relationDescTagID := relationFields[0].Descriptor()
	// relation.TagIDValidator is a validator for the "tag_id" field. It is called by the builders before save.
	relation.TagIDValidator = relationDescTagID.Validators[0].(func(int) error)
	// relationDescEntityID is the schema descriptor for entity_id field.
	relationDescEntityID := relationFields[1].Descriptor()
	// relation.EntityIDValidator is a validator for the "entity_id" field. It is called by the builders before save.
	relation.EntityIDValidator = relationDescEntityID.Validators[0].(func(int) error)
	// relationDescNamespaceID is the schema descriptor for namespace_id field.
	relationDescNamespaceID := relationFields[2].Descriptor()
	// relation.NamespaceIDValidator is a validator for the "namespace_id" field. It is called by the builders before save.
	relation.NamespaceIDValidator = relationDescNamespaceID.Validators[0].(func(int) error)
	// relationDescCreatedAt is the schema descriptor for created_at field.
	relationDescCreatedAt := relationFields[3].Descriptor()
	// relation.DefaultCreatedAt holds the default value on creation for the created_at field.
	relation.DefaultCreatedAt = relationDescCreatedAt.Default.(func() time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[0].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescCategoryID is the schema descriptor for category_id field.
	tagDescCategoryID := tagFields[3].Descriptor()
	// tag.CategoryIDValidator is a validator for the "category_id" field. It is called by the builders before save.
	tag.CategoryIDValidator = tagDescCategoryID.Validators[0].(func(int) error)
}
