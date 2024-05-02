package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Chunk holds the schema definition for the Chunk entity.
type Chunk struct {
	ent.Schema
}

// Fields of the Chunk.
func (Chunk) Fields() []ent.Field {
	return []ent.Field{
		field.Time("t").Immutable(),
		field.String("name"),
		field.Bytes("data").Default([]byte{}),
	}
}

// Edges of the Chunk.
func (Chunk) Edges() []ent.Edge {
	return nil
}
