FROM gitrhythm/devbase:latest

LABEL maintainer="Hiroyuki Hosokawa" \
    version="1.0" \
    description="gitrhythm/devbaseを元に構築したnodeg開発環境のベースイメージ" \
    url="https://github.com/gitrhythm2/myenv/tree/main/docker/images/nodebase"

RUN sudo apt-get -y update && \
    sudo apt-get -y upgrade && \
    brew upgrade

SHELL ["/usr/bin/bash", "-c"]

# nodenvインストール
RUN brew install nodenv && \
    echo 'eval "$(nodenv init -)"' >> ~/.profile && \
    echo 'status --is-interactive; and source (nodenv init -|psub)' >>  ~/.config/fish/config.fish

# nodeインストール
RUN eval "$(nodenv init -)" && \
    curl -fsSL https://github.com/nodenv/nodenv-installer/raw/master/bin/nodenv-doctor | bash && \
    nodenv install $(nodenv install -l | egrep "^[0-9]+.[0-9]+.[0-9]+\$" | tail -n 1) && \
    nodenv global $(nodenv versions | tr -d '  ')
