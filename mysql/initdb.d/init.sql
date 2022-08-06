DROP DATABASE IF EXISTS grass;
CREATE DATABASE grass;
USE grass;

CREATE TABLE training (
    date DATE,
    count INT
);

INSERT INTO training (date, count) VALUE ('2022-08-01', 3);
INSERT INTO training (date, count) VALUE ('2022-08-02', 4);
INSERT INTO training (date, count) VALUE ('2022-08-03', 1);
