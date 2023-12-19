# goblockchain

## 概要
- blockchainのデモを作成。
   - serverは1つでportを3つ擬似的に用意する。
   - ptopの通信はRestで擬似的に行う。

### 動作手順
1. 4つのserverを起動
- `% go run ./wallet_server/main.go ./wallet_server/wallet_server.go`
- `% go run ./blockchain_server/main.go  ./blockchain_server/blockchain_server.go`
- `% go run ./blockchain_server/main.go  ./blockchain_server/blockchain_server.go -port=5002`
- `% go run ./blockchain_server/main.go  ./blockchain_server/blockchain_server.go -port=5003`

2. moneyをsend
- `http://0.0.0.0:8082/`にアクセスして画面操作
<img width="1120" alt="スクリーンショット 2023-12-19 19 41 10" src="https://github.com/giver-yell/goblockchain/assets/68800629/73678f61-9a46-441d-b84b-6d8a1ce77097">


3. transaction poolを確認
- `http://0.0.0.0:5001/transactions`にアクセス
4. blockchainを確認
- `http://0.0.0.0:5001/chain`にアクセス
5. mining
- `http://0.0.0.0:5001/mine`にアクセス
6. transaction poolやblockchainを再確認

### 参考URL
#### Udemy 現役シリコンバレーエンジニアが教えるGoで始めるスクラッチからのブロックチェーン開発入門
- [https://www.udemy.com/course/go-blockchain/]


### Technical background of version 1 Bitcoin addresses
- [https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses]
