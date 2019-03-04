# grpc-test

## 概要
gRPCを試してみた

## Frontend/Backend
- frontendはtypescript
- backendはgo

## 自動生成
- front/backendそれぞれgrpcで型を自動生成
- tsのserverはfrontend側のmockとして利用
- goのclientはmockgenで生成してmockとして利用

## 実行スクリプト
- frontendはpackage.jsonを参照
- backendはMakefileを参照
