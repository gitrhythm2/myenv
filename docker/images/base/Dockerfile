
FROM ubuntu:focal

ENV USER linuxbrew
ENV HOME /home/linuxbrew

SHELL ["/bin/bash", "-c"]

# 最低限必要となるパッケージをインストール
RUN apt-get update
RUN apt-get update && \
    apt-get upgrade && \
    apt-get -y install build-essential procps curl file git sudo vim fish

# ユーザー作成
RUN useradd -m ${USER} && \
    echo "${USER}:hoge" | chpasswd && \
    usermod -G adm ${USER} && \
    echo "${USER} ALL=(ALL) NOPASSWD:ALL" >>/etc/sudoers

# ユーザーとホームディレクトリを変更
USER ${USER}
WORKDIR ${HOME}
