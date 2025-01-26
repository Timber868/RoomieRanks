CREATE TABLE IF NOT EXISTS chore_trade_item (
    `sender_username` VARCHAR(255) NOT NULL,
    `chore_id` INT UNSIGNED NOT NULL,
    `trade_request_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (sender_username, chore_id, trade_request_id),
    FOREIGN KEY (sender_username) REFERENCES users(username),
    FOREIGN KEY (chore_id) REFERENCES chore(id),
    FOREIGN KEY (trade_request_id) REFERENCES trade_request(id)
);