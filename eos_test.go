package explorer

import (
	"fmt"
	"testing"
)

func TestEosExplorer(t *testing.T) {
	b := EosExplorer("https://api.eosio.cr", "", "", 1000, 1010)

	fmt.Println("Connected Successfully to Eos rpc!")
	fmt.Println(b)

}
