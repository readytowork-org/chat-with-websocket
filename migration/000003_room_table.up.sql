CREATE TABLE IF NOT EXISTS rooms
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name        VARCHAR(45)  NOT NULL,
    is_private  BOOL         NOT NULL,
    follower_id INT UNSIGNED NOT NULL,
    created_at  DATETIME     NOT NULL,
    updated_at  DATETIME     NULL,
    deleted_at  DATETIME     NULL,
    PRIMARY KEY (id),
    CONSTRAINT user_rooms_follower_id_fk FOREIGN KEY (follower_id) REFERENCES followers (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = UTF8MB4;
