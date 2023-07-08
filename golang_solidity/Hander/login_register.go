package Hander

import ("charity/Config"
"github.com/gin-gonic/gin"
"github.com/zhangdapeng520/zdpgo_password"
"fmt"
)
type User struct { //定义用户结构体
	id       int    //id
	name     string //用户名     与数据库变量名一致
	passWord string //密码
}
func Register(c *gin.Context) {
	name := c.PostForm("username")         //将前端通过ajax传入的username的值传给name
	password := c.PostForm("password") //将前端通过ajax传入的password的值传给password
	p := zdpgo_password.New(zdpgo_password.PasswordConfig{})

	data := password
	fmt.Println(data)

	var (
		result string
		err    error
	)

// md5加密
result, err = p.Hash.Md5.EncryptString(data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(result)
	res :=Config.UserRegister(c, name, result)     //将name,password 传给UserRegister()方法
	tohtml(c, res)                           //返回给html
}


//登录

func Login(c *gin.Context) {
	name := c.PostForm("username")         //将前端通过ajax传入的username的值传给name
	password := c.PostForm("password")     //将前端通过ajax传入的password的值传给password
	p := zdpgo_password.New(zdpgo_password.PasswordConfig{})

	data := password
	// fmt.Println(data)

	var (
		result string
		err    error
	)
// md5加密
result, err = p.Hash.Md5.EncryptString(data)
if err != nil {
	// fmt.Println(err)
}
// fmt.Println(result)
	r:= Config.UserLogin(name, result) //将name ,password传给UserLogin()方法
	tohtml(c,r) 
	                    //返回 给html

}

func GetUser(c *gin.Context){
	name := c.PostForm("username")         //将前端通过ajax传入的username的值传给name
    // passWord :=c.PostForm("password")

	// var (
	// 	result string
	
	// )

	r,s:= Config.UserGet(name) //将name ,password传给UserLogin()方法
	tohtml1(c,r,s) 
	                    //返回 给html

}

func ChangeUser(c *gin.Context){
	name := c.PostForm("username")         //将前端通过ajax传入的username的值传给name
    passWord :=c.PostForm("password")

	// var (
	// 	result string
	
	// )

	r,s:= Config.UserChange(name,passWord) //将name ,password传给UserLogin()方法
	tohtml1(c,r,s) 
	                    //返回 给html

}


