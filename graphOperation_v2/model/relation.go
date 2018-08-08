package model

import (
	"github.com/cayleygraph/cayley/quad"
)

type Relation struct {
	Subject int `json:"Subject"`
	Predicate Predicate `json:"Predicate"`
	Object int `json:"Object"`
}

func (this Relation) Quad() quad.Quad {
	return quad.Make(quad.Int(this.Subject), quad.String(this.Predicate.String()), quad.Int(this.Object), quad.String("Relation"))
}




type Information struct {
	Subject interface{} `json:"Subject"`
	Predicate int `json:"Predicate"`
	Object int `json:"Object"`
}

func (this Information) Quad(subgraph string) quad.Quad {
	return quad.Make(this.Subject, quad.Int(this.Predicate), quad.Int(this.Object), quad.String(subgraph + "_info"))
}
