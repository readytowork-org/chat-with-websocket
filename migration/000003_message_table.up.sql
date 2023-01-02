CRAETE TABLE IF NOT EXISTS messages(
    id INT NOT NULL AUTO_INCREMENT,
    room_id INT,
    text LONG TEXT,
    user_id INT ,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (room_id) REFERENCES room(room_id)
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);