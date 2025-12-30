package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const rpcURL = "https://api.mainnet-beta.solana.com"

type RPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

func callRPC(method string, params interface{}, out interface{}) {
	req := RPCRequest{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}

	b, _ := json.Marshal(req)
	resp, _ := http.Post(rpcURL, "application/json", bytes.NewReader(b))
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(out)
}

func main() {
	// 1️⃣ Get latest slot
	var slotResp struct {
		Result int `json:"result"`
	}
	callRPC("getSlot", nil, &slotResp)

	fmt.Println("Latest slot:", slotResp.Result)

	// 2️⃣ Fetch block for that slot
	var blockResp struct {
		Result struct {
			Transactions []interface{} `json:"transactions"`
		} `json:"result"`
	}

	callRPC("getBlock", []interface{}{slotResp.Result}, &blockResp)

	fmt.Println("Transactions count:", len(blockResp.Result.Transactions))
}
