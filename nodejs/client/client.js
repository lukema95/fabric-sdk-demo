const path = require('path');
const fabricClient = require('fabric-client');

module.exports = {
  getClient: () => new Promise((resolve) => {
    const configfile = path.join(__dirname, '../config/network-config.yaml');
    const client = fabricClient.loadFromConfig(configfile);
    client.loadFromConfig(path.join(__dirname, '../config/Org1.yaml'));

    return client.initCredentialStores().then(() => {
      resolve(client);
    });
  }),
};
