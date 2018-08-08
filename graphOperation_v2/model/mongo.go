package model

import (
	//下面这个import必不可少 因为在这个import这个库的时候 它的init函数会注册这个类型的数据库 允许使用这个数据库进行存储
	_ "github.com/cayleygraph/cayley/graph/nosql/mongo"
	"github.com/astaxie/beego/config"
	"fmt"
	"github.com/cayleygraph/cayley/graph"
)

var (
	dbUrl string
	db string
	//store *graph.Handle
)

func init()  {
	conf, err := config.NewConfig("ini", "./conf/graph.yaml")
	if err != nil {
		dbUrl = "mongodb://120.92.100.60:27017"
		db = "mongo"
		//panic(err)
	}
	dbUrl = fmt.Sprintf("mongodb://%s:%s/%s", conf.String("mongodbIp"), conf.String("mongodbPort"), conf.String("dbName"))
	db = conf.String("database")
	err = graph.InitQuadStore(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	//fmt.Println("lala")
}