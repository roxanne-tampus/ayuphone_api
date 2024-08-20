-- Drop tables in reverse order to avoid foreign key issues
DROP TABLE IF EXISTS organization_users CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS organizations CASCADE;
DROP TABLE transactions CASCADE;
DROP TABLE transaction_technicians CASCADE;
DROP TABLE devices CASCADE;
DROP TABLE device_issues CASCADE;

