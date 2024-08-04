// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package app

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AppSensor is an auto generated low-level Go binding around an user-defined struct.
type AppSensor struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}

// AppMetaData contains all meta data concerning the App contract.
var AppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"createSensor\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structApp.Sensor\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getSensor\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structApp.Sensor\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSensors\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structApp.Sensor[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sensorByID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sensorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sensors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateSensor\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structApp.Sensor\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userSensors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506115468061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c80638519b29c116100595780638519b29c1461014d5780638badb85b14610180578063afb6e69f1461019e578063f90adc76146101d157610086565b8063321925661461008a578063412f2e82146100ba578063643344e8146100ed578063788044261461011d575b5f80fd5b6100a4600480360381019061009f9190610c42565b610201565b6040516100b19190610d9e565b60405180910390f35b6100d460048036038101906100cf9190610dbe565b6103df565b6040516100e49493929190610e4f565b60405180910390f35b61010760048036038101906101029190610efa565b6104bf565b6040516101149190610d9e565b60405180910390f35b61013760048036038101906101329190610dbe565b610757565b6040516101449190610d9e565b60405180910390f35b61016760048036038101906101629190610f81565b61087a565b6040516101779493929190610e4f565b60405180910390f35b610188610965565b6040516101959190610fbf565b60405180910390f35b6101b860048036038101906101b39190610dbe565b61096a565b6040516101c89493929190610e4f565b60405180910390f35b6101eb60048036038101906101e69190610fd8565b610a3b565b6040516101f8919061111e565b60405180910390f35b610209610bcc565b3373ffffffffffffffffffffffffffffffffffffffff1660025f8581526020019081526020015f206003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a190611188565b60405180910390fd5b8160025f8581526020019081526020015f206002018190555060025f8481526020019081526020015f206040518060800160405290815f82015481526020016001820180546102f8906111d3565b80601f0160208091040260200160405190810160405280929190818152602001828054610324906111d3565b801561036f5780601f106103465761010080835404028352916020019161036f565b820191905f5260205f20905b81548152906001019060200180831161035257829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050905092915050565b600181815481106103ee575f80fd5b905f5260205f2090600402015f91509050805f015490806001018054610413906111d3565b80601f016020809104026020016040519081016040528092919081815260200182805461043f906111d3565b801561048a5780601f106104615761010080835404028352916020019161048a565b820191905f5260205f20905b81548152906001019060200180831161046d57829003601f168201915b505050505090806002015490806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905084565b6104c7610bcc565b5f808154809291906104d890611230565b91905055505f60405180608001604052805f54815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018481526020013373ffffffffffffffffffffffffffffffffffffffff168152509050600181908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015560208201518160010190816105a19190611441565b50604082015181600201556060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505060035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015560208201518160010190816106759190611441565b50604082015181600201556060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050508060025f805481526020019081526020015f205f820151815f015560208201518160010190816106f89190611441565b50604082015181600201556060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050809150509392505050565b61075f610bcc565b60025f8381526020019081526020015f206040518060800160405290815f8201548152602001600182018054610794906111d3565b80601f01602080910402602001604051908101604052809291908181526020018280546107c0906111d3565b801561080b5780601f106107e25761010080835404028352916020019161080b565b820191905f5260205f20905b8154815290600101906020018083116107ee57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b6003602052815f5260405f208181548110610893575f80fd5b905f5260205f2090600402015f9150915050805f0154908060010180546108b9906111d3565b80601f01602080910402602001604051908101604052809291908181526020018280546108e5906111d3565b80156109305780601f1061090757610100808354040283529160200191610930565b820191905f5260205f20905b81548152906001019060200180831161091357829003601f168201915b505050505090806002015490806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905084565b5f5481565b6002602052805f5260405f205f91509050805f01549080600101805461098f906111d3565b80601f01602080910402602001604051908101604052809291908181526020018280546109bb906111d3565b8015610a065780601f106109dd57610100808354040283529160200191610a06565b820191905f5260205f20905b8154815290600101906020018083116109e957829003601f168201915b505050505090806002015490806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905084565b606060035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610bc1578382905f5260205f2090600402016040518060800160405290815f8201548152602001600182018054610ad3906111d3565b80601f0160208091040260200160405190810160405280929190818152602001828054610aff906111d3565b8015610b4a5780601f10610b2157610100808354040283529160200191610b4a565b820191905f5260205f20905b815481529060010190602001808311610b2d57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505081526020019060010190610a99565b505050509050919050565b60405180608001604052805f8152602001606081526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f80fd5b5f80fd5b5f819050919050565b610c2181610c0f565b8114610c2b575f80fd5b50565b5f81359050610c3c81610c18565b92915050565b5f8060408385031215610c5857610c57610c07565b5b5f610c6585828601610c2e565b9250506020610c7685828601610c2e565b9150509250929050565b610c8981610c0f565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610cd182610c8f565b610cdb8185610c99565b9350610ceb818560208601610ca9565b610cf481610cb7565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d2882610cff565b9050919050565b610d3881610d1e565b82525050565b5f608083015f830151610d535f860182610c80565b5060208301518482036020860152610d6b8282610cc7565b9150506040830151610d806040860182610c80565b506060830151610d936060860182610d2f565b508091505092915050565b5f6020820190508181035f830152610db68184610d3e565b905092915050565b5f60208284031215610dd357610dd2610c07565b5b5f610de084828501610c2e565b91505092915050565b610df281610c0f565b82525050565b5f82825260208201905092915050565b5f610e1282610c8f565b610e1c8185610df8565b9350610e2c818560208601610ca9565b610e3581610cb7565b840191505092915050565b610e4981610d1e565b82525050565b5f608082019050610e625f830187610de9565b8181036020830152610e748186610e08565b9050610e836040830185610de9565b610e906060830184610e40565b95945050505050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610eba57610eb9610e99565b5b8235905067ffffffffffffffff811115610ed757610ed6610e9d565b5b602083019150836001820283011115610ef357610ef2610ea1565b5b9250929050565b5f805f60408486031215610f1157610f10610c07565b5b5f84013567ffffffffffffffff811115610f2e57610f2d610c0b565b5b610f3a86828701610ea5565b93509350506020610f4d86828701610c2e565b9150509250925092565b610f6081610d1e565b8114610f6a575f80fd5b50565b5f81359050610f7b81610f57565b92915050565b5f8060408385031215610f9757610f96610c07565b5b5f610fa485828601610f6d565b9250506020610fb585828601610c2e565b9150509250929050565b5f602082019050610fd25f830184610de9565b92915050565b5f60208284031215610fed57610fec610c07565b5b5f610ffa84828501610f6d565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f608083015f8301516110415f860182610c80565b50602083015184820360208601526110598282610cc7565b915050604083015161106e6040860182610c80565b5060608301516110816060860182610d2f565b508091505092915050565b5f611097838361102c565b905092915050565b5f602082019050919050565b5f6110b582611003565b6110bf818561100d565b9350836020820285016110d18561101d565b805f5b8581101561110c57848403895281516110ed858261108c565b94506110f88361109f565b925060208a019950506001810190506110d4565b50829750879550505050505092915050565b5f6020820190508181035f83015261113681846110ab565b905092915050565b7f73656e646572206973206e6f74206f776e6572000000000000000000000000005f82015250565b5f611172601383610df8565b915061117d8261113e565b602082019050919050565b5f6020820190508181035f83015261119f81611166565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111ea57607f821691505b6020821081036111fd576111fc6111a6565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61123a82610c0f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361126c5761126b611203565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826112c5565b61130a86836112c5565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61134561134061133b84610c0f565b611322565b610c0f565b9050919050565b5f819050919050565b61135e8361132b565b61137261136a8261134c565b8484546112d1565b825550505050565b5f90565b61138661137a565b611391818484611355565b505050565b5b818110156113b4576113a95f8261137e565b600181019050611397565b5050565b601f8211156113f9576113ca816112a4565b6113d3846112b6565b810160208510156113e2578190505b6113f66113ee856112b6565b830182611396565b50505b505050565b5f82821c905092915050565b5f6114195f19846008026113fe565b1980831691505092915050565b5f611431838361140a565b9150826002028217905092915050565b61144a82610c8f565b67ffffffffffffffff81111561146357611462611277565b5b61146d82546111d3565b6114788282856113b8565b5f60209050601f8311600181146114a9575f8415611497578287015190505b6114a18582611426565b865550611508565b601f1984166114b7866112a4565b5f5b828110156114de578489015182556001820191506020850194506020810190506114b9565b868310156114fb57848901516114f7601f89168261140a565b8355505b6001600288020188555050505b50505050505056fea264697066735822122041faa28d7f4ba171e1beb835602cbc18dc051217f7521cf7bb82f0ca94cecbb864736f6c63430008190033",
}

// AppABI is the input ABI used to generate the binding from.
// Deprecated: Use AppMetaData.ABI instead.
var AppABI = AppMetaData.ABI

// AppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AppMetaData.Bin instead.
var AppBin = AppMetaData.Bin

// DeployApp deploys a new Ethereum contract, binding an instance of App to it.
func DeployApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *App, error) {
	parsed, err := AppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AppBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

// App is an auto generated Go binding around an Ethereum contract.
type App struct {
	AppCaller     // Read-only binding to the contract
	AppTransactor // Write-only binding to the contract
	AppFilterer   // Log filterer for contract events
}

// AppCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppSession struct {
	Contract     *App              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppCallerSession struct {
	Contract *AppCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppTransactorSession struct {
	Contract     *AppTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRaw struct {
	Contract *App // Generic contract binding to access the raw methods on
}

// AppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppCallerRaw struct {
	Contract *AppCaller // Generic read-only contract binding to access the raw methods on
}

// AppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppTransactorRaw struct {
	Contract *AppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApp creates a new instance of App, bound to a specific deployed contract.
func NewApp(address common.Address, backend bind.ContractBackend) (*App, error) {
	contract, err := bindApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

// NewAppCaller creates a new read-only instance of App, bound to a specific deployed contract.
func NewAppCaller(address common.Address, caller bind.ContractCaller) (*AppCaller, error) {
	contract, err := bindApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppCaller{contract: contract}, nil
}

// NewAppTransactor creates a new write-only instance of App, bound to a specific deployed contract.
func NewAppTransactor(address common.Address, transactor bind.ContractTransactor) (*AppTransactor, error) {
	contract, err := bindApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppTransactor{contract: contract}, nil
}

// NewAppFilterer creates a new log filterer instance of App, bound to a specific deployed contract.
func NewAppFilterer(address common.Address, filterer bind.ContractFilterer) (*AppFilterer, error) {
	contract, err := bindApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppFilterer{contract: contract}, nil
}

// bindApp binds a generic wrapper to an already deployed contract.
func bindApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.AppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.contract.Transact(opts, method, params...)
}

// GetSensor is a free data retrieval call binding the contract method 0x78804426.
//
// Solidity: function getSensor(uint256 id) view returns((uint256,string,uint256,address))
func (_App *AppCaller) GetSensor(opts *bind.CallOpts, id *big.Int) (AppSensor, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getSensor", id)

	if err != nil {
		return *new(AppSensor), err
	}

	out0 := *abi.ConvertType(out[0], new(AppSensor)).(*AppSensor)

	return out0, err

}

// GetSensor is a free data retrieval call binding the contract method 0x78804426.
//
// Solidity: function getSensor(uint256 id) view returns((uint256,string,uint256,address))
func (_App *AppSession) GetSensor(id *big.Int) (AppSensor, error) {
	return _App.Contract.GetSensor(&_App.CallOpts, id)
}

// GetSensor is a free data retrieval call binding the contract method 0x78804426.
//
// Solidity: function getSensor(uint256 id) view returns((uint256,string,uint256,address))
func (_App *AppCallerSession) GetSensor(id *big.Int) (AppSensor, error) {
	return _App.Contract.GetSensor(&_App.CallOpts, id)
}

// GetSensors is a free data retrieval call binding the contract method 0xf90adc76.
//
// Solidity: function getSensors(address user) view returns((uint256,string,uint256,address)[])
func (_App *AppCaller) GetSensors(opts *bind.CallOpts, user common.Address) ([]AppSensor, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getSensors", user)

	if err != nil {
		return *new([]AppSensor), err
	}

	out0 := *abi.ConvertType(out[0], new([]AppSensor)).(*[]AppSensor)

	return out0, err

}

// GetSensors is a free data retrieval call binding the contract method 0xf90adc76.
//
// Solidity: function getSensors(address user) view returns((uint256,string,uint256,address)[])
func (_App *AppSession) GetSensors(user common.Address) ([]AppSensor, error) {
	return _App.Contract.GetSensors(&_App.CallOpts, user)
}

// GetSensors is a free data retrieval call binding the contract method 0xf90adc76.
//
// Solidity: function getSensors(address user) view returns((uint256,string,uint256,address)[])
func (_App *AppCallerSession) GetSensors(user common.Address) ([]AppSensor, error) {
	return _App.Contract.GetSensors(&_App.CallOpts, user)
}

// SensorByID is a free data retrieval call binding the contract method 0xafb6e69f.
//
// Solidity: function sensorByID(uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppCaller) SensorByID(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "sensorByID", arg0)

	outstruct := new(struct {
		Id    *big.Int
		Name  string
		Value *big.Int
		User  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SensorByID is a free data retrieval call binding the contract method 0xafb6e69f.
//
// Solidity: function sensorByID(uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppSession) SensorByID(arg0 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	return _App.Contract.SensorByID(&_App.CallOpts, arg0)
}

// SensorByID is a free data retrieval call binding the contract method 0xafb6e69f.
//
// Solidity: function sensorByID(uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppCallerSession) SensorByID(arg0 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	return _App.Contract.SensorByID(&_App.CallOpts, arg0)
}

// SensorIndex is a free data retrieval call binding the contract method 0x8badb85b.
//
// Solidity: function sensorIndex() view returns(uint256)
func (_App *AppCaller) SensorIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "sensorIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SensorIndex is a free data retrieval call binding the contract method 0x8badb85b.
//
// Solidity: function sensorIndex() view returns(uint256)
func (_App *AppSession) SensorIndex() (*big.Int, error) {
	return _App.Contract.SensorIndex(&_App.CallOpts)
}

// SensorIndex is a free data retrieval call binding the contract method 0x8badb85b.
//
// Solidity: function sensorIndex() view returns(uint256)
func (_App *AppCallerSession) SensorIndex() (*big.Int, error) {
	return _App.Contract.SensorIndex(&_App.CallOpts)
}

// Sensors is a free data retrieval call binding the contract method 0x412f2e82.
//
// Solidity: function sensors(uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppCaller) Sensors(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "sensors", arg0)

	outstruct := new(struct {
		Id    *big.Int
		Name  string
		Value *big.Int
		User  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Sensors is a free data retrieval call binding the contract method 0x412f2e82.
//
// Solidity: function sensors(uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppSession) Sensors(arg0 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	return _App.Contract.Sensors(&_App.CallOpts, arg0)
}

// Sensors is a free data retrieval call binding the contract method 0x412f2e82.
//
// Solidity: function sensors(uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppCallerSession) Sensors(arg0 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	return _App.Contract.Sensors(&_App.CallOpts, arg0)
}

// UserSensors is a free data retrieval call binding the contract method 0x8519b29c.
//
// Solidity: function userSensors(address , uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppCaller) UserSensors(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "userSensors", arg0, arg1)

	outstruct := new(struct {
		Id    *big.Int
		Name  string
		Value *big.Int
		User  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// UserSensors is a free data retrieval call binding the contract method 0x8519b29c.
//
// Solidity: function userSensors(address , uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppSession) UserSensors(arg0 common.Address, arg1 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	return _App.Contract.UserSensors(&_App.CallOpts, arg0, arg1)
}

// UserSensors is a free data retrieval call binding the contract method 0x8519b29c.
//
// Solidity: function userSensors(address , uint256 ) view returns(uint256 id, string name, uint256 value, address user)
func (_App *AppCallerSession) UserSensors(arg0 common.Address, arg1 *big.Int) (struct {
	Id    *big.Int
	Name  string
	Value *big.Int
	User  common.Address
}, error) {
	return _App.Contract.UserSensors(&_App.CallOpts, arg0, arg1)
}

// CreateSensor is a paid mutator transaction binding the contract method 0x643344e8.
//
// Solidity: function createSensor(string name, uint256 value) returns((uint256,string,uint256,address))
func (_App *AppTransactor) CreateSensor(opts *bind.TransactOpts, name string, value *big.Int) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "createSensor", name, value)
}

// CreateSensor is a paid mutator transaction binding the contract method 0x643344e8.
//
// Solidity: function createSensor(string name, uint256 value) returns((uint256,string,uint256,address))
func (_App *AppSession) CreateSensor(name string, value *big.Int) (*types.Transaction, error) {
	return _App.Contract.CreateSensor(&_App.TransactOpts, name, value)
}

// CreateSensor is a paid mutator transaction binding the contract method 0x643344e8.
//
// Solidity: function createSensor(string name, uint256 value) returns((uint256,string,uint256,address))
func (_App *AppTransactorSession) CreateSensor(name string, value *big.Int) (*types.Transaction, error) {
	return _App.Contract.CreateSensor(&_App.TransactOpts, name, value)
}

// UpdateSensor is a paid mutator transaction binding the contract method 0x32192566.
//
// Solidity: function updateSensor(uint256 id, uint256 value) returns((uint256,string,uint256,address))
func (_App *AppTransactor) UpdateSensor(opts *bind.TransactOpts, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "updateSensor", id, value)
}

// UpdateSensor is a paid mutator transaction binding the contract method 0x32192566.
//
// Solidity: function updateSensor(uint256 id, uint256 value) returns((uint256,string,uint256,address))
func (_App *AppSession) UpdateSensor(id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _App.Contract.UpdateSensor(&_App.TransactOpts, id, value)
}

// UpdateSensor is a paid mutator transaction binding the contract method 0x32192566.
//
// Solidity: function updateSensor(uint256 id, uint256 value) returns((uint256,string,uint256,address))
func (_App *AppTransactorSession) UpdateSensor(id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _App.Contract.UpdateSensor(&_App.TransactOpts, id, value)
}
