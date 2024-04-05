package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

func (Room) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "chat_rooms"},
	}
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return nil
}
