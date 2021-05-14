
FROM ubuntu:focal

LABEL maintainer="Hiroyuki Hosokawa" \
      version="1.0" \
      description="ubuntu:focalを元にした、開発環境のベースとなるイメージ" \
      url="https://github.com/gitrhythm2/myenv/tree/main/docker/images/base"

SHELL ["/bin/bash", "-c"]

# 最低限必要となるパッケージをインストール
RUN apt-get update && \
    apt-get upgrade && \
    apt-get -y install build-essential procps curl file git sudo vim
