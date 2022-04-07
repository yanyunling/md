#!/bin/sh
cd ./web
npm install
npm run build
cd ../
rm -rf ./md/web
cp ./web/dist ./md/web
cd ./md
go build
echo md build finished
