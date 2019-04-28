# README

## セットアップ

GO をインストールする  
`TODO`

GOPATH を確認する  
```bash
$ go env | grep GOPATH
GOPATH="/Users/hoge/go"  # <- 環境変数に GOPATH が設定されていること
```

GOPATH 配下にソースを配置する
```bash
$ echo $GOPATH
/Users/hoge/go

$ pwd
/Users/hoge/go/src/github.com/sminoeee/sample-app  # <- こんな感じ. /Users/hoge/go/src 以降は Go の標準に従ったパス構成です
```

開発環境のビルド
```bash
$ pwd
GOPATH/src/github.com/sminoeee/sample-app  # <- GOPATH 配下の sample-app 配下であること

$ docker-compose build

...

Successfully built c93c3a35ec4b           # <- build に成功していること. c93c... は可変
Successfully tagged sample-app_web:latest
```

API サーバーを起動する
```bash
$ pwd 
GOPATH/src/github.com/sminoeee/sample-app

$ docker-compose up --force-recreate
...
sample_api | [06:52:35][SAMPLE-API] : Install completed in 6.827 s
sample_api | [06:52:35][SAMPLE-API] : Running..
sample_api | [06:52:35][SAMPLE-API] :    ____    __
sample_api | [06:52:35][SAMPLE-API] :   / __/___/ /  ___
sample_api | [06:52:35][SAMPLE-API] :  / _// __/ _ \/ _ \
sample_api | [06:52:35][SAMPLE-API] : /___/\__/_//_/\___/ v3.3.10-dev
sample_api | [06:52:35][SAMPLE-API] : High performance, minimalist Go web framework
sample_api | [06:52:35][SAMPLE-API] : https://echo.labstack.com
sample_api | [06:52:35][SAMPLE-API] : ____________________________________O/_______
sample_api | [06:52:35][SAMPLE-API] :                                     O\
sample_api | [06:52:35][SAMPLE-API] : http server started on [::]:1323

# ↑のように表示されたらOK. このログは API server (Echo) の起動ログ
```

## 開発してみる

API サーバーを起動した状態でソースコードを変更して保存すると自動ビルドが走ります

```log
sample_api | [06:58:09][SAMPLE-API] : GO changed /go/src/github.com/sminoeee/sample-app/go/usecase/seminar_usecase.go  # <- ビルド開始
sample_api | [06:58:09][SAMPLE-API] : Install started  # <- リスタート完了
```

## API を叩いてみる
```log
$ curl -X GET \
    http://localhost:1323/api/healthcheck \
    -H 'Postman-Token: 652ccaa2-239c-432c-bbd7-5e4dd81d2515' \
    -H 'cache-control: no-cache'

OK  # <- ヘルスチェックが通ること
```

## サーバーを停止する
```bash
# Ctrl+C で停止する
# 停止に成功したケース
Stopping sample_web      ... done
Stopping sample_api      ... done
Stopping sample_mysql_db ... done

# 失敗のケース
^CERROR: Aborting.

# 失敗したら↓を叩いて
$ docker-compose down
```

## Tips
- ざっくり GOPATH って何？  
 Go の作業ディレクトリ的なやつ。Ver 1.11 まではこのパスの配下にソースを配置する
