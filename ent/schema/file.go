package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// File holds file content and metadata.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			MinLen(1).
			Unique().
			Comment("The name of the file, used to reference it for downloads"),
		field.Time("CreationTime").
			Default(func() time.Time {
				return time.Now()
			}).
			Comment("The timestamp for when the File was created"),
		field.Time("LastModifiedTime").
			Comment("The timestamp for when the File was last modified"),
		field.Int("Size").
			Default(0).
			Min(0).
			Comment("The size of the file in bytes"),
		field.Bytes("Content").
			Comment("The content of the file"),
		field.String("Hash").
			Comment("A SHA3 digest of the content field"),
		field.String("ContentType").
			Comment("The content type of content"),
	}
}
