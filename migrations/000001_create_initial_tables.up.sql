CREATE TABLE IF NOT EXISTS users (
    id serial primary key,

    name varchar not null,
    privilege int not null default 0
);

CREATE TABLE IF NOT EXISTS contests (
    id serial primary key,

    name varchar not null,  --コンテスト名
    start_time timestamptz, --コンテスト開始時刻
    end_time timestamptz,   --コンテスト終了時刻
    type int not null,      --コンテスト種別(1:SingleOp, 2:MultiOp)
    cfg varchar not null,   --zLog CFG
    call varchar not null   --自局コールサイン
);

CREATE TABLE IF NOT EXISTS operators (
    id serial primary key,

    name varchar not null, --名前orコールサイン
    license int not null   --無線従事者免許の級(アマチュア無線技士相当)
);

CREATE TABLE IF NOT EXISTS logs (
    id serial primary key,
    contest int not null,

    time timestamptz not null, --ログ時刻
    call varchar not null,     --受信コールサイン
    rst varchar not null,      --受信RST
    rcvd varchar not null,     --受信コンテストナンバー
    band varchar not null,     --周波数帯
    mode varchar not null,     --通信モード
    pwr varchar,               --送信パワー(P,L,M,H)
    op int,                    --オペレータ
    note varchar not null,     --備考

    txrst varchar, --送信RST
    txd varchar,   --送信コンテストナンバー

    foreign key (contest) references contests (id) on delete cascade,
    foreign key (op) references operators (id)
);