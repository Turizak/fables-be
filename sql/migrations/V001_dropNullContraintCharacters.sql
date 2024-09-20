DO $$
DECLARE
table_name TEXT;
BEGIN
FOR table_name IN
SELECT tablename
FROM pg_tables
WHERE tablename LIKE '%_characters'
LOOP
EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN age DROP NOT NULL;';
EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN height DROP NOT NULL;';
EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN weight DROP NOT NULL;';
EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN eye_color DROP NOT NULL;';
EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN skin_color DROP NOT NULL;';
EXECUTE 'ALTER TABLE ' || table_name || ' ALTER COLUMN hair_color DROP NOT NULL;';
EXECUTE 'ALTER TABLE ' || table_name || ' ADD COLUMN gender VARCHAR NULL;';
END LOOP;
END $$;