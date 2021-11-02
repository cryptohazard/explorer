package qtum

import "time"

type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Strippedsize      int      `json:"strippedsize"`
	Size              int      `json:"size"`
	Weight            int      `json:"weight"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	VersionHex        string   `json:"versionHex"`
	Merkleroot        string   `json:"merkleroot"`
	HashStateRoot     string   `json:"hashStateRoot"`
	HashUTXORoot      string   `json:"hashUTXORoot"`
	Tx                []string `json:"tx"`
	Time              int      `json:"time"`
	Mediantime        int      `json:"mediantime"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	NTx               int      `json:"nTx"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
	Flags             string   `json:"flags"`
	Proofhash         string   `json:"proofhash"`
	Modifier          string   `json:"modifier"`
}

//GetBlockNumber is not needed actually :-)
func (b Block) GetBlockNumber() int {
	return b.Height
}
func (b Block) GetNumberTx() uint {
	return uint(b.NTx)
}

func (b Block) GetTimestamp() int64 {
	return int64(b.Time)
}

func (b Block) GetBlockHash() string {
	return b.Hash
}

func (b Block) GetDate() string {
	t := time.Unix(b.GetTimestamp(), 0)
	return t.Format("2006-01-02 15:04:05.000")

}

func (b Block) GetTxsList() []string {
	return b.Tx
}
