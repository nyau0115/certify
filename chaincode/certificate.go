package main

import (
	"encoding/json"
	"time"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/shopspring/decimal"
)

// Asset model for certificate. All asset models are kept in
// private scope i.e. they are not exported and ramin invisible to other packages
type certificateBase struct {
    Id string `json:"id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
	Timestamp time.Time `json:"timestamp"`
	Holder User `json:"holder"`
    Organisation Organisation `json:"organisation"`
    Endorsers []User `json:"endorsers"`
}

type Certificate struct {
    Id string `json:"id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
    Organisation Organisation `json:"organisation"`
    Holder User `json:"holder"`
    Endorsers []User `json:"endorsers"`
    Timestamp time.Time `json:"timestamp"`
    ExpiryDate time.Time `json:"expiryDate`
}

