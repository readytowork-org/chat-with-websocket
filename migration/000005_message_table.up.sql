CREATE TABLE IF NOT EXISTS messages
(
    id         INT          NOT NULL AUTO_INCREMENT,
    text       LONGTEXT,
    room_id    INT UNSIGNED NOT NULL,
    user_id    VARCHAR(28)  NOT NULL,
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NULL,
    deleted_at DATETIME     NULL,
    PRIMARY KEY (id),
    INDEX messages_id_idx (id),
    INDEX messages_room_id_fk_idx (room_id),
    CONSTRAINT messages_room_id_fk FOREIGN KEY (room_id) REFERENCES rooms (id),
    CONSTRAINT messages_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4;
