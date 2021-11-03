package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/cryptohazard/explorer"
)

func main() {
	chain := flag.String("c", "btc", "chain :eth,btc,qtum,eos")
	metric := flag.String("m", "et", "e: explorer, t:throughput, l:latency, s: block size, p propagation, and any combinaison lt,ps,ltps, ...")
	//TODO cut rpc in url+port+user+rpcpass
	//TODO use interface.
	rpc := flag.String("r", "127.0.0.1:8332", "rpc link (with port) to use to fetch the data for metric lts")
	userpass := flag.String("pwd", "user:pass", "rpc user:password")
	//	password:= flag.String("r", "127.0.0.1", "rpc link to use to fetch the data for metric lts")
	// eth rpc miannet: https://mainnet.infura.io
	//output := flag.String("o", "data.json", "name of the JSON output file")
	//logFile := flag.String("f", "ethlog", "ethlog,btclog,qtumlog,eoslog")
	//verbose := flag.Bool("v", false, "Print  some quick results from the data gathered")
	bStart := flag.Int("s", 230, "block number to start the exploration")
	bEnd := flag.Int("e", 235, "block number to end the exploration")
	flag.Parse()

	BlockExplorer := explorer.EthExplorer
	switch *chain {
	case "eth":
		BlockExplorer = explorer.EthExplorer
	case "btc":
		BlockExplorer = explorer.BtcExplorer
	case "qtum":
		BlockExplorer = explorer.QtumExplorer
	case "eos":
		BlockExplorer = explorer.EosExplorer
	default:
		fmt.Println(*chain, " is not (yet) supported")
		return
	}

	if *metric != "et" && *metric != "te" {
		fmt.Println(*metric, " is not yet supported")
		return
	}
	user := strings.Split(*userpass, ":")[0]
	pass := strings.Split(*userpass, ":")[1]

	data := BlockExplorer(*rpc, user, pass, *bStart, *bEnd)
	data.Print()
	m := data.Analyse()
	m.Print()

}
