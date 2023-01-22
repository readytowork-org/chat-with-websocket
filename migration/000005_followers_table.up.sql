CREATE TABLE IF NOT EXISTS followers
(
    user_id        VARCHAR(50) NOT NULL,
    follow_user_id VARCHAR(50) NOT NULL,
    created_at     DATETIME    NOT NULL,
    PRIMARY KEY (user_id, follow_user_id),
    INDEX follow_user_id_fk_idx (follow_user_id ASC) VISIBLE,
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
    DEFAULT CHARSET = utf8mb4;