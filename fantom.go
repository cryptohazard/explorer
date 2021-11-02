package explorer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	ethclient1922 "github.com/ethereum/go-ethereum/ethclient/1922"
)

func FtmExplorer(rpc, user, pass string, blockStart, blockEnd int) Blockchain {
	client, err := ethclient1922.Dial(rpc + user + pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected Successfully to Eth rpc!")
	bchain := Blockchain{}
	for i := blockStart; i <= blockEnd; i++ {
		b := Block{}
		b.BlockNumber = i
		blockInfo, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(uint64(i)))
		if err != nil {
			log.Fatal(err)
		}
		b.BlockHash = blockInfo.Hash().String()

		blockHeader, err := client.HeaderByNumber(context.Background(), new(big.Int).SetUint64(uint64(i)))
		if err != nil {
			log.Fatal(err)
		}
		b.Producer = blockHeader.Coinbase.String()

		// Call TransactionCount to return the count of transactions in the block.
		b.TxNumber, err = client.TransactionCount(context.Background(), blockInfo.Hash())
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(blockInfo)

		//TODO REDO this
		timeBlock := int64(blockInfo.Time())
		//fmt.Println(timeTemp)
		tm := time.Unix(timeBlock, 0)
		b.Timestamp = timeBlock
		b.Date = tm.Format("2006-01-02 15:04:05.000")

		bchain = append(bchain, b)

	}
	return bchain

}
