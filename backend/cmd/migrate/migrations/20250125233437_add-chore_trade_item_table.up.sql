CREATE TABLE IF NOT EXISTS chore_trade_item (
    sender_id INT NOT NULL,
    chore_id INT NOT NULL,
    trade_request_id INT NOT NULL,
    PRIMARY KEY (sender_id, chore_id, trade_request_id),
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (chore_id) REFERENCES chore(id),
    FOREIGN KEY (trade_request_id) REFERENCES trade_request(id)
);