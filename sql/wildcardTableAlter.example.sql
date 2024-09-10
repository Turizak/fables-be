DO $$
DECLARE
    table_name TEXT;
BEGIN
    FOR table_name IN
        SELECT tablename
        FROM pg_tables
        WHERE tablename LIKE '%_characters'
    LOOP
        EXECUTE 'ALTER TABLE ' || table_name || ' ADD COLUMN IF NOT EXISTS deleted BOOLEAN DEFAULT FALSE;';
    END LOOP;
END $$;

-- Will be useful reference for appending columns to tables that have a common suffix when many tables are involved and can't drop and recreate. 