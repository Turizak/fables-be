-- Begin V001_dropNullContraintCharacters.sql
DO $$ 
DECLARE table_name TEXT;

BEGIN FOR table_name IN
SELECT
    tablename
FROM
    pg_tables
WHERE
    tablename LIKE '%_characters' LOOP EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN age DROP NOT NULL;';

EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN height DROP NOT NULL;';

EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN weight DROP NOT NULL;';

EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN eye_color DROP NOT NULL;';

EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN skin_color DROP NOT NULL;';

EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN hair_color DROP NOT NULL;';

EXECUTE 'ALTER TABLE ' || table_name || ' ADD COLUMN gender VARCHAR NULL;';

END LOOP;

END $$;
-- End V001_dropNullContraintCharacters.sql

-- Begin V002_createLocationsTablesForExistingMonikers.sql
DO $$ 
DECLARE
    campaign_record RECORD;
    table_name VARCHAR(255);
BEGIN
    -- Loop through each campaign_uuid in the campaigns table
    FOR campaign_record IN 
        SELECT uuid, moniker FROM campaigns
    LOOP
        -- Generate the table name by prepending the moniker to '_locations'
        table_name := campaign_record.moniker || '_locations';

        -- Execute the dynamic SQL to create the table
        EXECUTE format('
            CREATE TABLE IF NOT EXISTS %I (
                id SERIAL PRIMARY KEY,
                campaign_uuid VARCHAR(255) NOT NULL,
                uuid VARCHAR(255) NOT NULL UNIQUE,
                creator_uuid VARCHAR(255) NOT NULL,
                name VARCHAR(255) NOT NULL,
                description VARCHAR(255) NOT NULL,
                created TIMESTAMP WITH TIME ZONE NOT NULL,
                last_updated TIMESTAMP WITH TIME ZONE
            );
        ', table_name);

        RAISE NOTICE 'Table % created successfully for campaign UUID: %.', table_name, campaign_record.uuid;

    END LOOP;

END $$;

-- End V002_createLocationsTablesForExistingMonikers.sql

-- Begin V003_addPublicBoolToCharDefaultTrue.sql
DO $$ 
DECLARE table_name TEXT;

BEGIN FOR table_name IN
SELECT
    tablename
FROM
    pg_tables
WHERE
    tablename LIKE '%_characters' LOOP EXECUTE 'ALTER TABLE ' || table_name || ' ADD COLUMN is_public boolean DEFAULT true;';

END LOOP;

END $$;
-- End V003_addPublicBoolToCharDefaultTrue.sql

-- Begin V004_createNpcsTablesForExistingCampaigns.sql
DO $$
DECLARE
    campaign_record RECORD;
    table_name VARCHAR(255);
BEGIN
    -- Loop through each campaign_uuid in the campaigns table
    FOR campaign_record IN 
        SELECT uuid, moniker FROM campaigns
    LOOP
        -- Generate the table name by prepending the moniker to '_npcs'
        table_name := campaign_record.moniker || '_npcs';

        -- Execute the dynamic SQL to create the table
        EXECUTE format('
            CREATE TABLE IF NOT EXISTS %I (
                id SERIAL PRIMARY KEY,
                campaign_uuid VARCHAR(255) NOT NULL,
                uuid VARCHAR(255) NOT NULL UNIQUE,
                creator_uuid VARCHAR(255) NOT NULL,
                first_name VARCHAR(255) NOT NULL,
                last_name VARCHAR(255) NOT NULL,
                race VARCHAR(255) NOT NULL,
                class VARCHAR(255) NOT NULL,
                description VARCHAR(255) NOT NULL,
                is_quest_boss BOOLEAN NOT NULL,
                created TIMESTAMP WITH TIME ZONE NOT NULL,
                last_updated TIMESTAMP WITH TIME ZONE
            );
        ', table_name);

        RAISE NOTICE 'Table % created successfully for campaign UUID: %.', table_name, campaign_record.uuid;

    END LOOP;

END $$;

-- End V004_createNpcsTablesForExistingCampaigns.sql

-- Begin V005_createSessionsTablesForExistingCampaigns.sql
DO $$
DECLARE
    campaign_record RECORD;
    table_name VARCHAR(255);
BEGIN
    -- Loop through each campaign_uuid in the campaigns table
    FOR campaign_record IN 
        SELECT uuid, moniker FROM campaigns
    LOOP
        -- Generate the table name by prepending the moniker to '_sessions'
        table_name := campaign_record.moniker || '_sessions';

        -- Execute the dynamic SQL to create the table
        EXECUTE format('
            CREATE TABLE IF NOT EXISTS %I (
                id SERIAL PRIMARY KEY,
                campaign_uuid VARCHAR(255) NOT NULL,
                creator_uuid VARCHAR(255) NOT NULL,
                uuid VARCHAR(255) NOT NULL UNIQUE,
                party_uuids character varying[],
                date_occured TIMESTAMP WITH TIME ZONE NOT NULL,
                created TIMESTAMP WITH TIME ZONE NOT NULL,
                last_updated TIMESTAMP WITH TIME ZONE,
                deleted boolean DEFAULT false
            );
        ', table_name);

        RAISE NOTICE 'Table % created successfully for campaign UUID: %.', table_name, campaign_record.uuid;

    END LOOP;

END $$;

-- End V005_createSessionsTablesForExistingCampaigns.sql

-- Begin V006_addSessionIdToSession.sql
DO $$ 
DECLARE table_name TEXT;

BEGIN FOR table_name IN
SELECT
    tablename
FROM
    pg_tables
WHERE
    tablename LIKE '%_sessions' LOOP EXECUTE 'ALTER TABLE ' || table_name || ' ADD COLUMN session_id INT';

END LOOP;

END $$;
-- End V006_addSessionIdToSession.sql

-- Begin V007_addDeletedNpcAndLocation.sql
DO $$ 
DECLARE 
    table_name TEXT;
    BEGIN 
        FOR table_name IN
                SELECT tablename
                    FROM pg_tables
                        WHERE tablename LIKE '%npcs' OR tablename LIKE '%locations'
                            LOOP 
                                EXECUTE 'ALTER TABLE ' || quote_ident(table_name) || ' ADD COLUMN IF NOT EXISTS deleted boolean DEFAULT false';
                            END LOOP;
END $$;
-- End V007_addDeletedNpcAndLocation.sql

-- Begin V008_addNpcAndLocationUUIDsSession.sql
DO $$ 
DECLARE 
    table_name TEXT;
    BEGIN 
        FOR table_name IN
                SELECT tablename
                    FROM pg_tables
                        WHERE tablename LIKE '%sessions'
                            LOOP 
                                EXECUTE 'ALTER TABLE ' || quote_ident(table_name) || ' ADD COLUMN IF NOT EXISTS npc_uuids character varying[]';
                                EXECUTE 'ALTER TABLE ' || quote_ident(table_name) || ' ADD COLUMN IF NOT EXISTS location_uuids character varying[]';
                            END LOOP;
END $$;
-- End V008_addNpcAndLocationUUIDsSession.sql

-- Begin V009_addQuestBossUuidNpc.sql
DO $$ 
DECLARE 
    table_name TEXT;
    BEGIN 
        FOR table_name IN
                SELECT tablename
                    FROM pg_tables
                        WHERE tablename LIKE '%npcs'
                            LOOP 
                                EXECUTE 'ALTER TABLE ' || quote_ident(table_name) || ' ADD COLUMN IF NOT EXISTS quest_boss_uuid VARCHAR(255)';
                            END LOOP;
END $$;
-- End V009_addQuestBossUuidNpc.sql

-- Begin V010_createQuestsTablesForExisting.sql
DO $$ 
DECLARE campaign_record RECORD;

table_name VARCHAR(255);

BEGIN -- Loop through each campaign_uuid in the campaigns table
FOR campaign_record IN
SELECT
    uuid,
    moniker
FROM
    campaigns LOOP -- Generate the table name by prepending the moniker to '_quests'
    table_name := campaign_record.moniker || '_quests';

-- Execute the dynamic SQL to create the table
EXECUTE format(
    '
            CREATE TABLE IF NOT EXISTS %I (
                id SERIAL PRIMARY KEY,
                campaign_uuid VARCHAR(255) NOT NULL,
                uuid VARCHAR(255) NOT NULL UNIQUE,
                creator_uuid VARCHAR(255) NOT NULL,
                name VARCHAR(255) NOT NULL,
                description VARCHAR(255),
                quest_giver VARCHAR(255) NOT NULL,
                reward_uuids VARCHAR(255)[],
                location_uuids VARCHAR(255)[],
                npc_uuids VARCHAR(255)[],
                party_uuids VARCHAR(255)[],
                boss_uuids VARCHAR(255)[],
                starting_session_uuid VARCHAR(255),
                ending_session_uuid VARCHAR(255),
                status VARCHAR(255) NOT NULL,
                created TIMESTAMP WITH TIME ZONE NOT NULL,
                last_updated TIMESTAMP WITH TIME ZONE,
                deleted BOOLEAN DEFAULT FALSE
            );
        ',
    table_name
);

RAISE NOTICE 'Table % created successfully for campaign UUID: %.',
table_name,
campaign_record.uuid;

END LOOP;

END $$;
-- End V010_createQuestsTablesForExisting.sql

-- Begin V011_createNotesTableForExisting.sql
DO $$ 
DECLARE 
    campaign_record RECORD;
    table_name VARCHAR(255);
BEGIN 
    -- Loop through each campaign_uuid in the campaigns table
    FOR campaign_record IN
        SELECT
            uuid,
            moniker
        FROM
            campaigns 
    LOOP 
        -- Generate the table name by prepending the moniker to '_notes'
        table_name := campaign_record.moniker || '_notes';

        -- Execute the dynamic SQL to create the table
        EXECUTE format(
            '
            CREATE TABLE IF NOT EXISTS %I (
                id SERIAL PRIMARY KEY,
                campaign_uuid VARCHAR(255) NOT NULL,
                session_uuid VARCHAR(255),
                uuid VARCHAR(255) NOT NULL UNIQUE,
                creator_uuid VARCHAR(255) NOT NULL,
                title VARCHAR(255) NOT NULL,
                content VARCHAR(255) NOT NULL,
                created TIMESTAMP WITH TIME ZONE NOT NULL,
                last_updated TIMESTAMP WITH TIME ZONE,
                deleted BOOLEAN DEFAULT FALSE
            );
            ',
            table_name
        );

        RAISE NOTICE 'Table % created successfully for campaign UUID: %.',
            table_name,
            campaign_record.uuid;
    END LOOP;

END $$;

-- End V011_createNotesTableForExisting.sql

