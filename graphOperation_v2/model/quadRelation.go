package model

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
)

type QuadRelation struct {
	Subject int `json:"Subject"`
	Predicate Predicate `json:"Predicate"`
	Object int `json:"Object"`
	Para interface{} `json:"Para"`
}

func (this QuadRelation) Quads() []quad.Quad {
	var quads []quad.Quad
	r := Relation{this.Subject, this.Predicate, this.Object}
	info := Information{this.Para, this.Subject, this.Object}
	quads = append(quads, r.Quad())
	quads = append(quads, info.Quad(this.Predicate.String()))
	return quads
}


func AddQuadRelation(relation QuadRelation) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(relation.Quads())
}

func AddQuadRlations(relations []QuadRelation) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	var quads []quad.Quad
	for _, r := range relations {
		s := r.Quads()
		for _, i := range s {
			quads = append(quads, i)
		}
	}
	return store.AddQuadSet(quads)
}
