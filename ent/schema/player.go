package schema

import (
	"database/sql/driver"
	"encoding/json"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
		field.Other("equip", &Equip{}).
			SchemaType(map[string]string{
				"mysql": "text",
			}).
			Optional().
			Nillable(),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("players").
			Unique(),
	}
}

type Equip struct {
	Id    int64
	Attr  string
	Attrs map[int]int64
}

func (e *Equip) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), e)
}
func (e Equip) Value() (driver.Value, error) {
	return json.Marshal(e)
}
