package rpc

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func NewClient(uri string) (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error

	for i := 5; i > 0; i-- {
		client, err = ethclient.Dial(uri)
		if err == nil {
			return client, nil
		}
		time.Sleep(time.Second * 10)
		fmt.Println("retry connection")
	}
	return nil, err
}
