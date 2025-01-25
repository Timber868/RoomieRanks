CREATE TABLE IF NOT EXISTS collectible_trade_item (
    sender_id INT NOT NULL,
    collectible_id INT NOT NULL,
    trade_request_id INT NOT NULL,
    PRIMARY KEY (sender_id, collectible_id, trade_request_id),
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (collectible_id) REFERENCES collectible(id),
    FOREIGN KEY (trade_request_id) REFERENCES trade_request(id)
);