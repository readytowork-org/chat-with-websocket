CREATE TABLE IF NOT EXISTS users
(
    id         VARCHAR(28)  NOT NULL,
    full_name  VARCHAR(45)  NULL,
    email      VARCHAR(100) NOT NULL,
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NULL,
    deleted_at DATETIME     NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;