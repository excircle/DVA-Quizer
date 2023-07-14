CREATE TABLE questions (
    id              INT AUTO_INCREMENT NOT NULL,
    domain          INT NOT NULL,
    question        VARCHAR(255) NOT NULL,
    answer          VARCHAR(255) NOT NULL,
    description     TEXT NOT NULL,
    miscellaneous   TINYINT NOT NULL,
    PRIMARY KEY(`id`),
    KEY(`domain`),
    KEY(`miscellaneous`)
);
