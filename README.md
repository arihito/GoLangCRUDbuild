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
GOPATH="/Users/hoge/go"  # <- 環境変数に GOPATH が設定されていること. 以降 /Users/hoge/go は GOPATH と表記します
```

GOPATH/bin にパスを通してリフレッシュ
```bash
# ↓ これを .bash_profile に追加（自分の環境に合わせて）
export PATH=$PATH:$GOPATH/bin

# リフレッシュ. こちらも自分の環境に合わせて
$ source ~/.bash_profile 
```

ソース一式のダウンロード
- git 設定済みの場合
```bash
$ go get -u github.com/sminoeee/sample-app
package github.com/sminoeee/sample-app: no Go files in /Users/hoge/go/src/github.com/sminoeee/sample-app
# no Go files と表示されるけど気にしない...
```

- git 未設定の場合  
```bash
# ↓のパス構成でソースコード一式を配置する. ※ Go の標準パス構成に従って配置してください.
$ pwd
GOPATH/src/github.com/sminoeee/sample-app  # <- こんな感じ. GOPATH/src 以降は Go の標準に従ったパス構成です
```

依存関係のダウンロード
```bash
$ go get -u github.com/golang/dep/cmd/dep

$ pwd
GOPATH/src/github.com/sminoeee/sample-app/go  # <- このパスに移動

$ dep ensure -v  # <- ちょっと長いけど実行完了を待つ
```

開発環境のビルド
```bash
$ pwd
GOPATH/src/github.com/sminoeee/sample-app  # <- 作業パスを確認

$ docker-compose build

...

Successfully built c93c3a35ec4b           # <- build に成功していること. c93c... は可変
Successfully tagged sample-app_web:latest
```

API サーバーを起動する (docker 上で実行)
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

## ハンズオン開始

realize をインストールする ( ホットリロードのライブラリ )
```bash
$ go get -u github.com/oxequa/realize
```

DBを起動
```bash
$ pwd
GOPATH/src/github.com/sminoeee/sample-app

$ docker-compose up --force-recreate db
```

API サーバーを起動
```bash
$ pwd
GOPATH/src/github.com/sminoeee/sample-app/go

$ realize start --name='api-server' --run
```

API サーバーを起動した状態でソースコードを変更して保存すると自動ビルドが走ります  
```log
[01:41:26][API-SERVER] : GO changed /Users/tas/go/src/github.com/sminoeee/sample-app/go/external/db/connection.go
[01:41:26][API-SERVER] : Install started
[01:41:27][API-SERVER] : Install completed in 1.523 s
[01:41:27][API-SERVER] : Running..
[01:41:28][API-SERVER] :    ____    __
[01:41:28][API-SERVER] :   / __/___/ /  ___
[01:41:28][API-SERVER] :  / _// __/ _ \/ _ \
[01:41:28][API-SERVER] : /___/\__/_//_/\___/ v3.3.10-dev
[01:41:28][API-SERVER] : High performance, minimalist Go web framework
[01:41:28][API-SERVER] : https://echo.labstack.com
[01:41:28][API-SERVER] : ____________________________________O/_______
[01:41:28][API-SERVER] :                                     O\
[01:41:28][API-SERVER] : ⇨ http server started on [::]:1323   # <- ここまで出たらOK!
```

## 作るもの
- ワイヤーフレーム  
  https://overflow.io/s/GUGA0B/