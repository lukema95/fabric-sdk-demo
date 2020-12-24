package client

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"log"
)

// RegisterBlockEvent registers for block events. If the caller does not have permission
func (c * Client) RegisterBlockEvent() (fab.Registration, error){
	registration, _, err := c.eventClient.RegisterBlockEvent()
	if err != nil{
		return nil, err
	}
	log.Println("Register block event successful, registration : ", registration)

	return registration, nil
}

// RegisterChaincodeEvent registers for chaincode events. Unregister must be called when the registration is no longer needed.
func (c *Client)RegisterChaincodeEvent(ccID, eventFilter string) (fab.Registration, error){
	registration, _, err := c.eventClient.RegisterChaincodeEvent(ccID,eventFilter)
	if err != nil{
		return nil, err
	}
	log.Println("RegisterChaincodeEvent, registration : ", registration)
	return registration, nil

}

// RegisterTxStatusEvent registers for transaction status events. Unregister must be called when the registration is no longer needed.
func (c *Client)RegisterTxStatusEvent(txID string) (fab.Registration, error)   {
	registration, _, err := c.eventClient.RegisterTxStatusEvent(txID)
	if err != nil{
		return nil, err
	}
	log.Println("RegisterTxStatusEvent successful, registration : ", registration)

	return registration, nil
}

// Unregister removes the given registration and closes the event channel.
func (c *Client) UnRegister(reg fab.Registration){
	c.eventClient.Unregister(reg)
	log.Println("UnRegister txStatus successful")
}

