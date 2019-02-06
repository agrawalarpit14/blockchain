'use strict';

var Fabric_Client = require('fabric-client');
var Fabric_CA_Client = require('fabric-ca-client');

var path = require('path');

var fabric_client = new Fabric_Client();
var fabric_ca_client = null;
var admin_user = null;
var member_user = null;
var store_path = path.join(__dirname, 'hfc-key-store');
console.log(' Store path:'+store_path);


Fabric_Client.newDefaultKeyValueStore({ path: store_path
}).then((state_store) => {
    fabric_client.setStateStore(state_store);
    var crypto_suite = Fabric_Client.newCryptoSuite();
    var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
    crypto_suite.setCryptoKeyStore(crypto_store);
    fabric_client.setCryptoSuite(crypto_suite);
    fabric_ca_client = new Fabric_CA_Client('https://admin:fc25967498@n0f406e527ffb41b2bb76b0439815db5a-org1-ca.uk05.blockchain.ibm.com:31011', null ,'org1CA', crypto_suite);
    return fabric_client.getUserContext('admin', true);
}).then((user_from_store) => {
    if (user_from_store && user_from_store.isEnrolled()) {
        console.log('Successfully loaded admin from persistence');
        admin_user = user_from_store;
    } else {
        throw new Error('Failed to get admin.... run enrollAdmin.js');
    }

    return fabric_ca_client.register({enrollmentID: 'ajay', affiliation: 'org1',role: 'client'}, admin_user);
}).then((secret) => {
    console.log('Successfully registered user - secret:'+ secret);

    return fabric_ca_client.enroll({enrollmentID: 'ajay', enrollmentSecret: secret});
}).then((enrollment) => {
  console.log('Successfully enrolled member user "user" ');
  return fabric_client.createUser(
     {username: 'ajay',
     mspid: 'org1',
     cryptoContent: { privateKeyPEM: enrollment.key.toBytes(), signedCertPEM: enrollment.certificate }
     });
}).then((user) => {
     member_user = user;

     return fabric_client.setUserContext(member_user);
}).then(()=>{
     console.log('User was successfully registered and enrolled and is ready to interact with the fabric network');

}).catch((err) => {
    console.error('Failed to register: ' + err);
	if(err.toString().indexOf('Authorization') > -1) {
		console.error('Authorization failures may be caused by having admin credentials from a previous CA instance.\n' +
		'Try again after deleting the contents of the store directory '+store_path);
	}
});
