CREATE TABLE IF NOT EXISTS messages
(
    id           INT          NOT NULL AUTO_INCREMENT,
    text         LONGTEXT,
    user_room_id INT UNSIGNED NOT NULL,
    created_at   DATETIME     NOT NULL,
    updated_at   DATETIME     NULL,
    deleted_at   DATETIME     NULL,
    PRIMARY KEY (id),
    INDEX messages_id_idx (id),
    INDEX messages_user_room_id_fk_idx (user_room_id),
    CONSTRAINT messages_user_room_id_fk FOREIGN KEY (user_room_id) REFERENCES user_rooms (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4;
