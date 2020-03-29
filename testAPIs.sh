#!/bin/bash

echo -e "\nEnrolling user 'Jim' on Org1..."
curl --request POST \
  --url http://localhost:3000/users \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data username=Jim \
  --data orgname=Org1

echo -e "\nEnrolling user 'Barry' on Org2..."
curl --request POST \
  --url http://localhost:3000/users \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data username=Barry \
  --data orgname=Org2

echo -e "\nCreating channel 'mychannel'..."
curl --request POST \
  --url http://localhost:3000/channel \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel \
  --data configpath=network/channel-artifacts/channel.tx

 echo -e "\nJoining Org1 peers to the channel..."
 curl --request POST \
  --url http://localhost:3000/join \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel

echo -e "\nJoining Org2 peers to the channel..."
  curl --request POST \
  --url http://localhost:3000/join \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data channelid=mychannel

echo -e "\nInstalling chaincode on Org1 peer..."
curl --request POST \
  --url http://localhost:3000/install \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel \
  --data path=gosdk-example/chaincode/golang \
  --data name=mycc \
  --data version=v0 \
  --data peerurl=localhost:7051

echo -e "\nInstalling chaincode on Org2 peer..."z
curl --request POST \
  --url http://localhost:3000/install \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data channelid=mychannel \
  --data path=gosdk-example/chaincode/golang \
  --data name=mycc \
  --data version=v0 \
  --data peerurl=localhost:9051

echo -e "\nInstantiating chaincode..."
curl --request POST \
  --url http://localhost:3000/instantiate \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data channelid=mychannel \
  --data path=gosdk-example/chaincode/golang \
  --data name=mycc \
  --data version=v0 \
  --data policy="AND('Org1MSP.member','Org2MSP.member')" \
  --data peer=localhost:7051 \
  --data peer=localhost:9051

echo -e "\n\nCalling 'initLedger' chaincode function..."
curl --request POST \
  --url http://localhost:3000/execute \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org1 \
  --data username=Jim \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=initLedger 

echo -e "\nCalling 'queryAllcars' chaincode function..."
curl --request POST \
  --url http://localhost:3000/query \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data username=Barry \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=queryAllCars

echo -e "\nCalling 'createCar' chaincode function..."
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

echo -e "\nCalling 'queryCar' chaincode function..."
curl --request POST \
  --url http://localhost:3000/query \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data username=Barry \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=queryCar \
  --data args=CAR12

echo -e "\nCalling 'changeCarOwner' chaincode function..."
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

echo -e "\nCalling 'queryCar' chaincode function..."
curl --request POST \
  --url http://localhost:3000/query \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data orgname=Org2 \
  --data username=Barry \
  --data channelid=mychannel \
  --data chaincodeid=mycc \
  --data function=queryCar \
  --data args=CAR12