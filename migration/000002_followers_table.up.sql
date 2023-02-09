CREATE TABLE IF NOT EXISTS followers
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id        VARCHAR(28)  NOT NULL,
    follow_user_id VARCHAR(28)  NOT NULL,
    created_at     DATETIME     NOT NULL,
    updated_at     DATETIME     NOT NULL,
    deleted_at     DATETIME     NULL,
    PRIMARY KEY (id),
    CONSTRAINT followers_unique
        UNIQUE (follow_user_id, user_id),
    INDEX follow_user_id_fk_idx (id ASC),
    CONSTRAINT f_user_id_fk
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT follow_user_id_fk
        FOREIGN KEY (follow_user_id)
            REFERENCES users (id)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARSET = UTF8MB4;

