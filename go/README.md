# Go アプリケーション実装のサンプル

---

## ソース構成説明

---

- usecase  
  サービス仕様  
  1 usecase 1 public function
  
- domain  
  - model  
    ドメインモデル
  - repository  
    interface のみ。実装は adapter/gateway で行う
  - service  
    ドメインサービス

- external  
  DB詳細、ルーティング、ロガーなどの外部サービス  
  - mysql  
    DB 詳細

- adapter  
  - controllers  
    handler
  - gateway  
    sql...
  - interfaces  
    ロガーやHTTPコンテキスト などの external の interface  
    