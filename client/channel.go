package client

import (
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"log"
)

// SaveChannel creates or updates channel.
func (c *Client) SaveChannel(channelId string, channelConfigPath string) (string, error) {
	signingIdentity, err := c.mspClient.GetSigningIdentity(c.OrgAdmin)
	if err != nil{
		log.Panic("failed to create signingIdentity")
	}

	//r, _ := os.Open("/Users/jjvincent/work/go/src/github.com/hyperledger/fabric-sdk-go-sample-gm/config/businesschannel.tx")
	req := resmgmt.SaveChannelRequest{
		ChannelID: channelId,
		//ChannelConfig: r,
		ChannelConfigPath: channelConfigPath,
		SigningIdentities: []msp.SigningIdentity{signingIdentity},
	}

	//resps, err := c.resMgtClient.SaveChannel(req)
	resps, err := c.resMgtClient.SaveChannel(req,resmgmt.WithOrdererEndpoint("grpcs://localhost:7050"))

	if err != nil {
		return "",err
	}
	log.Println("Save channel successful, response txId : ", resps.TransactionID)

	return string(resps.TransactionID), nil
}

// JoinChannel allows for peers to join existing channel with optional custom options (specific peers, filtered peers).
func (c *Client) JoinChannel(channelID string, options ...resmgmt.RequestOption) error {

	err := c.resMgtClient.JoinChannel(channelID, options...)
	if err != nil{
		return err
	}
	log.Println("Join channel successful")
	return nil
}

// QueryChannels queries the names of all the channels that a peer has joined.
func (c *Client) QueryChannels(options ...resmgmt.RequestOption) (*peer.ChannelQueryResponse, error) {
	resp, err:= c.resMgtClient.QueryChannels(options...)
	if err != nil{
		return resp, err
	}
	log.Println("Query channel successful, response : ", resp)
	return nil, err
}
