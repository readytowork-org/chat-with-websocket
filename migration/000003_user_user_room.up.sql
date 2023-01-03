CREATE TABLE IF NOT EXISTS user_rooms(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT, 
    room_id INT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);