package explorer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	ethclient11011 "github.com/ethereum/go-ethereum/ethclient/11011"
)

func EthExplorer(rpc, user, pass string, blockStart, blockEnd int) Blockchain {
	client, err := ethclient11011.Dial(rpc + user + pass)
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

//older function
func EthThroughput(blockStart, blockEnd int) float32 {
	var sum float32
	var throughPut float32
	for i := blockStart; i <= blockEnd; i++ {
		sum = sum + EthBlockThroughput(i)
	}

	fmt.Println("Sum of Speed: ", sum)
	n := float32(blockEnd - blockStart + 1)
	fmt.Printf("We test %f blocks! ", n)
	throughPut = sum / n

	return throughPut

}

//TODO do case blocknumber=0
func EthBlockThroughput(blockNumber int) float32 {
	client, err := ethclient11011.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected Successfully!")

	//Get all informations of the block by calling Call the client's BlockByNumber method
	temp := new(big.Int).SetUint64(uint64(blockNumber)) // The number of lastest block in type big.Int, we use it in BlockByNumber
	block, err := client.BlockByNumber(context.Background(), temp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Block's number is :", blockNumber)
	fmt.Println("Block info: ", block)
	//Number of parentBlock
	blockParent := blockNumber - 1

	//Get all informations of the parentBlock
	temp2 := new(big.Int).SetUint64(uint64(blockParent)) // The number of parent block of type big.Int, we use it in BlockByNumber
	parentBlock, err := client.BlockByNumber(context.Background(), temp2)
	if err != nil {
		log.Fatal(err)
	}

	// Call TransactionCount to return the count of transactions in the block.
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	countTx := float32(count)
	fmt.Printf("This block has %v transactions! \n", count)

	timeBlock := block.Time()
	var timeTemp int64 = int64(timeBlock)
	tm := time.Unix(timeTemp, 0)
	fmt.Println("Timestamps of the block: ", timeBlock)
	fmt.Println("Real time of the block: ", tm.Format("2006-01-02 15:04:05 "))

	timeParentBlock := parentBlock.Time()
	var timeTemp2 int64 = int64(timeParentBlock)

	deltaTime := float32(timeTemp - timeTemp2)
	fmt.Printf("The blocktime is: %f secs\n", deltaTime)

	speed := (countTx) / deltaTime
	fmt.Printf("The ethereum blockchain currently supports roughly %f transactions per second\n", speed)

	return speed

}
