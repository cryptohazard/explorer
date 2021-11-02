package explorer

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func BtcExplorer(rpc, user, pass string, blockStart, blockEnd int) Blockchain {
	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         rpc,
		User:         user,
		Pass:         pass,
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

	fmt.Println("Connected Successfully to Btc rpc!")
	bchain := Blockchain{}
	for i := blockStart; i <= blockEnd; i++ {
		b := Block{}
		b.BlockNumber = i
		hash, err := client.GetBlockHash(int64(i))
		if err != nil {
			log.Fatal(err)
		}
		blockInfo, err := client.GetBlock(hash)
		if err != nil {
			log.Fatal(err)
		}
		b.BlockHash = hash.String()

		// coinbase tx is counted in
		b.TxNumber = uint(len(blockInfo.Transactions))
		//coinbaseTx:=blockInfo.Transactions[0].TxOut[0]

		b.Timestamp = int64(blockInfo.Header.Timestamp.Unix())
		b.Date = blockInfo.Header.Timestamp.Format("2006-01-02 15:04:05.000")

		bchain = append(bchain, b)

	}
	return bchain

}
