-- Drop foreign key constraints that depend on to_do_list
ALTER TABLE to_do DROP CONSTRAINT IF EXISTS to_do_list_id_fkey;

-- Drop the to_do_list table
DROP TABLE IF EXISTS to_do_list;