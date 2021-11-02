package explorer

import (
	"context"
	"fmt"
	"log"

	eos "github.com/eoscanada/eos-go"
)

func EosExplorer(rpc, user, pass string, blockStart, blockEnd int) Blockchain {
	client := eos.New(rpc)
	ctx := context.TODO()

	info, err := client.GetInfo(ctx)
	if err != nil {
		log.Fatalf("error creating new Eos rpc client: %v", err)
	}
	fmt.Println("Connected Successfully to Eos rpc!\n", info)
	bchain := Blockchain{}
	for i := blockStart; i <= blockEnd; i++ {
		b := Block{}
		b.BlockNumber = i
		blockInfo, err := client.GetBlockByNum(ctx, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		b.BlockHash = blockInfo.ID.String()
		b.Producer = string(blockInfo.SignedBlock.SignedBlockHeader.BlockHeader.Producer)

		// coinbase tx is counted in
		b.TxNumber = uint(len(blockInfo.SignedBlock.Transactions))

		b.Timestamp = int64(blockInfo.SignedBlock.BlockHeader.Timestamp.Unix())
		b.Date = blockInfo.SignedBlock.BlockHeader.Timestamp.Format("2006-01-02 15:04:05.000")

		bchain = append(bchain, b)

	}
	return bchain

}
