
FROM gitrhythm/base:latest

ENV USER linuxbrew
ENV HOME /home/linuxbrew

SHELL ["/bin/bash", "-c"]

# ユーザー作成
RUN useradd -m ${USER} && \
    echo "${USER}:hoge" | chpasswd && \
    usermod -G adm ${USER} && \
    echo "${USER} ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

# ユーザーとホームディレクトリを変更
USER ${USER}
WORKDIR ${HOME}

# bash関連
RUN echo 'PS1="\[\e[1;35m\]${debian_chroot:+($debian_chroot)}\u@\h:\w\$ "' >> .bashrc
COPY --chown=linuxbrew:linuxbrew ./config/.bash_aliases .

# fish関連
RUN mkdir -p .config/fish
COPY --chown=linuxbrew:linuxbrew ./config/ .
