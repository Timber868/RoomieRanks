CREATE TABLE IF NOT EXISTS bill (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `purchase_date` DATE NOT NULL,
    `payer_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);