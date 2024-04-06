package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "chat_rooms_members"},
	}
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("RoomID"),
		field.Uint32("UserID"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
