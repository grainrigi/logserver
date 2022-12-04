# 開発用コンテナ

FROM golang:1.19-bullseye

# nodejsのインストール
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
  && apt-get install -y nodejs \
  && npm -g i yarn

# golang-migrateのインストール
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash \
 && apt -y install migrate

RUN apt -y update \
 && apt -y install vim sudo postgresql-client docker docker-compose moreutils

# Go言語ツールのインストール
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
  && go install github.com/ramya-rao-a/go-outline@latest \
  && go install github.com/nsf/gocode@latest \
  && go install github.com/acroca/go-symbols@latest \
  && go install github.com/fatih/gomodifytags@latest \
  && go install github.com/josharian/impl@latest \
  && go install github.com/haya14busa/goplay/cmd/goplay@latest \
  && go install github.com/go-delve/delve/cmd/dlv@latest \
  && go install golang.org/x/lint/golint@latest \
  && go install golang.org/x/tools/gopls@latest

# 一般ユーザーの作成
ARG USERNAME=root
ARG UID
ARG GID

# ホストと同一のUID・GID・ユーザー名でユーザー作成
RUN [ -n "$UID" ] && (groupadd --gid $GID $USERNAME \
 && useradd --uid $UID --gid $GID -m $USERNAME \
 && echo "$USERNAME ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USERNAME \
 && chmod 0440 /etc/sudoers.d/$USERNAME \
 && chown -R $UID:$GID /go)

# .bashrcの準備
COPY .bashrc /tmp/.bashrc
RUN [ -n "$UID" ] && (cat /tmp/.bashrc >> /home/$USERNAME/.bashrc) \
 || (cp /etc/skel/. /root/ && cat /tmp/.bashrc >> /root/.bashrc)

USER $USERNAME
WORKDIR /workspace
