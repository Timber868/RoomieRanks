CREATE TABLE IF NOT EXISTS debt (
    bill_id INT NOT NULL,
    user_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (bill_id, user_id),
    FOREIGN KEY (debtor_id) REFERENCES bill (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);