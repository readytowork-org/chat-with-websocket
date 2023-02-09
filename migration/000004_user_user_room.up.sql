CREATE TABLE IF NOT EXISTS user_rooms
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    follower_id INT UNSIGNED NOT NULL,
    room_id     INT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    INDEX user_rooms_room_id_idx (room_id),
    CONSTRAINT user_rooms_follower_id_fk FOREIGN KEY (follower_id) REFERENCES followers (id),
    CONSTRAINT user_rooms_room_id_fk FOREIGN KEY (room_id) REFERENCES rooms (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4;

