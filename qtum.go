package explorer

import (
	"fmt"
	"log"

	"github.com/cryptohazard/explorer/qtum"
)

func QtumExplorer(rpc, user, pass string, blockStart, blockEnd int) Blockchain {
	client := qtum.QtumRPC{
		Host:    rpc, //"http://127.0.0.1:13889",
		ID:      "test",
		Version: "1.0",
		User:    user,
		Pass:    pass,
	}

	_, err := client.Call("getblockchaininfo", nil)
	if err != nil {
		log.Fatalf("error creating new Qtum rpc client: %v", err)
	}
	fmt.Println("Connected Successfully to Qtum rpc!")

	bchain := Blockchain{}
	for i := blockStart; i <= blockEnd; i++ {
		b := Block{}
		b.BlockNumber = i
		blockInfo, err := client.GetBlock(i)
		if err != nil {
			log.Fatal(err)
		}
		b.BlockHash = blockInfo.GetBlockHash()

		// coinbase tx is counted in
		b.TxNumber = blockInfo.GetNumberTx()

		b.Timestamp = blockInfo.GetTimestamp()
		b.Date = blockInfo.GetDate()

		bchain = append(bchain, b)

	}
	return bchain

}
