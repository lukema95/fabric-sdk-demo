package client

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"log"
)

//EnrollUser enroll a user have registerd
func (c *Client) EnrollUser(username string, password string) (bool, error) {

	_, err := c.mspClient.GetSigningIdentity(username)
	//if err == msp.ErrUserNotFound {
	if true{
		log.Println("Going to enroll user")
		err = c.mspClient.Enroll(username, msp.WithSecret(password))
		if err != nil {
			log.Printf("Failed to enroll user: %s\n", err)
			return false, err
		}
		log.Printf("Success enroll user: %s\n", username)
	} else if err != nil {
		log.Printf("Failed to get user: %s\n", err)
		return false, err
	}
	//log.Printf("User %s already enrolled, skip enrollment.\n", username)
	return true, err
}

//Register a new user with username , password and department.
func (c *Client) RegisterUser(username, password, department string) (string, error) {
	request := &msp.RegistrationRequest{
		Name:        username,
		Type:        "user",
		Affiliation: department,
		Secret:      password,
	}

	secret, err := c.mspClient.Register(request)
	if err != nil {
		fmt.Printf("register %s [%s]\n", username, err)
		return "",err
	}
	log.Printf("register %s successfully,with password %s\n", username, secret)
	return secret, nil
}

