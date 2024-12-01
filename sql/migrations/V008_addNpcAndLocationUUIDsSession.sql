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