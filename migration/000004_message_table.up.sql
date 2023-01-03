CREATE TABLE IF NOT EXISTS messages (
    id INT NOT NULL AUTO_INCREMENT,
    text LONGTEXT,
    user_room_id INT ,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_room_id) REFERENCES user_rooms(id)
);