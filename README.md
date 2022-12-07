# 開発環境構築

まずはDocker Desktop(Windows, MacOS)またはdockerd(Linux)をインストールしてください。

Linuxでdockerdを使う場合はdocker-compose(またはdocker compose)もインストールしてください。

## VSCodeを使う場合(おすすめ)

まず、拡張機能「Dev Containers」をインストールします。

次に、このリポジトリのディレクトリをVSCodeで再度開き直します。

開き直したら右下にTipsが出てきて「Reopen in Container」というボタンが押せるはずなので押します。
(もし逃してしまった場合はコマンドパレットを開いて「Dev Container: Reopen in Container」を実行してください。)

しばらくするとDev Containerが立ち上がります。

※自動でターミナルが開いてくれないので、ターミナル右上の「+」ボタンから自分でターミナルを開いてください

## VSCodeを使わない場合

Windowsの場合は`startdev.cmd`を、Linux/MacOSの場合は`startdev.sh`を実行します。

いずれの場合もDev Container内のシェルが立ち上がります。

# Dev Container内での操作

## サーバーのコンパイル&実行

```sh
$ go run .
```

## DBのマイグレーション

[golang-migrate](https://github.com/golang-migrate/migrate)の`migrate`コマンドを使って操作できます。
開発用DBのマイグレーションをしたい場合は`migrate-dev`というエイリアスを用意してあるので、
そちらを使えばいちいちデータベースURLやmigrationディレクトリを指定しなくても使えます。

### テーブルの構築

```sh
migrate-dev up
```

### テーブルの削除

※ すべてのデータが削除されます！

```sh
migrate-dev down
```

### マイグレーションの追加

```sh
migrate create -ext sql -dir migrations -seq create_xxx_table
```

## E2Eテスト

HTTPリクエスト〜DBアクセスまで一気通貫でテストできます。

※Dev Container内のみから実行可能

```sh
/workspace$ ./e2e/run.sh
Preparing test environment...
yarn run v1.22.19
$ jest --runInBand
 PASS  tests/operators.test.ts
  operators
    ✓ GET (14 ms)

Test Suites: 1 passed, 1 total
Tests:       1 passed, 1 total
Snapshots:   0 total
Time:        0.663 s
Ran all test suites.
Done in 1.57s.
```