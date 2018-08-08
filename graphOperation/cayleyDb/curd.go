package cayleyDb


import (
	//下面这个import必不可少 因为在这个import这个库的时候 它的init函数会注册这个类型的数据库 允许使用这个数据库进行存储
	_ "github.com/cayleygraph/cayley/graph/nosql/mongo"
	"github.com/cayleygraph/cayley/graph"
	"github.com/astaxie/beego/config"
	"fmt"
	"github.com/cayleygraph/cayley"
	"relation-graph/graphRelation/graphOperation/model"
	"github.com/cayleygraph/cayley/quad"
	"reflect"
	"errors"
	"github.com/cayleygraph/cayley/graph/iterator"
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

func InsertOneRelation(_relation model.Relation) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	_quadSet := _relation.Quads()
	return store.AddQuadSet(_quadSet)
}

func InsertManyRelation(_relations []model.Relation) error  {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	var _quadSet []quad.Quad
	for _, r := range _relations {
		newQuad := r.Quads()
		for _, q := range newQuad {
			_quadSet = append(_quadSet, q)
		}
	}
	return store.AddQuadSet(_quadSet)
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


//根据给定的主语谓语找到所有的 宾语-标签键值对
func Find_All_By_SubjectId_Predicate(id int, state model.State) ([]model.SearchResult, error) {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	p := cayley.StartPath(store, quad.Int(id)).Labels()
	labels := []int{}
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			labels = append(labels, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, errors.New("cayley query fail")
	}
	if len(labels) == 0 {
		return nil, nil
	}
	results := []model.SearchResult{}
	for _, l := range labels {
		p = cayley.StartPath(store, quad.Int(id)).LabelContext(quad.Int(l)).Out(state.String())
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
				var name string
				name, err = FindNameById(store, nativeValue.(int))
				if err != nil {
					panic(err)
				}
				u := model.User{name, nativeValue.(int)}
				result := model.SearchResult{u, l}
				results	= append(results, result)
			}
		})
		if err != nil {
			return nil, errors.New("cayley query fail")
		}
	}
	return results, nil
}



// 根据主语谓语和给定的的最小值标签 找到跟最小值标签符合comparison关系的多个 宾语-标签键值对
// comparison是一个Operator类型 本质上是一个int类型
// 0 <
// 1 <=
// 2 >
// 3 >=
func Find_Range_By_SubjectId_Predicate(id int, state model.State, comparison iterator.Operator, threshold int) ([]model.SearchResult, error) {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	//r := make(map[int]int)
	p := cayley.StartPath(store, quad.Int(id)).Labels().Filter(comparison, quad.Int(threshold))
	labels := []int{}
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			labels = append(labels, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, errors.New("cayley query fail")
	}
	if len(labels) == 0 {
		return nil, nil
	}
	results := []model.SearchResult{}
	for _, l := range labels {
		p = cayley.StartPath(store, quad.Int(id)).LabelContext(quad.Int(l)).Out(state.String())
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
				var name string
				name, err = FindNameById(store, nativeValue.(int))
				if err != nil {
					panic(err)
				}
				u := model.User{name, nativeValue.(int)}
				result := model.SearchResult{u, l}
				results	= append(results, result)
			}
		})
		if err != nil {
			return nil, errors.New("cayley query fail")
		}
	}
	return results, nil
}



//根据给定的宾语谓语找到所有的 主语-标签键值对
func Find_All_By_ObjectId_Predicate(id int, state model.State) ([]model.SearchResult, error) {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	p := cayley.StartPath(store, quad.Int(id)).Labels()
	labels := []int{}
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			labels = append(labels, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, errors.New("cayley query fail")
	}
	if len(labels) == 0 {
		return nil, nil
	}
	results := []model.SearchResult{}
	for _, l := range labels {
		//fmt.Println("l", l)
		p = cayley.StartPath(store, quad.Int(id)).LabelContext(quad.Int(l)).In(state.String())
		//fmt.Println("P", name, l, )
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
				var name string
				name, err = FindNameById(store, nativeValue.(int))
				if err != nil {
					panic(err)
				}
				u := model.User{name, nativeValue.(int)}
				result := model.SearchResult{u, l}
				results	= append(results, result)
			}
		})
		if err != nil {
			return nil, errors.New("cayley query fail")
		}
	}
	return results, nil
}



// 根据宾语谓语和给定的的最小值标签 找到跟最小值标签符合comparison关系的多个 主语-标签键值对
// comparison是一个Operator类型 本质上是一个int类型
// 0 <
// 1 <=
// 2 >
// 3 >=
func Find_Range_ByObjectId_Predicate(id int, state model.State, comparison iterator.Operator, threshold int) ([]model.SearchResult, error) {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	//r := make(map[int]int)
	p := cayley.StartPath(store, quad.Int(id)).Labels().Filter(comparison, quad.Int(threshold))
	labels := []int{}
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
			labels = append(labels, nativeValue.(int))
		}
	})
	if err != nil {
		return nil, errors.New("cayley query fail")
	}
	if len(labels) == 0 {
		return nil, nil
	}
	results := []model.SearchResult{}
	for _, l := range labels {
		p = cayley.StartPath(store, quad.Int(id)).LabelContext(quad.Int(l)).In(state.String())
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			nativeValue := quad.NativeOf(value)
			if reflect.TypeOf(nativeValue).Kind() == reflect.Int {
				var name string
				name, err = FindNameById(store, nativeValue.(int))
				if err != nil {
					panic(err)
				}
				u := model.User{name, nativeValue.(int)}
				result := model.SearchResult{u, l}
				results	= append(results, result)
			}
		})
		if err != nil {
			return nil, errors.New("cayley query fail")
		}
	}
	return results, nil
}

//更新标签
func UpdateLabel(q quad.Quad, new_label interface{}) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	err = store.RemoveQuad(q)
	if err != nil {
		return errors.New("remove quad from cayley fail")
	}
	new_quad := quad.Make(q.Subject, q.Predicate, q.Object, new_label)
	return store.AddQuad(new_quad)
}

//更新名字
func UpdateName(id int, name string) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	old_name, e := FindNameById(store, id)
	if e != nil {
		return errors.New("no such user")
	}
	new_quad := quad.Make(id, model.Name.String(), name, nil)
	old_quad := quad.Make(id, model.Name.String(), old_name, nil)
	err = store.RemoveQuad(old_quad)
	if err != nil {
		return errors.New("remove old quad error")
	}
	return store.AddQuad(new_quad)
}


//更新主语
func UpdateSubject(q quad.Quad, new_user model.User) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	//fmt.Println("s", q.Subject, "p", q.Predicate, "o", q.Object, "l", q.Label)
	//删除原有的四元组
	err = store.RemoveQuad(q)
	if err != nil {
		return err
	}
	//为了保险 从新添加一个这个主语的Name声明四元组
	err = store.AddQuad(quad.Make(new_user.Id, model.Name.String(), new_user.Name, nil))
	if err != nil {
		return errors.New("add quad to cayley fail")
	}
	//添加新的四元组
	err = store.AddQuad(quad.Make(new_user.Id, q.Predicate, q.Object, q.Label))
	if err != nil {
		return errors.New("add quad to cayley fail")
	}
	return nil
}

//更新宾语
func UpdateObject(q quad.Quad, new_user model.User) error {
	store, err := cayley.NewGraph(db, dbUrl, nil)
	if err != nil {
		panic(err)
	}
	//删除原有的四元组
	err = store.RemoveQuad(q)
	if err != nil {
		return errors.New("remove quad from cayley fail")
	}
	//为了保险 从新添加一个这个主语的Name声明四元组
	err = store.AddQuad(quad.Make(new_user.Id, model.Name.String(), new_user.Name, nil))
	if err != nil {
		return errors.New("add quad to cayley fail")
	}
	//添加新的四元组
	err = store.AddQuad(quad.Make(q.Subject, q.Predicate, new_user.Id, q.Label))
	if err != nil {
		return errors.New("add quad to cayley fail")
	}
	return nil
}