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
