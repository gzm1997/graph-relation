package cayleyDb

import (
	//下面这个import必不可少 因为在这个import这个库的时候 它的init函数会注册这个类型的数据库 允许使用这个数据库进行存储
	_ "github.com/cayleygraph/cayley/graph/nosql/mongo"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"reflect"
	"relation-graph/graphRelation/graphOperation_v2/model"
	"github.com/astaxie/beego/config"
	"fmt"
	"errors"
)



var (
	dbUrl string
	db string
	//store *graph.Handle
)

func init()  {
	conf, err := config.NewConfig("ini", "./conf/graph.yaml")
	if err != nil {
		dbUrl = "mongodb://10.13.84.8:27017"
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




//根据id查找names
func FindNameById(store *graph.Handle, id int) (string, error) {
	//fmt.Println("to find name", id)
	p := cayley.StartPath(store, quad.Int(id)).Out(quad.String(model.Name.String()))
	var name string
	err := p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//fmt.Println("nav", nativeValue)
		if reflect.TypeOf(nativeValue).Kind() == reflect.String {
			name = nativeValue.(string)
			//fmt.Println("this is name", name)
		}
	})
	//fmt.Println("return name", name)
	return name, err
}


// 强关系1 根据主语谓语 查询宾语跟属性

func FindBySubjectAndPredicate(subject int, predicate model.Predicate) ([]map[string]interface{}, error) {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	p := cayley.StartPath(store, quad.Int(subject)).Out(predicate.String())
	var objects []int
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			objects = append(objects, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, errors.New("cayley query fail")
	}
	if len(objects) == 0 {
		return nil, nil
	}
	var r []map[string]interface{}
	for _, o := range objects {
		p = cayley.StartPath(store, quad.Int(o)).LabelContext(predicate.String() + "_info").In(quad.Int(subject))
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			var tMap = make(map[string]interface{})
			tMap["para"] = nativeValue
			tMap["name"], _ = FindNameById(store, o)
			tMap["id"] = o
			r = append(r, tMap)
			//if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			//	objects = append(objects, nativeValue.(int))
			//}
		})
	}
	return r, nil
}



// 强关系2 根据谓语宾语 查询主语和属性
// 返回值为一个map[string]interface{}数组
// 每个元素具体为
// {
//	 "name": "myname",
//	 "id": 123,
//	 "para": xxx
// }
func FindByPredicateAndObject(predicate model.Predicate, object int) ([]map[string]interface{}, error) {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	p := cayley.StartPath(store, quad.Int(object)).In(predicate.String())
	var subject []int
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			subject = append(subject, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, errors.New("cayley query fail")
	}
	if len(subject) == 0 {
		return nil, nil
	}
	var r []map[string]interface{}
	for _, s := range subject {
		p = cayley.StartPath(store, quad.Int(object)).LabelContext(predicate.String() + "_info").In(quad.Int(s))
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			var tMap = make(map[string]interface{})
			tMap["para"] = nativeValue
			tMap["name"], _ = FindNameById(store, s)
			tMap["id"] = s
			r = append(r, tMap)
			//if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			//	objects = append(objects, nativeValue.(int))
			//}
		})
	}
	return r, nil
}






// 弱关系 根据关系属性和宾语 查询可能认识的人
// 返回值为一个map[string]interface{}数组
// 每个元素细节为
// {
//	 "name": "myname",
//	 "id": 123
// }
func FindWeakRelationshipOnShareFile(object int, fileId int) ([]map[string]interface{}, error) {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	//var p *path.Path
	//var paraNode quad.Value
	//if reflect.TypeOf(para).Kind() == reflect.String {
	//	paraNode = quad.String(para.(string))
	//	//p = cayley.StartPath(store, quad.String(para.(string))).Labels()
	//} else if reflect.TypeOf(para).Kind() == reflect.Int {
	//	paraNode = quad.Int(para.(int))
	//	//p = cayley.StartPath(store, quad.Int(para.(int))).Labels()
	//} else  {
	//	return nil, errors.New("para type unknown")
	//}
	//p = cayley.StartPath(store, paraNode).Labels()
	//var label string
	//err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
	//	nativeValue := quad.NativeOf(value)
	//	if reflect.TypeOf(nativeValue).Kind() == reflect.String {
	//		label = nativeValue.(string)
	//	}
	//})
	//if err != nil {
	//	return nil, errors.New("cayley query fail")
	//}
	//l := len(label)
	//predicate := string([]rune(label)[:l - 5])
	var subject []int
	p := cayley.StartPath(store, quad.Int(object)).In(quad.String(model.ShareFile.String()))
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			subject = append(subject, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, errors.New("cayley query fail")
	}
	//fmt.Println("subjects", subject)
	var r []map[string]interface{}
	for _, s := range subject {
		p = cayley.StartPath(store, quad.Int(fileId)).LabelContext(model.ShareFile.String() + "_info").Out(quad.Int(s))
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			//fmt.Println("get id", nativeValue.(int))
			if nativeValue.(int) != object {
				var tMap = make(map[string]interface{})
				tMap["name"], _ = FindNameById(store, nativeValue.(int))
				tMap["id"] = nativeValue.(int)
				r = append(r, tMap)
				//if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
				//	objects = append(objects, nativeValue.(int))
				//}
			}
		})
	}
	return r, nil
}




