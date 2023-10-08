# https://hub.docker.com/_/golang からイメージを選ぶ
FROM golang:1.21.2-alpine3.18
# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app/
