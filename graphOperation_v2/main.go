package main

import (
	"relation-graph/graphRelation/graphOperation_v2/model"
	"fmt"
	"relation-graph/graphRelation/graphOperation_v2/cayleyDb"
)

func main()  {
	// 添加三个用户
	u1 := model.User{"tom", 1}
	u2 := model.User{"amy", 2}
	u3 := model.User{"lorry", 3}
	u4 := model.User{"kevin", 4}
	u5 := model.User{"cousin", 5}
	u6 := model.User{"harry", 6}
	users := []model.User{u1, u2, u3, u4, u5, u6}
	if model.AddUsers(users) == nil {
		fmt.Println("insert user data successfully\n")
	}

	// 添加5条关系
	r1 := model.QuadRelation{u1.Id, model.WriteNum, u4.Id, 1}
	r2 := model.QuadRelation{u1.Id, model.ShareFile, u2.Id, 123456}
	r3 := model.QuadRelation{u1.Id, model.ShareFolder, u3.Id, 3}
	r4 := model.QuadRelation{u5.Id, model.ShareFile, u2.Id, 123456}
	r5 := model.QuadRelation{u5.Id, model.WriteNum, u2.Id, 6}
	r6 := model.QuadRelation{u6.Id, model.ShareFile, u2.Id, 654321}
	r7 := model.QuadRelation{u1.Id, model.ShareFile, u5.Id, 123456}
	relations := []model.QuadRelation{r1, r2, r3, r4, r5, r6, r7}
	if model.AddQuadRlations(relations) == nil {
		fmt.Println("insert relation successfully\n")
	}
	// 查询u1 分享了文件夹给谁 数量是多少
	result, err := cayleyDb.FindBySubjectAndPredicate(u1.Id, model.ShareFolder)
	if err == nil {
		fmt.Println("查询u1 分享了文件夹给谁 数量是多少: ")
		for _, r := range result {
			fmt.Println(r)
		}
		fmt.Println("")
	}


	// 查询谁分享了文件夹给u2 分享了多少个
	result, err = cayleyDb.FindByPredicateAndObject(model.ShareFile, u2.Id)
	if err == nil {
		fmt.Println("查询谁分享了文件给u2 分享了多少个: ")
		for _, r := range result {
			fmt.Println(r)
		}
		fmt.Println("")
	}

	// 查询谁跟u2在获得文件123456上存在弱关系
	result, err = cayleyDb.FindWeakRelationshipOnShareFile(u2.Id, 123456)
	if err == nil {
		fmt.Println("查询谁跟u2在获得文件123456上存在弱关系: ")
		for _, r := range result {
			fmt.Println(r)
		}
	}
}

