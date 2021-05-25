FROM gitrhythm/devbase:latest

LABEL maintainer="Hiroyuki Hosokawa" \
    version="1.0" \
    description="gitrhythm/devbaseを元に構築したgolang開発環境のベースイメージ" \
    url="https://github.com/gitrhythm2/myenv/tree/main/docker/images/gobase"

RUN sudo apt-get -y update && \
    sudo apt-get -y upgrade && \
    brew upgrade

ENV GOENV_ROOT=$HOME/.goenv
ENV PATH $GOENV_ROOT/bin:$PATH

RUN git clone https://github.com/syndbg/goenv.git ${GOENV_ROOT}

RUN echo 'eval "$(goenv init -)"' >> ~/.profile && \
    echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.profile && \
    # fish用
    echo 'eval (goenv init - | source)' >> ~/.config/fish/config.fish && \
    echo 'set -x PATH $GOPATH/bin $PATH' >> ~/.config/fish/config.fish

# 最新のgoをインストール
RUN goenv install "$(goenv install -l | egrep "^  [0-9]+.[0-9]+.[0-9]+\$" | tail -n 1 | tr -d '  ')" && \
    goenv global $(goenv versions | tr -d '  ')

RUN eval "$(goenv init -)" && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/motemen/gore/cmd/gore@latest

# vscodeのgo extension
RUN eval "$(goenv init -)" && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@master && \
    go install honnef.co/go/tools/cmd/staticcheck@latest
