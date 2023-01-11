CREATE TABLE IF NOT EXISTS rooms
(
    id         INT         NOT NULL AUTO_INCREMENT,
    name       VARCHAR(45) NOT NULL,
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;