DROP TABLE IF EXISTS campaigns;

CREATE TABLE IF NOT EXISTS campaigns (
    id SERIAL PRIMARY KEY,
    uuid character varying NOT NULL UNIQUE,
    name character varying NOT NULL,
    moniker character varying(20) NOT NULL UNIQUE,
    creator_uuid character varying NOT NULL,
    dm_uuid character varying NOT NULL,
    party_uuids character varying[],
    completed boolean NOT NULL DEFAULT FALSE,
    active boolean NOT NULL DEFAULT TRUE,
    ruleset character varying NOT NULL,
    max_players integer NOT NULL,
    created timestamp with time zone NOT NULL,
    last_updated timestamp with time zone
);