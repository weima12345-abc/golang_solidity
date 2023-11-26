<div align=center> 
<!-- <img src="https://avatars.githubusercontent.com/u/88427645?s=200&v=4" style="border-radius:10px"> -->
</div>

<center> Starter Kits 快速构建 DAPP </center>

<center> An Polygon Starter Kit Tutorial containing React, @web3-react, Infura. </center>
<div align=center> <a href="https://docs.matic.network/docs/develop/getting-started"> Developer Docs</a> - <a href="https://polygon-tutorial.soildstake.net"> Tutorial </a></div>    


链上公益致力于打造国内最安全真实的援助平台，该平台聚焦于国内公益信息的收集，捐献财物，等等功能为一体，为社会各界人
士提供一个可以无忧虑帮助需要帮助的人的平台。由于区块链的高安全性，高透明性以及不可篡改性，社会各界对区块链的关注度 也越来越高，所以我们联合其他高校共同打造了这个基于区块链的为广大社会弱势群体服务的系统平台。 链上公益的系统主要包 括公益信息展示系统，用户个人信息展示，该项目具有良好的可扩充性、高可用性，不可篡改性，由于公益事业的特殊性，这个行 业需要被社会各界人士所认同，这个系统可以完美解决问题。

项目工具：Golang(Gin) +智能合约

一。 功能介绍

1.登录注册
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/dl.png" width="600" alt="npm start">
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/zc.png" width="600" alt="npm start">
2.首页显示所有的需要帮助的人和事
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/index.png" width="600" alt="npm start">
3.用户捐赠页面，记录所有已经帮助的公益项目
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/w_xqgn.png" width="600" alt="npm start">
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/jz.png" width="600" alt="npm start">
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/xqgn.png" width="600" alt="npm start">

4.收藏页面，用户浏览到想要捐赠的公益项目但不确定时，选择收藏，之后再进行捐赠。
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/sc.png" width="600" alt="npm start">
5.个人页面，储存用户个人信息，如头像，用户名和密码；用户的账户余额等。
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/gr.png" width="600" alt="npm start">
<img src="https://github.com/weima12345-abc/golang_solidity/tree/main/img_goSodilidy/grgn.png" width="600" alt="npm start">


二。 项目启动


## 0. 环境配置 

### Requirements 

#### 安装 ganache 
```javascripts
Download Package from https://www.trufflesuite.com/ganache
```

#### 快速开始

1.在remix智能合约在线编译部署平台得到abi和合约地址。替换到相应的位置上。
2.开启本地私链，如 ganache和本地挖矿。如没有ganache和本地挖矿，请下载：
```javascripts
Download Package from https://www.trufflesuite.com/ganache
```
3.在前端也需要替换刚刚在remix编译部署过的abi和合约地址
4.完成这一切后，启动项目：
```javascripts
go run main.go
```
然后在浏览器打开：
```javascripts
http://localhost:8080/
```




