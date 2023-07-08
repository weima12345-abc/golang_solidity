package Hander

import (
	mid "charity/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
	"github.com/ethereum/go-ethereum/common"
	
)
//获得当前账户余额,从合约
func CountAmount(c *gin.Context) {
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetContractAmount(contract)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

//获得所有的公益项目从合约
func GetallFundings(c *gin.Context) {
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetallFundings(contract)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}
//获得某个公益项目  中间件
func GetoneFunding(c *gin.Context) {
	id := c.PostForm("id")

	//string转int
i_id,err :=strconv.Atoi(id)

//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetnewFunding(contract,big1)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

//用户捐助  中间件
func Contribute(c *gin.Context) {
	id := c.PostForm("id")
    value :=c.PostForm("value")

   //string转int
   i_id, err := strconv.ParseInt(id, 10, 64)
   if err != nil {
	   respError(c, err)

	   return
   }
   //int先转uint64 再转*big.Int类型
   big1 := big.NewInt(i_id)
	
   big2, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		respError(c, err)
		return
	}

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.Contribute(client,contract,big1,big2)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}



//新增公益项目  中间件
func NewFunding(c *gin.Context) {
	
	initiator := c.PostForm("initiator")
	
	title := c.PostForm("title")
	info := c.PostForm("info")
	goal := c.PostForm("goal")
a,err :=strconv.ParseInt(goal, 10, 64)
	//string转int
g_goal := big.NewInt(a)
// i_initiator := String2Bytes(initiator)
// add_i_initiator :=BytesAddress(i_initiator)
add_i_initiator :=common.HexToAddress(initiator)
//int先转uint64 再转*big.Int类型
	// big1 := new(big.Int).SetUint64(uint64(g_goal))
	

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.NewFunding(client,contract,add_i_initiator,title,info,g_goal)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}


//删除公益项目  中间件
func DelFunding(c *gin.Context) {
	

	id := c.PostForm("id")
a,err :=strconv.ParseInt(id, 10, 64)

g_goal := big.NewInt(a)


	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.DelFunding(client,contract,g_goal)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

//收藏公益项目  中间件
func InsFunding(c *gin.Context) {
	

	id := c.PostForm("id")
a,err :=strconv.ParseInt(id, 10, 64)

g_goal := big.NewInt(a)


	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.InsFunding(client,contract,g_goal)
	fmt.Println(err)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

//获取已加入收藏的公益项目  中间件
func GetInsFunding(c *gin.Context) {
	
	initiator := c.PostForm("initiator")
	add_i_initiator :=common.HexToAddress(initiator)

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetallInsFunding(contract,add_i_initiator)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

//获取用户已捐助的所有公益项目  中间件
func GetallContribute(c *gin.Context) {
	initiator := c.PostForm("initiator")
	add_i_initiator :=common.HexToAddress(initiator)


	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetallContribute(contract,add_i_initiator)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}


//删除用户收藏的某个公益项目  中间件
func DelInscollect(c *gin.Context) {
	

	id := c.PostForm("id")
a,err :=strconv.ParseInt(id, 10, 64)

g_goal := big.NewInt(a)


	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.DelInscollect(client,contract,g_goal)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

// //获得用户在某个项目上捐助的金额  中间件
// func GetOneFamout(c *gin.Context) {
// 	initiator := c.PostForm("initiator")
// 	add_i_initiator :=common.HexToAddress(initiator)


// 	id := c.PostForm("id")
// a,err :=strconv.ParseInt(id, 10, 64)

// g_goal := big.NewInt(a)


// 	client, err := mid.GetClient()
// 	if err != nil {
// 		respError(c, err)
// 		fmt.Println(err)
// 		return
// 	}
// 	contract, err := mid.GetCallFaceTx(*client)
// 	if err != nil {
// 		respError(c, err)
// 		fmt.Println(err)
// 		return
// 	}
// 	res, err := mid.DetOneFamout(contract,add_i_initiator,g_goal)
// 	if err != nil {
// 		respError(c, err)
// 		fmt.Println(err)
// 		return
// 	}
// 	respOK(c, res)
// 	client.Close()
// }











