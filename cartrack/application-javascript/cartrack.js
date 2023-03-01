/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const { buildCCPOrg1, buildCCPOrg2, buildWallet, prettyJSONString} = require('../../test-application/javascript/AppUtil.js');

const myChannel = 'mychannel';
const myChaincodeName = 'cartrack';

async function trackCar(ccp,wallet,user,carId) {
	try {

		const gateway = new Gateway();

		//connect using Discovery enabled
		await gateway.connect(ccp,
			{ wallet: wallet, identity: user, discovery: { enabled: true, asLocalhost: true } });

		const network = await gateway.getNetwork(myChannel);
		const contract = network.getContract(myChaincodeName);

		let result = await contract.evaluateTransaction('trackCarById',carId);
		console.log('*** Result: Auction: ' + prettyJSONString(result.toString()));

		
		gateway.disconnect();
	} catch (error) {
		console.error(`******** FAILED to get history of CarId: ${error}`);
	}
}

async function main() {
	try {

		if (process.argv[2] === undefined || process.argv[3] === undefined ||
            process.argv[4] === undefined ) {
			console.log('Usage: node cartrack.js org userID carId');
			process.exit(1);
		}

		const org = process.argv[2];
		const user = process.argv[3];
		const carId = process.argv[4];
		

		if (org === 'Org1' || org === 'org1') {
			const ccp = buildCCPOrg1();
			const walletPath = path.join(__dirname, 'wallet/org1');
			const wallet = await buildWallet(Wallets, walletPath);
			await trackCar(ccp,wallet,user,carId);
		}
		else if (org === 'Org2' || org === 'org2') {
			const ccp = buildCCPOrg2();
			const walletPath = path.join(__dirname, 'wallet/org2');
			const wallet = await buildWallet(Wallets, walletPath);
			await trackCar(ccp,wallet,user,carId);
		}  else {
			console.log('Usage: node cartrack.js org userID carId');
			console.log('Org must be Org1 or Org2');
		}
	} catch (error) {
		console.error(`******** FAILED to run the application: ${error}`);
	}
}


main();
