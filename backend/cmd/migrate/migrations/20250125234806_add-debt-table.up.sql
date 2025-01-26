CREATE TABLE IF NOT EXISTS debt (
    `bill_id` INT UNSIGNED NOT NULL,
    `user_username` VARCHAR(255) NOT NULL,
    `amount` DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (bill_id, user_username),
    FOREIGN KEY (bill_id) REFERENCES bill (id),
    FOREIGN KEY (user_username) REFERENCES users (username)
);