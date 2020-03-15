#!/bin/bash

echo -e "\nEnroll user on Org1..."
curl --request POST \
  --url http://localhost:3000/users \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data username=Jim \
  --data orgname=Org1

echo -e "\nEnroll user on Org2..."
curl --request POST \
  --url http://localhost:3000/users \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data username=Barry \
  --data orgname=Org2

echo -e "\nCreate channel..."
curl --request POST \
  --url http://localhost:3000/channel \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel \
  --data configpath=network/channel-artifacts/channel.tx

 echo -e "\nJoin Org1 peers to channel..."
 curl --request POST \
  --url http://localhost:3000/join \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel

echo -e "\nJoin Org2 peers to channel..."
  curl --request POST \
  --url http://localhost:3000/join \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data channelid=mychannel

echo -e "\nInstall chaincode on Org1 peer..."
curl --request POST \
  --url http://localhost:3000/install \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel \
  --data path=gosdk-example/chaincode/golang \
  --data name=mycc \
  --data version=v0 \
  --data peerurl=localhost:7051

echo -e "\nInstall chaincode on Org2 peer..."
curl --request POST \
  --url http://localhost:3000/install \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data channelid=mychannel \
  --data path=gosdk-example/chaincode/golang \
  --data name=mycc \
  --data version=v0 \
  --data peerurl=localhost:9051

echo -e "\nInstantiate chaincode..."
curl --request POST \
  --url http://localhost:3000/instantiate \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel \
  --data path=gosdk-example/chaincode/golang \
  --data name=mycc \
  --data version=v0 \
  --data peer=localhost:7051 \
  --data peer=localhost:9051

echo -e "\nCall initLedger chaincode function..."
curl --request POST \
  --url http://localhost:3000/execute \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data username=Jim \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=initLedger 

echo -e "\nCall queryAllcars chaincode function..."
curl --request POST \
  --url http://localhost:3000/query \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data username=Barry \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=queryAllCars

echo -e "\nCall createCar chaincode function..."
curl --request POST \
  --url http://localhost:3000/execute \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data username=Jim \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=createCar \
  --data args=CAR12 \
  --data args=Honda \
  --data args=Accord \
  --data args=Black \
  --data args=Tom

echo -e "\n Call queryCar chaincode function..."
curl --request POST \
  --url http://localhost:3000/query \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data username=Barry \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=queryCar \
  --data args=CAR12

echo -e "\nCall changeCarOwner chaincode function..."
curl --request POST \
  --url http://localhost:3000/execute \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data username=Jim \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=changeCarOwner \
  --data args=CAR12 \
  --data args=Dave

echo -e "\n Call queryCar chaincode function..."
curl --request POST \
  --url http://localhost:3000/query \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data username=Barry \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=queryCar \
  --data args=CAR12