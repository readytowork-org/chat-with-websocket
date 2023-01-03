CREATE TABLE IF NOT EXISTS rooms(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(45) NOT NULL,
    owner_id INT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (owner_id) REFERENCES users(id)
);