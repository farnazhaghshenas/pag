CREATE TABLE cities
(
    id            int auto_increment PRIMARY KEY,
    name         mediumtext   NOT NULL
);

CREATE TABLE trips
(
    id            int auto_increment,
    originId      int        NOT NULL,
    destinationId int        NOT NULL,
    dates         mediumtext NOT NULL,
    price         float(23, 0) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (originId) REFERENCES cities(id),
    FOREIGN KEY (destinationId) REFERENCES cities(id)
);
