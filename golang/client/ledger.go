package client

import (
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"log"
)

// QueryBlock queries the ledger for Block by block number.
func (c *Client)QueryBlock(blockNum uint64,options ...ledger.RequestOption) (*common.Block, error){
	resp, err := c.ledgerClient.QueryBlock(blockNum, options...)
	if err != nil{
		return nil, err
	}
	log.Println("QueryBlock successful, response: ", resp)

	return resp, nil
}

// QueryInfo queries for various useful blockchain information on this channel such as block height and current block hash.
func (c *Client) QueryInfo(options ...ledger.RequestOption) (*fab.BlockchainInfoResponse, error) {
	resp, err := c.ledgerClient.QueryInfo(options...)
	if err != nil{
		return nil, err
	}
	log.Println("QueryInfo successful, resp: ", resp)
	return resp, nil
}

// QueryBlockByTxID queries for block which contains a transaction.
func (c *Client) QueryBlockByTxID(txID fab.TransactionID, options ...ledger.RequestOption) (*common.Block, error){
	resp, err := c.ledgerClient.QueryBlockByTxID(txID, options...)
	if err != nil{
		return nil, err
	}
	log.Println("QueryBlockByTxID successful, response : ", resp)
	return resp, nil
}

// QueryTransaction queries the ledger for processed transaction by transaction ID.
func (c *Client) QueryTransaction(transactionID fab.TransactionID, options ...ledger.RequestOption) (*peer.ProcessedTransaction, error){
	resp, err := c.ledgerClient.QueryTransaction(transactionID, options...)
	if err != nil{
		return nil, err
	}
	log.Println("QueryTransaction successful, response: ", resp)
	return resp, nil
}

// QueryConfig queries for channel configuration.
func (c *Client) QueryConfig(options ...ledger.RequestOption) (fab.ChannelCfg, error){
	resp, err := c.ledgerClient.QueryConfig(options...)

	if err != nil{
		return nil, err
	}
	log.Println("QueryConfigBlock successful, response: ", resp)
	return resp, nil
}

// QueryConfigBlock returns the current configuration block for the specified channel.
func (c *Client) QueryConfigBlock(options ...ledger.RequestOption) (*common.Block, error){
	resp, err := c.ledgerClient.QueryConfigBlock(options...)
	if err != nil{
		return nil, err
	}
	log.Println("Query config block successful, response : ", resp)
	return resp, nil
}