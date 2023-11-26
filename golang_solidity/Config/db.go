package Config

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct { //定义用户结构体
	id       int    //id
	name     string //用户名     与数据库变量名一致
	passWord string //密码
}

//数据库连接
func init() {
	var err error
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/01") // 设置连接数据库的参数
	if err != nil {
		panic(err)
	}
	err = db.Ping() //连接数据库
	if err != nil {
		fmt.Println("数据库连接失败")
	}
}

// UserLogin /*登入账户*/
func UserLogin(username string, password string) int {
	rows, err := db.Query("SELECT passWord FROM a where name=?", username) //通过账户名查询密码
	if err != nil {
		fmt.Println("查询失败")
	}

	var a string //查询到的密码
	//遍历返回结果
	for rows.Next() {
		err = rows.Scan(&a)
		if err != nil {
			fmt.Println(err)
		}
	}
	var b int

	if a == password { //判断数据库密码 与前端输入的密码是否相同
		b = 1
		
	} else {
		b = 0
	
	}

	
	
	// fmt.Println(a)
	return b
}

// UserRegister 用户注册
func UserRegister(c *gin.Context, username string, password string) int {

	var a int
	rows, err := db.Query("SELECT  *  from a where name=?",username) //查询superoot内所有信息
	fmt.Println(err)
	if err != nil {
		fmt.Println("数据库无此用户，可添加") //处理错误机制
	
	}
	

	var u User                            //已知User为结构体
	err = rows.Scan(&u.name, &u.passWord) //扫描name,passWord数据
	if err != nil {
		fmt.Println(err) //处理错误机制
	}
	fmt.Println(u)
	if username != u.name { 
		//判断数据库内的u.name是否与前端页面传入的username相匹配 若不匹配 则无此用户

		results, err := db.Exec("insert into a(name,passWord) values(?,?)", username, password) //执行插入sql语句
		if err != nil {
			
			fmt.Println("执行失败") //处理错误机制
			a=0
			return a
		} else {
			rows, _ := results.RowsAffected() //输出执行的行数
			if rows != 1 {
				c.JSON(200, gin.H{
					"success": false,
				})
				a=0
				return a
			} else {
				c.JSON(200, gin.H{
					"success":  true,
					"username": username,
					"password": password,
				})
				a=1
				return a
			}
		}
	}

a=0
return a

}



func UserGet( username string) (string,string) {

	
	rows, err := db.Query("SELECT  *  from a where name=?",username) //查询superoot内所有信息
	fmt.Println(err)
	if err != nil {
		fmt.Println("数据库无此用户，可添加") //处理错误机制
	
	}
	
        var u User
		//遍历返回结果
		for rows.Next() {
			err = rows.Scan(&u.id,&u.name,&u.passWord)
			if err != nil {
				// fmt.Println(err)
			}
		}
		// fmt.Println(u.name)


return u.name,u.passWord

}
func UserChange( username string,password string) (string,string) {

	
	results, err := db.Query("update  a set name=?,passWord=? where name=戴加涛123",username,password) //查询superoot内所有信息
	  
		if err != nil {
			
			fmt.Println("执行失败") //处理错误机制
			
			
		} 
		fmt.Println(results)
			
		
		
	
return username,password

}

//定义书籍结构体
type books struct { // 结果集，参数名需大写
	Bookid    int
	BookName  string
	CreatTime string
}

