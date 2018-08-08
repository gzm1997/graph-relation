package models

import (
	"relation-graph/graphRelation/createTriple/modelBase"
	"relation-graph/graphRelation/createTriple/modelRelation"
)

type File = modelBase.File
type User = modelBase.User
type Group = modelBase.Group
type CreateFileLink = modelRelation.CreateFileLink
type ClickFileLink = modelRelation.ClickFileLink
type CreateGroupShareLink = modelRelation.CreateGroupShareLink
type ClickGroupShareLink = modelRelation.ClickGroupShareLink

type Fileinfo struct {
	Fileid int `json:"fileid"`
	Groupid int `json:"groupid"`
	GroupType string `json:"group_type"`
	ParentId int `json:"parentid"`
	LastUpdateUserid int `json:"last_update_userid"`
	Fname string `json:"fname"`
	Ftype string `json:"ftype"`
	Fver int `json:"fver"`
	Fsize int `json:"fsize"`
	Store int `json:"store"`
	Fsha string `json:"fsha"`
	Ctime int `json:"ctime"`
	Mtime int `json:"mtime"`
}

type Creator struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Corpid int `json:"corpid"`
}

type Groupinfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	GroupType string `json:"group_type"`
	Creator int `json:"creator"`
	Price int `json:"price"`
	Ctime int `json:"ctime"`
	Mtime int `json:"mtime"`
	Utime int `json:"utime"`
	Atime int `json:"atime"`
	Status int `json:"status"`
}

type Link struct {
	Sid string `json:"sid"`
	FileId int `json:"fileid"`
	UserId int `json:"userid"`
	Ctime int `json:"ctime"`
	Chkcode string `json:"chkcode"`
	Clicked int `json:"clicked"`
	GroupId int `json:"groupid"`
	Status string `json:"status"`
	Ranges string `json:"ranges"`
	Permission string `json:"permission"`
	ExpirePeriod int `json:"expire_period"`
	ExpireTime int `json:"expire_time"`
	Creator Creator `json:"creator"`
}

type CreateFileLinkRaw struct {
	Topic string `json:"topic"`
	Type string `json:"type"`
	Operator int `json:"operator"`
	Time int `json:"time"`
	Fileinfo Fileinfo `json:"fileinfo"`
	Creator Creator `json:"creator"`
	Groupinfo Groupinfo `json:"groupinfo"`
	Link Link `json:"link"`
	LinkUrl string  `json:"link_url"`
}

func (this CreateFileLinkRaw) ConvertRawDataToCreateFileLink() CreateFileLink {
	var permission modelRelation.FileLinkPermission
	switch this.Link.Permission {
	case "write": permission = modelRelation.Write
	case "read": permission = modelRelation.Read
	default:
		permission = -1
	}
	cfl := CreateFileLink{User{this.Creator.Id, this.Creator.Name}, File{this.Fileinfo.Fileid, this.Fileinfo.Fname, this.Fileinfo.Groupid}, permission, this.Link.Ctime}
	return cfl
}

