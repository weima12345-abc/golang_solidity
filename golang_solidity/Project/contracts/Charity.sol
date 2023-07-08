// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Charity{

address public host;  //管理员

constructor(){
  host=msg.sender;
}

modifier onlyhost(){
      require(msg.sender==host);
      _;
}
  // 捐赠人结构体
  struct Funder {
    
    address payable addr;   // 捐赠人地址
    uint amount;            // 捐赠数额
  }

  // 资金使用请求结构体
  struct Use {
    string info;                     // 使用请求的说明
    uint goal;                       // 使用请求的数额
    uint agreeAmount;                // 目前的同意数额
    uint disagree;                   // 目前的不同意数额
    bool over;                       // 请求是否结束
    mapping(uint => uint) agree;     // 出资人是否同意 0: 还没决定，1：同意，2：不同意
  }


  // 捐募项目的结构体
  struct Funding {
    address  initiator;       // 求助人
    string title;                    // 项目标题
    string info;                     // 项目简介
    uint goal;                       // 目标金额
    // uint endTime;                    // 捐赠结束时间
bool exist;  //判断这个项目是否存在
uint id; //第几个项目
    bool success;                    // 捐赠是否成功，成功则 amount 含义为项目剩余的钱

    uint amount;                     // 当前已经筹集到的金额
    uint numFunders;                 // 投资记录数量
    uint numUses;                    // 使用请求数量
   
    // mapping(uint => Use) uses;       // 所有的使用请求
  }





 mapping(uint => Funder) public funders; // 投资记录具体信息收藏收藏
 mapping(address=>mapping(uint=> Funder)) public msgFunder;
uint public numContributeFundings;  //捐赠项目数量
  uint public numFundings;                  // 发布公益项目数量
   uint  public numCollect;                  // 收藏项目，数量
  mapping(uint => Funding) public fundings; // 所有的捐赠项目
  mapping(uint => Funding[]) public allfundings; // 所有的捐赠项目
 mapping(address => mapping(uint=> bool))  tfContribute;

 mapping(address => mapping(uint=> Funding))  tfCollectFunding;
 mapping(address =>Funding)  oneCollectFunding; //某个人的某个收藏项目
  mapping(address =>Funding[])  allCollectFunding; //某个人所有的收藏项目
  
//   mapping(address=>mapping(uint=>Funding)) tfContribute;  //所有被捐助项目
  mapping(address=>Funding[]) allContribute;
 //发布一个招募项目
  function newFunding(address  initiator, string memory title, string memory info, uint goal)  public   returns(uint) {
    // require(endTime > block.timestamp);

    numFundings = numFundings + 1;
    Funding storage f = fundings[numFundings];
    f.initiator = initiator;
    f.title = title;
    f.info = info;
    f.goal = goal;
    f.id=numFundings;
    // f.endTime = endTime;
    f.exist = true;
    f.success = false;
    f.amount = 0;
    f.numFunders = 0;
    f.numUses = 0;
    
    
    allfundings[1].push(f);
    return numFundings;
  }
  
  
  //删除已发布的公益·项目
  function deleteFunding(uint index)  public   returns(uint) {

require(index<=allfundings[1].length);

 
    for (uint i = index; i < allfundings[1].length; i++) {
      allfundings[1][i-1] = allfundings[1][i];
      fundings[i]=fundings[i+1];
    }
 
    delete allfundings[1][allfundings[1].length-1];
   delete fundings[numFundings];
   numFundings=numFundings-1;
  
  allfundings[1].pop();
   
//   numContributeFundings=numContributeFundings-1;
//  Funding storage f = fundings[index];
// funders[index].addr.transfer(funders[f.numFunders].amount);
//     funders[f.numFunders].amount=0;
    
    return numFundings;
  }

//用户捐助项目
  function contribute( uint ID) public payable {

require(fundings[ID].amount+msg.value<= fundings[ID].goal,"less than goal");
require(fundings[ID].exist=true,"Funding is not exist");


Funding storage f=fundings[ID];


  
    f.amount += msg.value;
  if(tfContribute[msg.sender][ID]==false){
      allContribute[msg.sender].push(f);
      tfContribute[msg.sender][ID]==true;
  }
    
 allfundings[1][ID-1].amount +=msg.value;
    funders[ID].amount += msg.value;
   numContributeFundings=numContributeFundings+1;
  
  msgFunder[msg.sender][ID].amount +=msg.value;
  

   
 
//     // 考虑本项目是否达成目标
//     if(fundings[ID].amount >= fundings[ID].goal){
//     fundings[ID].success = true;
    
//   for (uint i = ID; i < allfundings[1].length; i++) {
//       allfundings[1][i-1] = allfundings[1][i];
       
        
//     }
//     delete allfundings[1][allfundings[1].length-1];
//   numContributeFundings=numContributeFundings-1;
//     allfundings[1].pop();
  
//     }
    
    
    
    fundings[ID].numFunders = fundings[ID].numFunders + 1;
   
    


   
  }

  // 退钱
  function returnMoney(uint ID) public {
    require(ID <= numFundings && ID >= 1);
    require(fundings[ID].success == false);

    Funding storage f = fundings[ID];
    for(uint i=1; i<=f.numFunders; i++)
      if(funders[i].addr == msg.sender) {
        funders[i].addr.transfer(funders[i].amount);
        funders[i].amount = 0;
        f.amount -= funders[i].amount;
      }
  }
  //获得当前账户余额
  function getMsgBalance() public view returns (uint) {
    return address(msg.sender).balance;
  }
    //获得合约账户余额

  function getBalance() public view returns (uint) {
    return address(this).balance;
  }
  function getTime() public view returns(uint){
      return block.timestamp;
  }
  //获得某一个项目的被捐助所有金额
  function getMyFundings(address addr, uint ID) public view returns (uint) {
      uint res = 0;
      for(uint i=1; i<=fundings[ID].numFunders; i++) {
          if(funders[i].addr == addr)
            res += funders[i].amount;
      }
      return res;
  } 
  
  
//获取所有的公益项目
function GetAllfundings() public view returns(Funding[] memory){
    return allfundings[1];
}


//获得某个公益项目的被捐助金额

function GetFunder(uint key) public view returns(Funder memory){

  require(fundings[key].success==true);
  return funders[key];
}
//将某个公益项目加入某个人的收藏

function InsCollectFunding(uint key) public  returns(uint){

  require(fundings[key].exist==true,"Funding is not exist");
  require(tfCollectFunding[msg.sender][key].exist==false,"already insert!");
    numCollect=numCollect+1;
  Funding storage c=oneCollectFunding[msg.sender];

c.initiator=msg.sender; 
c.title=fundings[key].title;
c.info=fundings[key].info;
c.goal=fundings[key].goal;
c.exist=fundings[key].exist;
c.id=key;

allCollectFunding[msg.sender].push(oneCollectFunding[msg.sender]);
tfCollectFunding[msg.sender][key].exist=true;
  return numCollect;
}

//获得某个公益项目

function getFunding(uint key) public view returns(Funding memory){

   require(fundings[key].exist==true);

  return fundings[key];
}

//获得某个人收藏的所有公益项目

function getallCollectFunding(address addr) public view returns(Funding[] memory){

//   require(fundings[key].exist==true);
  return allCollectFunding[addr];
}

//删除某个人捐助的某个公益项目

function delOneCollectFunding(uint index) public  returns(bool){

    for (uint i = index; i < allCollectFunding[msg.sender].length; i++) {
      allCollectFunding[msg.sender][i-1] = allCollectFunding[msg.sender][i];
    
    }
 
    delete  allCollectFunding[msg.sender][allCollectFunding[msg.sender].length-1];
   
   numCollect=numCollect-1;
  
  allCollectFunding[msg.sender].pop();
  tfCollectFunding[msg.sender][index].exist=false;
//   require(fundings[key].exist==true);
  return true;
}

//获得某个人捐助的所有公益项目

function getallContribute(address addr) public view returns(Funding[] memory){

//   require(fundings[key].exist==true);


  return allContribute[addr];
}








}






