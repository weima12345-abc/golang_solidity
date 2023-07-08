package Config
type charityFunding struct {
	initiator string
	title string 
	info int64  
	goal  int64 
	 success bool                   // 捐赠是否成功，成功则 amount 含义为项目剩余的钱
	 endTime uint64
     amount uint                     // 当前已经筹集到的金额
     numFunders uint                 // 投资记录数量
     numUses uint                 // 使用请求数量
   

}

