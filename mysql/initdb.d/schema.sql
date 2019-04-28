/*
 FIXME: timezone to JST
*/
USE sample_db;

/*  users: ユーザー */
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(256) NOT NULL,
    display_name VARCHAR(256) DEFAULT NULL,
    created_at   DATETIME     DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

/*  study_groups: グループ */
DROP TABLE IF EXISTS study_groups;
CREATE TABLE study_groups
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(2000)   NOT NULL comment 'グループ名',
    page_url   VARCHAR(2000)   NOT NULL comment 'グループページURL',
    user_id    BIGINT UNSIGNED NOT NULL comment '主催者',
    published  BOOLEAN  DEFAULT FALSE comment '公開済みかどうか',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id)
);

/* study_group_members: グループメンバー */
DROP TABLE IF EXISTS study_group_members;
CREATE TABLE study_group_members
(
    id             SERIAL PRIMARY KEY,
    study_group_id BIGINT UNSIGNED  NOT NULL comment 'グループ',
    user_id        BIGINT UNSIGNED  NOT NULL comment 'ユーザー',
    member_kind    TINYINT UNSIGNED NOT NULL comment 'メンバー種別',
    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (study_group_id, user_id),
    INDEX user_id (user_id),
    FOREIGN KEY (study_group_id) REFERENCES study_groups (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

/*  events: イベント */
DROP TABLE IF EXISTS events;
CREATE TABLE events
(
    id             SERIAL PRIMARY KEY,
    title          VARCHAR(256)    NOT NULL comment 'タイトル',
    sub_title      VARCHAR(256) DEFAULT NULL comment 'サブタイトル',
    image_path     VARCHAR(256) DEFAULT NULL comment '画像パス',
    study_group_id BIGINT UNSIGNED NOT NULL comment '所属グループ',
    event_start    DATETIME        NOT NULL comment '開催開始',
    event_end      DATETIME        NOT NULL comment '開催終了',
    apply_start    DATETIME        NOT NULL comment '募集開始',
    apply_end      DATETIME        NOT NULL comment '募集終了',
    summary        VARCHAR(2000)   NOT NULL comment 'イベント説明',
    user_id        BIGINT UNSIGNED NOT NULL comment '主催者',
    published      BOOLEAN      DEFAULT FALSE comment '公開済みかどうか',
    created_at     DATETIME     DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (study_group_id) REFERENCES study_groups (id)
);

/*  tags: タグ */
DROP TABLE IF EXISTS tags;
CREATE TABLE tags
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(128) NOT NULL comment 'タグ名',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (name)
);

/*  event_tags: イベントタグ */
DROP TABLE IF EXISTS event_tags;
CREATE TABLE event_tags
(
    id         SERIAL PRIMARY KEY,
    event_id   BIGINT UNSIGNED NOT NULL comment 'イベント',
    tag_id     BIGINT UNSIGNED NOT NULL comment 'タグ',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (event_id, tag_id),
    FOREIGN KEY (event_id) REFERENCES events (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id)
);

/*  event_seats: イベント参加枠 */
DROP TABLE IF EXISTS event_seats;
CREATE TABLE event_seats
(
    id              SERIAL PRIMARY KEY,
    event_id        BIGINT UNSIGNED             NOT NULL,
    title           VARCHAR(256)                NOT NULL comment '参加枠名',
    maximum_limit   SMALLINT UNSIGNED           NOT NULL comment '定員数',
    reservation_num SMALLINT UNSIGNED DEFAULT 0 NOT NULL comment '予約数',
    fee             INT               DEFAULT 0 comment '参加費',
    raffle_method   TINYINT UNSIGNED            NOT NULL comment '抽選方法',
    created_at      DATETIME          DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME          DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (event_id, title),
    FOREIGN KEY (event_id) REFERENCES events (id)
);

/* event_reserves: イベント予約 */
DROP TABLE IF EXISTS event_reserves;
CREATE TABLE event_reserves
(
    id            SERIAL PRIMARY KEY,
    user_id       BIGINT UNSIGNED NOT NULL comment '予約者',
    evnet_id      BIGINT UNSIGNED NOT NULL comment 'イベント',
    event_seat_id BIGINT UNSIGNED NOT NULL comment 'イベント参加枠',
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (user_id, evnet_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (evnet_id) REFERENCES events (id),
    FOREIGN KEY (event_seat_id) REFERENCES event_seats (id)
);
