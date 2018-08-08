package model

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley"
)


type User struct {
	Name string `json:"UserName"`
	Id   int `json:"UserId"`
}

func (this User) Quad() quad.Quad {
	return quad.Make(quad.Int(this.Id), quad.String(Name.String()), quad.String(this.Name), quad.String("Name_info"))
}

func AddUser(user User) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuad(user.Quad())
}

func AddUsers(users []User) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	var quads []quad.Quad
	for _, u := range users {
		quads = append(quads, u.Quad())
	}
	return store.AddQuadSet(quads)
}


