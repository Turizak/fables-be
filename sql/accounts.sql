CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    email character varying NOT NULL UNIQUE,
    password character varying NOT NULL,
    uuid character varying NOT NULL UNIQUE,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    created timestamp with time zone NOT NULL
);