package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// https://coston2-api.flare.network/ext/bc/C/rpc

const (
	xrp   = "testXRP"
	ltc   = "testLTC"
	xlm   = "testXLM"
	doge  = "testDOGE"
	ada   = "testADA"
	algo  = "testALGO"
	btc   = "testBTC"
	eth   = "testETH"
	fil   = "testFIL"
	flr   = "testFLR"
	arb   = "testARB"
	avax  = "testAVAX"
	bnb   = "testBNB"
	matic = "testMATIC"
	sol   = "testSOL"
	usdc  = "testUSDC"
	usdt  = "testUSDT"
	xdc   = "testXDC"
)

type ChainQuery struct {
	client       *ethclient.Client
	auth         *bind.TransactOpts
	mumbaiClient *ethclient.Client
	symbols      []string
}

type Price struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

func NewChainQuery() ChainQuery {
	fmt.Println("aadr:", Config.FlareNode)
	client, err := ethclient.Dial(Config.FlareNode)
	if err != nil {
		panic(err)
	}
	auth, err := GetAuth(client)
	if err != nil {
		panic(err)
	}
	mumbaiClient, err := ethclient.Dial(Config.MumbaiUrl)
	if err != nil {
		panic(err)
	}
	return ChainQuery{client: client, mumbaiClient: mumbaiClient, auth: auth, symbols: []string{xrp, ltc, xlm, doge, ada, algo, btc, eth, fil, flr, arb, avax, bnb, matic, sol, usdc, usdt, xdc}}
}

func GetAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(Config.Key)
	if err != nil {
		return nil, err
	}

	pubKey := privKey.Public()
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(114))

	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	return auth, nil
}

func (c *ChainQuery) UpdatedAuth() *bind.TransactOpts {
	nonce, err := c.client.PendingNonceAt(context.Background(), c.auth.From)
	if err != nil {
		panic(err)
	}

	c.auth.Nonce = big.NewInt(int64(nonce))
	return c.auth
}

func (c *ChainQuery) QueryPrices() ([]Price, error) {
	priceData, err := NewUnfuckPriceData(common.HexToAddress(Config.PriceDataAddr), c.client)
	if err != nil {
		return []Price{}, err
	}

	var prices []Price

	for _, s := range c.symbols {
		pInfo, err := priceData.LatestPriceForToken(nil, s)
		if err != nil {
			continue
		}

		zeros := strings.Repeat("0", int(pInfo.Decimals.Int64()))
		toDiv, err := strconv.Atoi("1" + zeros)
		if err != nil {
			continue
		}

		price := float64(pInfo.Price.Int64()) / float64(toDiv)
		fmt.Println("price: ", s, price)

		prices = append(prices, Price{Symbol: s, Price: price})
	}

	return prices, nil
}

func (c *ChainQuery) RefreshPricesAndCreateReport() {
	lastPrices, err := getLastPrices()
	if err != nil {
		fmt.Println("error getting last prices: ", err)
		return
	}
	c.RefreshPrices()

	prices, err := c.QueryPrices()
	if err != nil {
		fmt.Println("error querying prices: ", prices)
		return
	}

	if err := S.PSaver.Save(prices, time.Now().Unix()); err != nil {
		fmt.Println("error saving prices: ", err)
	}

	prompt := G.GeneratePricePrompt(lastPrices.Prices, prices)
	fmt.Println(prompt)
	resp := G.GetResponse(prompt)

	if err := S.PASaver.Save(resp, time.Now().Unix()); err != nil {
		fmt.Println("error saving price analysis: ", err)
	}
}

func (c *ChainQuery) RefreshPrices() {
	priceData, err := NewUnfuckPriceData(common.HexToAddress(Config.PriceDataAddr), c.client)
	if err != nil {
		fmt.Println("error creating price data contract: ", err)
		return
	}

	_, err = priceData.RefreshPrices(c.auth)
	if err != nil {
		fmt.Println("error refreshing prices: ", err)
	}
	// fmt.Println("prices refreshed, tx hash: ", tx.Hash().String())
}

func (c *ChainQuery) IsMember(address string) bool {
	core, err := NewUnfuckCore(common.HexToAddress(Config.UnfuckCore), c.mumbaiClient)
	if err != nil {
		fmt.Println("error init core: ", err)
		return false
	}
	isMember, err := core.Members(nil, common.HexToAddress(address))
	if err != nil {
		fmt.Println("error get is member: ", err)
		return false
	}
	return isMember
}

func getLastPrices() (PriceInfo, error) {
	fBytes, err := os.ReadFile("prices.json")
	if err != nil {
		return PriceInfo{}, err
	}

	var priceInfo []PriceInfo
	if err := json.Unmarshal(fBytes, &priceInfo); err != nil {
		return PriceInfo{}, err
	}

	if len(priceInfo) > 0 {
		return priceInfo[0], nil
	} else {
		return PriceInfo{}, nil
	}

}
