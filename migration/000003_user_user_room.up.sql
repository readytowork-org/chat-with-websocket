CREATE TABLE IF NOT EXISTS user_rooms
(
    user_id    VARCHAR(28) NOT NULL,
    room_id    INT         NOT NULL,
    is_private BOOL NOT NULL,
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (user_id, room_id),
    INDEX user_rooms_room_id_idx (room_id),
    CONSTRAINT user_rooms_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT user_rooms_room_id_fk FOREIGN KEY (room_id) REFERENCES rooms (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

