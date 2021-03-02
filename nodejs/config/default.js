/**
 * 默认环境配置
 */
const path = require('path');

module.exports = {
  env: 'dev',
  chaincode: { // 链码相关
    chaincodePath: path.join('chaincode/'),
    chaincodeName: 'mycc',
    chaincodeVersion: 'v1.0',
    chaincodeType: 'golang',
    chaincodePackage: '',
    targets: ['peer0.org1.example.com'],
  },
  channel: { // 通道相关
    channelName: 'mychannel',
  },
};
