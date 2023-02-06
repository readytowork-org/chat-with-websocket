ALTER TABLE messages
    ADD status ENUM ('Sending', 'Sent') NOT NULL AFTER text;
