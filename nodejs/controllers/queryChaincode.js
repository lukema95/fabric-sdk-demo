
const Client = require('../client/client');
const config = require('config-lite')(__dirname);
module.exports = {
  query: async function(fcn, args) {
    try {
        let result;
        let client = await Client.getClient();
        let channel = client.getChannel(config.channel.channelName);
        if(!channel) {
      console.error('mychannel was not defined in the connection profile');
      throw new Error(message);
        }
        let tx_id = client.newTransactionID(true);
        
        var request = {
      targets : config.chaincode.targets,
      chaincodeId: config.chaincode.chaincodeName,
      fcn: fcn,
            args: args,
            txId: tx_id,
        };
  
        let response_payloads = await channel.queryByChaincode(request);
  
        var res = new Array();
        if(response_payloads[0].message == 'equipment does not exit' ) {
            return result = {
                code:400,
                msg: res
            };
        }
        if(response_payloads[0].toString('utf-8') == ''){
            return result = {
                code:200,
                msg: res
            };
        }
    if (response_payloads) {
      for (let i = 0; i < response_payloads.length; i++) {
                res.push(JSON.parse(response_payloads[i].toString('utf-8'))) ;
        
      }
      return result = {
                code: response_payloads[0].status,
                msg: res
            };
    } else {
            return result = {
                code: 500,
                msg: 'response_payloads is null'
            };
    }
  
    } catch (error) {
        return result = {
            code: 500,
            msg: 'Failed to evaluate transaction: ' + error,
        }
    }
  },
}

