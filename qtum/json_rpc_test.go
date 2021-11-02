package qtum

import (
	"fmt"
	"testing"
)

func TestQtumRPC(t *testing.T) {
	rpc := QtumRPC{
		Host:    "http://127.0.0.1:13889",
		ID:      "test",
		Version: "1.0",
		User:    "user",
		Pass:    "pass",
	}

	r, err := rpc.Call("getblockchaininfo", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("#1 blockchaininfo: %+v\n", r)

	r, err = rpc.Call("getblockhash", []interface{}{600})
	if err != nil {
		panic(err)
	}
	fmt.Printf("#2 addresses list: %+v\n", r)

	r, err = rpc.GetBlockHash(600)
	if err != nil {
		panic(err)
	}
	fmt.Printf("#3 address: %+v\n", r)

	r, err = rpc.GetBlock(1460)
	if err != nil {
		panic(err)
	}
	fmt.Printf("#4 address: %+v\n", r)
	b := r.(Block)
	fmt.Println(b.GetTimestamp())
	fmt.Println(b.GetBlockHash())
	fmt.Println(b.GetNumberTx())
	fmt.Println(b.GetDate())
	fmt.Println(b.GetTxsList())

	/*
		r, err = rpc.Call("getaccountaddress", []interface{}{"product.1"})
		if err != nil {
			panic(err)
		}
		fmt.Printf("#3 address: %+v\n", r)

		r, err = rpc.Call("sendtoaddress", []interface{}{"qcqCSPnAMJBMszPSu2puFNPtMivaYQG6Ep", 21.12345678})
		if err != nil {
			if err.Error() == "[code: -6] Insufficient funds" {
				fmt.Println("#4 Insufficient funds")
			} else {
				panic(err)
			}
		} else {
			fmt.Printf("#4 txid: %+v\n", r)
		}

		r, err = rpc.Call("getreceivedbyaddress", []interface{}{"qcqCSPnAMJBMszPSu2puFNPtMivaYQG6Ep", 1})
		if err != nil {
			panic(err)
		}
		fmt.Printf("#5 received: %+v\n", r)

		r, err = rpc.Call("getblockcount", nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("#6 blockcount: %+v\n", r)
	*/
}
