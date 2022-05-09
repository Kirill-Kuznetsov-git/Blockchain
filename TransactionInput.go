package main

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature string
	//Signature []byte
	//PubKey    []byte
}
