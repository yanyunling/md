#!/bin/sh
rm -rf ./web/dist
cd ./web
npm install
npm run build
cd ../
rm -rf ./md/web
cp ./web/dist -r ./md/web
cd ./md
go build
echo md build finished
