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