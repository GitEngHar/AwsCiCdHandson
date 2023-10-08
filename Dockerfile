# goバージョン
FROM golang:1.21.2
# アップデートとgitのインストール
RUN apk update && apk add git
# boiler-plateディレクトリの作成
WORKDIR /app
COPY . .
RUN mkdir /go/src/github.com/boiler-plate
# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/boiler-plate
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/boiler-plate
# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports