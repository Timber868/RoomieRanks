CREATE TABLE IF NOT EXISTS bill (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `purchase_date` DATE NOT NULL,
    `payer_username` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`payer_username`) REFERENCES users(`username`)
);