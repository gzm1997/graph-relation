package model


type Predicate int

const (
	ShareFolder Predicate = iota // value --> 0
	WriteNum                 // value --> 1
	LastTime                 // value --> 2
	Name                     // value --> 3
	ShareFile
	CreateGroup
)


func (this Predicate) String() string {
	switch this {
	case ShareFolder:
		return "ShareFolder"
	case WriteNum:
		return "WriteNum"
	case LastTime:
		return "LastTime"
	case Name:
		return "UserName"
	case ShareFile:
		return "ShareFile"
	case CreateGroup:
		return "CreateGroup"
	default:
		return "Unknow"
	}
}