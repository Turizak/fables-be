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