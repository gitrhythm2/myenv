
FROM gitrhythm/base:latest

LABEL maintainer="Hiroyuki Hosokawa" \
      version="1.0" \
      description="homebrewを元にした開発環境のベースイメージ" \
      url="https://github.com/gitrhythm2/myenv/tree/main/docker/images/devbase"

ENV USER linuxbrew
ENV HOME /home/linuxbrew
ENV PATH /home/${USER}/.linuxbrew/bin:$PATH

SHELL ["/bin/bash", "-c"]

# ユーザー作成
RUN useradd -m ${USER} && \
    echo "${USER}:hoge" | chpasswd && \
    usermod -G adm ${USER} && \
    echo "${USER} ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

# ユーザーとホームディレクトリを変更
USER ${USER}
WORKDIR ${HOME}

# HomeBrew環境構築
RUN git clone https://github.com/Homebrew/brew /home/${USER}/.linuxbrew/Homebrew && \
    mkdir /home/${USER}/.linuxbrew/bin && \
    ln -s /home/${USER}/.linuxbrew/Homebrew/bin/brew /home/${USER}/.linuxbrew/bin && \
    echo "eval $(/home/${USER}/.linuxbrew/bin/brew shellenv)" >>~/.profile && \
    brew doctor

# 開発環境のベースをインストール
RUN brew install git && \
    brew install fish && \
    brew install anyenv

# bashとfishの設定ファイル関連
RUN mkdir -p .config/fish
COPY --chown=linuxbrew:linuxbrew config/ .
RUN echo 'eval "$(anyenv init -)"' >> ~/.profile && \
    echo 'set -x PATH $HOME/.anyenv/bin $PATH' >> ~/.config/fish/config.fish && \
    echo 'eval (anyenv init - | source)' >> ~/.config/fish/config.fish && \
    # プロンプトの色をホストOSのターミナルと異なる色にする
    echo 'PS1="\[\e[1;35m\]${debian_chroot:+($debian_chroot)}\u@\h:\w\$ "' >> .bashrc
