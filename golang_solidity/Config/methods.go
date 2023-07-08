package Config

import (
	Faucet "charity/Project/go"
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/core/types"

	// "fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	chainID       = 987655678
	gasLimit      = uint64(21000)
	FaucetAddress = "0x0b5D04a4b97AeCc7E8D636d233a895a6Ca9E02e7"
)

// GetClient 获取 ethclient.Client 对象
func GetClient() (*ethclient.Client, error) {
	rpcDial, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(rpcDial)
	return client, nil
}

// GetFacetTx 获取 Faucet.ApinameTransactor 对象
func GetFacetTx(client *ethclient.Client) (*Faucet.CalldemoTransactor, error) {
	contract, err := Faucet.NewCalldemoTransactor(common.HexToAddress(FaucetAddress), client)
	return contract, err
}

// GetCallFaceTx 获取 Faucet.ApinameCaller 对象
func GetCallFaceTx(client ethclient.Client) (*Faucet.CalldemoCaller, error) {
	contract, err := Faucet.NewCalldemoCaller(common.HexToAddress(FaucetAddress), &client)
	return contract, err
}

// 初步获取 *bind.TransactOpts 对象
func setopts(client *ethclient.Client, privateKey *ecdsa.PrivateKey, address common.Address, value int64) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		panic(err)
	}
	nonce, err := Getnonce(client, address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := GetgasPrice(client)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}

// GetTxopts 获取最终 *bind.TransactOpts 对象
func GetTxopts(client *ethclient.Client, value int64) *bind.TransactOpts {
	privateKey, publicKey := GetPublicKeyPrivateKey()
	opts := setopts(client, privateKey, publicKey, value)
	return opts
}

// Getnonce 获取 nonce
func Getnonce(client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	} else {
		return nonce, nil
	}

}

// GetPublicKeyPrivateKey 获取公钥和私钥
func GetPublicKeyPrivateKey() (*ecdsa.PrivateKey, common.Address) {
	PK := "e2298eb11d40f767c515af75cd9f4b14e6c4d35685768afde6b0c3c4ed8945de"
	privateKey, err := crypto.HexToECDSA(PK)
	if err != nil {
		panic(err)
	}
	publicKey := common.HexToAddress("0xE13BE5b9d25f42fC0a9f391d58962cB947F27A81")
	return privateKey, publicKey
}

// GetgasPrice 获取 gasPrice
func GetgasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	} else {
		return gasPrice, nil
	}

}

// GetContractAmount 调用 charity合约合约 查看合约ETH数目
func GetContractAmount(contract *Faucet.CalldemoCaller) (*big.Int, error) {

	res, err := contract.GetMsgBalance(nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return res, nil
}

// GetallFundings 调用 charity合约 查看合约所有公益项目，并在前台展示
func GetallFundings(contract *Faucet.CalldemoCaller) ([]Faucet.CharityFunding, error) {

	res, err := contract.GetAllfundings(nil)
	if err != nil {
		return *new([]Faucet.CharityFunding), err
	}
	// fmt.Println(res)
	return res, nil
}

// newFunding 调用 charity合约 查看合约某个公益项目，并在前台展示
func GetnewFunding(contract *Faucet.CalldemoCaller, id *big.Int) (Faucet.CharityFunding, error) {

	res, err := contract.GetFunding(nil, id)
	if err != nil {
		return *new(Faucet.CharityFunding), err
	}
	return res, nil
}

// contribute 调用 charity合约 用户捐助
func Contribute(client *ethclient.Client, contract *Faucet.CalldemoTransactor, id *big.Int, value int64) (*types.Transaction, error) {
	opts := GetTxopts(client, value)
	res, err := contract.Contribute(opts, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewFunding 调用 charity合约 新增公益项目
func NewFunding(client *ethclient.Client, contract *Faucet.CalldemoTransactor, initiator common.Address, title string, info string, goal *big.Int) (*types.Transaction, error) {
	opts := GetTxopts(client, 0)
	res, err := contract.NewFunding(opts, initiator, title, info, goal)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelFunding 调用 charity合约 删除公益项目
func DelFunding(client *ethclient.Client, contract *Faucet.CalldemoTransactor,  goal *big.Int) (*types.Transaction, error) {
	opts := GetTxopts(client, 0)
	res, err := contract.DeleteFunding(opts, goal)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelFunding 调用 charity合约 删除公益项目
func DelInscollect(client *ethclient.Client, contract *Faucet.CalldemoTransactor,  goal *big.Int) (*types.Transaction, error) {
	opts := GetTxopts(client, 0)
	res, err := contract.DelOneCollectFunding(opts, goal)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// InsFunding 调用 charity合约 把公益项目加入收藏
func InsFunding(client *ethclient.Client, contract *Faucet.CalldemoTransactor,  goal *big.Int) (*types.Transaction, error) {
	opts := GetTxopts(client, 0)
	res, err := contract.InsCollectFunding(opts, goal)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetallInsFunding 调用 charity合约 获取已加入收藏的公益项目
func GetallInsFunding(contract *Faucet.CalldemoCaller,  initiator common.Address) ([]Faucet.CharityFunding, error) {

	res, err := contract.GetallCollectFunding(nil,initiator)
	if err != nil {
		return *new([]Faucet.CharityFunding), err
	}
	return res, nil
}

// GetallContribute 调用 charity合约 获取用户所有已捐助的项目
func GetallContribute(contract *Faucet.CalldemoCaller,  initiator common.Address) ([]Faucet.CharityFunding, error) {

	res, err := contract.GetallContribute(nil,initiator)
	if err != nil {
		return *new([]Faucet.CharityFunding), err
	}
	return res, nil
}

// GetallContribute 调用 charity合约 获取用户所有已捐助的项目
func DetOneFamout(contract *Faucet.CalldemoCaller,  initiator common.Address,goal *big.Int) (Faucet.CharityFunder, error) {

	res, err := contract.MsgFunder(nil,initiator,goal)
	if err != nil {
		return *new(Faucet.CharityFunder), err
	}
	return res, nil
}


