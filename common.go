package explorer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Metric struct {
	Total   float32
	Minimum float32
	Average float32
	Maximum float32
}

//TODO make Latence and throughput Metric
type Metrics struct {
	Throughput float32 `json:"Throughput (tx/sec)"`
	Latence    float32 `json:"Latence (sec)"`
	Size       Metric  `json:"Size (#tx)"`
}

func (metrics Metrics) Print() {
	json, err := json.MarshalIndent(metrics, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}

type Block struct {
	BlockNumber int
	BlockHash   string
	Producer    string
	TxNumber    uint
	Timestamp   int64
	Date        string
}

type Blockchain []Block

func (bchain Blockchain) Print() {
	json, err := json.MarshalIndent(bchain, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}

func (bchain Blockchain) Analyse() Metrics {
	metrics := Metrics{}
	var txCount uint

	temp, _ := time.Parse("2006-01-02 15:04:05", bchain[0].Date)
	tPrevious := temp.Unix()
	//init first value for min and Maximum
	txMin := bchain[0].TxNumber
	txMax := bchain[0].TxNumber

	var deltas int64

	for _, b := range bchain {
		// number tx
		// sum
		txCount += b.TxNumber
		//count max and Minimum
		if b.TxNumber < txMin {
			txMin = b.TxNumber
		}
		if b.TxNumber > txMax {
			txMax = b.TxNumber

		}

		tmp, _ := time.Parse("2006-01-02 15:04:05", b.Date)
		t := tmp.Unix()
		deltas += t - tPrevious
		tPrevious = t

	}
	txCount = txCount - bchain[0].TxNumber

	metrics.Size = Metric{float32(txCount), float32(txMin), float32(txCount) / float32(len(bchain)), float32(txMax)}

	metrics.Latence = float32(deltas) / float32(len(bchain)-1)
	metrics.Throughput = float32(txCount) / float32(deltas)
	return metrics
}
