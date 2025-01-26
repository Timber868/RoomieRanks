CREATE TABLE IF NOT EXISTS collectible_trade_item (
    `sender_username` VARCHAR(255) NOT NULL,
    `collectible_id` INT UNSIGNED NOT NULL,
    `trade_request_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (sender_username, collectible_id, trade_request_id),
    FOREIGN KEY (sender_username) REFERENCES users(username),
    FOREIGN KEY (collectible_id) REFERENCES collectible(id),
    FOREIGN KEY (trade_request_id) REFERENCES trade_request(id)
);