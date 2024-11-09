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