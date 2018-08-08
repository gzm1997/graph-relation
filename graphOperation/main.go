package main

import (
	"relation-graph/graphRelation/graphOperation/model"
	"fmt"
	"github.com/cayleygraph/cayley/graph/iterator"
	"relation-graph/graphRelation/graphOperation/cayleyDb"
	"github.com/cayleygraph/cayley/quad"
)




func main()  {
	u1 := model.User{"tom", 15331094}
	u2 := model.User{"amy", 15331082}
	u3 := model.User{"kevin", 15331080}
	r1 := model.Relation{u1, model.ShareFolder, u2, 3}
	r2 := model.Relation{u2, model.WriteNum, u3, 2}
	r3 := model.Relation{u1, model.ShareFolder, u3, 4}
	fmt.Println(cayleyDb.InsertOneRelation(r1))
	fmt.Println(cayleyDb.InsertOneRelation(r2))
	fmt.Println(cayleyDb.InsertOneRelation(r3))

	fmt.Println(cayleyDb.InsertManyRelation([]model.Relation{r1, r2, r3}))


	result1, _ := cayleyDb.Find_All_By_SubjectId_Predicate(u1.Id, model.ShareFolder)
	model.ShowResult(result1)


	result4, _ := cayleyDb.Find_Range_By_SubjectId_Predicate(u1.Id, model.ShareFolder, iterator.CompareLTE, 3)
	model.ShowResult(result4)

	result3, _ := cayleyDb.Find_All_By_ObjectId_Predicate(u3.Id, model.WriteNum)
	model.ShowResult(result3)

	result2, _ := cayleyDb.Find_Range_ByObjectId_Predicate(u3.Id, model.ShareFolder, iterator.CompareGTE, 4)
	model.ShowResult(result2)

    q := quad.Make(u1.Id, model.ShareFolder.String(), u2.Id, 3)
	fmt.Println("update label", cayleyDb.UpdateLabel(q, 12))

	result5, _ := cayleyDb.Find_All_By_SubjectId_Predicate(u1.Id, model.ShareFolder)
	model.ShowResult(result5)


	fmt.Println("update name", cayleyDb.UpdateName(u1.Id, "Tom new names"))

	result6, _ := cayleyDb.Find_All_By_ObjectId_Predicate(u2.Id, model.ShareFolder)
	model.ShowResult(result6)

	fmt.Println("update subject", cayleyDb.UpdateSubject(quad.Make(u1.Id, q.Predicate, q.Object, 12), u3))
	result7, _ := cayleyDb.Find_All_By_ObjectId_Predicate(u2.Id, model.ShareFolder)
	model.ShowResult(result7)


	result8, _ := cayleyDb.Find_All_By_ObjectId_Predicate(u3.Id, model.ShareFolder)
	model.ShowResult(result8)

	nu := model.User{"lxm", 15331067}
	fmt.Println(cayleyDb.UpdateObject(quad.Make(u1.Id, model.ShareFolder, u3.Id, 4), nu))

	r, _ := cayleyDb.Find_All_By_SubjectId_Predicate(u1.Id, model.ShareFolder)
	model.ShowResult(r)


}
