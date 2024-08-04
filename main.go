package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	app "github.com/emiliomarin/blockchain-sample/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type wallet struct {
	PubKey  common.Address
	PrivKey *ecdsa.PrivateKey
}

func main() {
	// Create a contract client using the deployed contract address and the ganache client
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x70e0a92C340C3Ae7b3f9e3B17a9F9B900497DFB9")
	appClient, err := app.NewApp(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	wallets := generateWallets()
	nSensors := 3

	// Create 3 sensors with each wallet
	// Get all sensors for the wallet
	// Update first sensor in array
	// Get sensor
	for _, wallet := range wallets {
		printAddressBalance(client, wallet.PubKey.Hex())
		for i := 0; i < nSensors; i++ {
			opts, err := buildTxOpts(client, wallet.PubKey, wallet.PrivKey)
			if err != nil {
				log.Fatal("error building tx opts", err)
			}

			_, err = appClient.CreateSensor(opts, "Temp Sensor", big.NewInt(10))
			if err != nil {
				log.Fatal("error creating sensor ", err)
			}
		}

		sensors, err := appClient.GetSensors(&bind.CallOpts{Pending: false}, wallet.PubKey)
		if err != nil {
			log.Fatal("error getting sensors", err)
		}
		fmt.Println(sensors)

		opts, err := buildTxOpts(client, wallet.PubKey, wallet.PrivKey)
		if err != nil {
			log.Fatal("error building tx opts", err)
		}
		_, err = appClient.UpdateSensor(opts, sensors[0].Id, big.NewInt(100))
		if err != nil {
			log.Fatal("error updating sensor ", err)
		}

		sensor, err := appClient.GetSensor(&bind.CallOpts{Pending: true}, sensors[0].Id)
		if err != nil {
			log.Fatal("error getting sensors", err)
		}
		fmt.Println("Updated sensor value is: ", sensor.Value)

		printAddressBalance(client, wallet.PubKey.Hex())
	}

	// try to update sensor of other user, should fail
	sensors, err := appClient.GetSensors(&bind.CallOpts{Pending: false}, wallets[0].PubKey)
	if err != nil {
		log.Fatal("error getting sensors", err)
	}
	opts, err := buildTxOpts(client, wallets[1].PubKey, wallets[1].PrivKey)
	if err != nil {
		log.Fatal("error building tx opts", err)
	}
	_, err = appClient.UpdateSensor(opts, sensors[0].Id, big.NewInt(999))
	fmt.Println("Should fail to update: ", err)

	sensor, err := appClient.GetSensor(&bind.CallOpts{Pending: true}, sensors[0].Id)
	if err != nil {
		log.Fatal("error getting sensors", err)
	}
	fmt.Println("Sensor should remain at 100: ", sensor.Value)
}

func buildTxOpts(client *ethclient.Client, publicKey common.Address, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func generateWallets() []wallet {
	pubKeys := []string{
		"0xf6B83Fee9EaAe7bE4b07E3E014ff9C8991CE1332",
		"0x5c0435F477f9E8987CfA2F2AFc154cF4762DEEc5",
	}
	privKeys := []string{
		"0xc46fe947411ab22fa7d151d12534f4d540978f2ed911877ea16a9e08a88b2b8f",
		"0x54663ae6feef54a83361cc1fd09c300b1a3d7d7196070e2b364be60af2164657",
	}
	wallets := []wallet{}

	for i := range pubKeys {
		userPublicKey := common.HexToAddress(pubKeys[i])
		privateKeyBin, _ := hexutil.Decode(privKeys[i])
		userPrivateKey, err := crypto.ToECDSA(privateKeyBin)
		if err != nil {
			log.Fatal(err)
		}
		wallets = append(wallets, wallet{
			PubKey:  userPublicKey,
			PrivKey: userPrivateKey,
		})
	}
	return wallets
}

func printAddressBalance(client *ethclient.Client, address string) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("error getting balance", err)
	}
	fmt.Printf("Address %s Balance: %s\n", address, balance)
}
