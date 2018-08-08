package model

import (
	"github.com/cayleygraph/cayley/quad"
	"encoding/json"
	"fmt"
)



type State int

const (
	ShareFolder State = iota // value --> 0
	WriteNum                 // value --> 1
	LastTime                 // value --> 2
	Name                     // value --> 3
)

func (this State) String() string {
	switch this {
	case ShareFolder:
		return "ShareFolder"
	case WriteNum:
		return "WriteNum"
	case LastTime:
		return "LastTime"
	case Name:
		return "UserName"
	default:
		return "Unknow"
	}
}

type User struct {
	Name string `json:"UserName"`
	Id   int `json:"UserId"`
}

type Relation struct {
	Subject   User
	Predicate State
	Object    User
	Label     interface{}
}

func (this Relation) Quads() []quad.Quad {
	_quad := quad.Make(this.Subject.Id, this.Predicate.String(), this.Object.Id, this.Label)
	subjectQuad := quad.Make(this.Subject.Id, Name.String(), this.Subject.Name, nil)
	objectQuad := quad.Make(this.Object.Id, Name.String(), this.Object.Name, nil)
	return []quad.Quad{_quad, subjectQuad, objectQuad}
}

func (this Relation) GetRelationQuad() quad.Quad {
	return quad.Make(this.Subject.Id, this.Predicate.String(), this.Object.Id, this.Label)
}

func (this Relation) ConstructByJsonByte(b []byte) error {
	err := json.Unmarshal(b, &this)
	return err
}

type SearchResult struct {
	User `json:"User"`
	Label interface{} `json:"Label,omitempty"`
}


func (this SearchResult) Json() ([]byte, error) {
	b, err := json.Marshal(this)
	return b, err
}

func ShowResult(rl []SearchResult)  {
	for _, r := range rl {
		j, _ := r.Json()
		fmt.Printf("%s\n", j)
	}
	fmt.Println("")
}