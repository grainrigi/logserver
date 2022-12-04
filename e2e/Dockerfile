FROM golang:1.19-bullseye

# nodejsのインストール
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
  && apt-get install -y nodejs \
  && npm -g i yarn

# golang-migrateのインストール
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash \
 && apt -y install migrate

# psql, fuser, chronicのインストール
RUN apt -y update \
 && apt -y install postgresql-client psmisc moreutils

WORKDIR /workspace/e2e
