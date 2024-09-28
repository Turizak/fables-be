DO $$
DECLARE
table_name TEXT;
BEGIN
FOR table_name IN
SELECT tablename
FROM pg_tables
WHERE tablename LIKE '%_characters'
LOOP
EXECUTE 'ALTER TABLE ' || table_name || ' ADD COLUMN is_public boolean DEFAULT true;';
END LOOP;
END $$;