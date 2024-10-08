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
