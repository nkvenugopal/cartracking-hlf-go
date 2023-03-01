
 smart contract that tracks a car from a manufacturer to an owner.


When the car is manufactured by the manufacturer it is CREATED state. After it is delivered to a dealer it will be in READY_FOR_SALE state. Once it is sold to a customer it will be in SOLD state.

Car can only be manufactured by a manufacturer. Can only be sold by a dealer.


## Deploy the chaincode

We will run the cartrack smart contract using the Fabric test network. Open a command terminal and navigate to the test network directory:
```
cd fabric-samples/test-network
```

You can then run the following command to deploy the test network.
```
./network.sh up createChannel -ca
```

Note that we use the `-ca` flag to deploy the network using certificate authorities. We will use the CA to register and enroll our users.


```
./network.sh deployCC -ccn cartrack -ccp ../cartrack/chaincode-go/ -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"
```

## Install the application dependencies

We will interact with the cartrack smart contract through a set of Node.js applications. Change into the `application-javascript` directory:
```
cd fabric-samples/cartrack/application-javascript
```

From this directory, run the following command to download the application dependencies:
```
npm install
```

## Register and enroll the application identities

To interact with the network, you will need to enroll the Certificate Authority administrators of Org1 and Org2. You can use the `enrollAdmin.js` program for this task. Run the following command to enroll the Org1 admin:
```
node enrollAdmin.js org1
```
You should see the logs of the admin wallet being created on your local file system. Now run the command to enroll the CA admin of Org2:
```
node enrollAdmin.js org2
```


TBD

## Clean up

When your are done using the cartrack smart contract, you can bring down the network and clean up the environment. In the `cartrack/application-javascript` directory, run the following command to remove the wallets used to run the applications:
```
rm -rf wallet
```

You can then navigate to the test network directory and bring down the network:
````
cd ../../test-network/
./network.sh down
````
