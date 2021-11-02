package explorer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Result struct {
	Chain      string
	Type       string
	BlockStart int
	BlockEnd   int
	Result     []float32
	Average    float32
}

type Latency struct {
	Type      string
	Blocktime []float32
	Average   float32
}

func (s *Result) ShowResult() {
	fmt.Println("Show Result of Tests :")
	fmt.Println("\tChain\t:", s.Chain)
	fmt.Println("\tType\t:", s.Type)
	fmt.Println("\tBlockStart\t:", s.BlockStart)
	fmt.Println("\tBlockEnd\t:", s.BlockEnd)
	fmt.Println("\tResult\t:", s.Result)
	fmt.Println("\tAverage\t:", s.Average)

}

func (l *Latency) TransLatency() {
	fmt.Println("Show latency of Tests :")
	fmt.Println("\tType\t:", l.Type)
	fmt.Println("\tBlocktime\t:", l.Blocktime)
	fmt.Println("\tAverage\t:", l.Average)
}

func PrintRes() float32 {
	showRes := Result{
		"Ethereum",
		"throughput",
		7950400,
		7950401,
		[]float32{7.592593, 7.823529, 12.923077, 10.125000, 10.772727, 0.000000, 189.000000, 1.666667, 108.500000, 4.391304},
		35.3,
	}
	showRes.ShowResult()

	b, err := json.Marshal(showRes)
	if err != nil {
		fmt.Println("error:", err)
	}

	err = ioutil.WriteFile("result.json", b, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("data", data)

	//Calcule latence
	showLat := Latency{
		"Latency",
		[]float32{27, 17, 13, 16, 22, 16, 1, 9, 2, 23},
		14.6,
	}
	showLat.TransLatency()

	t, err := json.Marshal(showLat)
	if err != nil {
		fmt.Println("error:", err)
	}

	err = ioutil.WriteFile("latency.json", t, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	var btime interface{}
	err = json.Unmarshal(t, &btime)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("blocktime", btime)

	return 1.0
}
