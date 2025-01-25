CREATE TABLE IF NOT EXISTS household_chore (
    household_id INT NOT NULL,
    chore_id INT NOT NULL,
    PRIMARY KEY (household_id, chore_id),
    FOREIGN KEY (household_id) REFERENCES household(id),
    FOREIGN KEY (chore_id) REFERENCES chore(id)
);