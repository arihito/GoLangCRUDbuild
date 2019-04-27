USE sample_db;

/*  users: ユーザー */
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(256) NOT NULL,
    display_name VARCHAR(256) DEFAULT NULL,
    canceled_at  DATETIME     DEFAULT NULL,
    created_at   DATETIME     DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


/*  seminars: セミナー */
DROP TABLE IF EXISTS seminars;
CREATE TABLE seminars
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(256)    NOT NULL,
    sub_title  VARCHAR(256) DEFAULT NULL,
    event_date DATE            NOT NULL,
    summary    VARCHAR(2000)   NOT NULL,
    organizer  BIGINT UNSIGNED NOT NULL,
    start      DATETIME        NOT NULL,
    end        DATETIME        NOT NULL,
    location   VARCHAR(1000)   NOT NULL,
    created_at DATETIME     DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (organizer) REFERENCES users (id)
);

/*  seminar_speakers: セミナー発表者 */
DROP TABLE IF EXISTS seminar_speakers;
CREATE TABLE seminar_speakers
(
    id         SERIAL PRIMARY KEY,
    seminar_id BIGINT UNSIGNED NOT NULL,
    speak_by   BIGINT UNSIGNED NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (seminar_id) REFERENCES seminars (id),
    FOREIGN KEY (speak_by) REFERENCES users (id)
);

/*  seminar_seats: セミナー座席 */
DROP TABLE IF EXISTS seminar_seats;
CREATE TABLE seminar_seats
(
    id              SERIAL PRIMARY KEY,
    seminar_id      BIGINT UNSIGNED             NOT NULL,
    summary         VARCHAR(256)                NOT NULL,
    raffle_date     DATETIME                    NOT NULL,
    maximum_limit   SMALLINT UNSIGNED           NOT NULL,
    reservation_num SMALLINT UNSIGNED DEFAULT 0 NOT NULL,
    created_at      DATETIME          DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME          DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (seminar_id, summary),
    FOREIGN KEY (seminar_id) REFERENCES seminars (id)
);


/*  seminar_tags: セミナータグ */
DROP TABLE IF EXISTS seminar_tags;
CREATE TABLE seminar_tags
(
    id         SERIAL PRIMARY KEY,
    seminar_id BIGINT UNSIGNED NOT NULL,
    tag        VARCHAR(128)    NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (seminar_id, tag),
    FOREIGN KEY (seminar_id) REFERENCES seminars (id)
);

/*
    tags: タグ
    todo: そのうち作るかも
*/

/*
    reserves: 予約
    ユーザーのセミナー予約
*/
DROP TABLE IF EXISTS reserves;
CREATE TABLE reserves
(
    id              SERIAL PRIMARY KEY,
    reserved_by     BIGINT UNSIGNED NOT NULL,
    seminar_seat_id BIGINT UNSIGNED NOT NULL,
    seminar_id      BIGINT UNSIGNED NOT NULL,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (reserved_by, seminar_id),
    FOREIGN KEY (reserved_by) REFERENCES users (id),
    FOREIGN KEY (seminar_id) REFERENCES seminars (id)
);


/*  study_groups: 勉強会 */
DROP TABLE IF EXISTS study_groups;
CREATE TABLE study_groups
(
    id         SERIAL PRIMARY KEY,
    summary    VARCHAR(2000)   NOT NULL,
    organizer  BIGINT UNSIGNED NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (organizer) REFERENCES users (id)
);

/*
    study_group_members: 勉強会メンバー
    勉強会に所属するメンバー
*/
DROP TABLE IF EXISTS study_group_members;
CREATE TABLE study_group_members
(
    id             SERIAL PRIMARY KEY,
    study_group_id BIGINT UNSIGNED  NOT NULL,
    user_id        BIGINT UNSIGNED  NOT NULL,
    kind           TINYINT UNSIGNED NOT NULL,
    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY (study_group_id, user_id),
    FOREIGN KEY (study_group_id) REFERENCES study_groups (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

