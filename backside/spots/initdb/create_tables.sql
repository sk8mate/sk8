CREATE DATABASE spotsDb;

DROP TABLE IF EXISTS spots;

CREATE TYPE coordinates AS
(
    lat  double precision,
    long double precision
);

CREATE TABLE IF NOT EXISTS spots
(
    id          INT          NOT NULL,
    name        varchar(250) NOT NULL,
    address     varchar(250) NOT NULL,
    coordinates coordinates  NOT NULL,
    lighting    boolean      NOT NULL,
    friendly    boolean      NOT NULL,
    verified    boolean      NOT NULL,
    PRIMARY KEY (id)
);