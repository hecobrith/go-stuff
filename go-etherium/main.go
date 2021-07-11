package main

import (
    "context"
    "fmt"
    "log"
	"github.com/hecobrith/go-etherium/wallet"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
    if err != nil {
        log.Fatal(err)
    }

    account := common.HexToAddress("0x7B73A0dFf76402dCae018d58CCc7F8D76d9192b9")
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balance) // 25893180161173005034

    pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
    fmt.Println(pendingBalance) // 25729324269165216042

	wallet.NewWallet()
}