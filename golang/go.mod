module fabric-sdk-sample

go 1.12

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta2
	github.com/maluning/fabric-sdk-sample/golang v0.0.0-20201224091207-bba746045960
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
)

replace github.com/maluning/fabric-sdk-sample/golang v0.0.0-20201224091207-bba746045960 => ./
