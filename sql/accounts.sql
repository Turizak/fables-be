DROP TABLE IF EXISTS accounts;

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    uuid character varying NOT NULL UNIQUE,
    email character varying NOT NULL UNIQUE,
    username character varying NOT NULL UNIQUE,
    password character varying NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    created timestamp with time zone NOT NULL
);