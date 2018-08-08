// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"relation-graph/graphRelation/graphServer/controllers"
)

func init() {
	//文件相关的路由
	fileNameSpace := beego.NewNamespace("/file",
		//创建文件分享链接
		beego.NSRouter("/createfilelink", &controllers.FileController{}, "post:CreateFileLink"),
		//点击文件分享链接
		beego.NSRouter("/clickfilelink", &controllers.FileController{}, "post:ClickFileLink"))

	//群邀请相关的路由
	groupNameSpace := beego.NewNamespace("/group",
		//创建群邀请链接
		beego.NSRouter("/creategroupsharelink", &controllers.GroupController{}, "post:CreateGroupShareLink"),
		//点击群邀请链接
		beego.NSRouter("/clickgroupsharelink", &controllers.GroupController{}, "post:ClickGroupShareLink"))


	//将命名空间中的所有路由函数注册
	beego.AddNamespace(fileNameSpace, groupNameSpace)
}
