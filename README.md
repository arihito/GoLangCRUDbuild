# README

## セットアップ

GO をインストールする  
windows, mac と環境ごとに異なるため [golang install] みたいな感じでググってインストールする  
mac は homebrew を使うと楽  

```bash
$ brew install go
```

GOPATH を確認する  
GOPATH は Go の作業ディレクトリ的なやつ。Ver 1.11 まではこのパスの配下にソースを配置する  

```bash
$ go env | grep GOPATH
GOPATH="/Users/hoge/go"  # <- 環境変数に GOPATH が設定されていること
```

GOPATH 配下にソースコード一式を配置する. Go の標準パス構成に従って配置してください.
```bash
$ echo $GOPATH
/Users/hoge/go  # <- 環境変数 GOPATH が設定されていること. 以降 /Users/hoge/go は GOPATH と表記します

$ pwd
GOPATH/src/github.com/sminoeee/sample-app  # <- こんな感じ. GOPATH/src 以降は Go の標準に従ったパス構成です
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
GOPATH/src/github.com/sminoeee/sample-app  # <- パスを確認

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

API サーバーの起動確認: health check を叩いてみる
```log
$ curl -X GET http://localhost:1323/api/healthcheck
OK  # <- ヘルスチェックが通ること
```

API サーバーを停止する
```bash
# Ctrl+C で停止する
# 停止に成功したケース
Stopping sample_api      ... done
Stopping sample_mysql_db ... done

# 失敗のケース
^CERROR: Aborting.

# 失敗したら↓を叩いて
$ docker-compose down
```

## 開発してみる

`TODO:: realize`

API サーバーを起動した状態でソースコードを変更して保存すると自動ビルドが走ります  

```log
sample_api | [06:58:09][SAMPLE-API] : GO changed /go/src/github.com/sminoeee/sample-app/go/usecase/seminar_usecase.go  # <- ビルド開始
sample_api | [06:58:09][SAMPLE-API] : Install started  # <- リスタート完了
```

## 作るもの
- ワイヤーフレーム  
  https://overflow.io/s/GUGA0B/