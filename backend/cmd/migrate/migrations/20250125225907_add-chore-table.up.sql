CREATE TABLE IF NOT EXISTS chore (
    `ID` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `difficulty` INT UNSIGNED NOT NULL,
    `timeEstimate` INT UNSIGNED NOT NULL,
    `completionTime` INT UNSIGNED NOT NULL,
    `householdID` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`householdID`) REFERENCES households(`id`)
);