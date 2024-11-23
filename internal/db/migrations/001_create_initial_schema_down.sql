-- Drop tables in reverse order to avoid foreign key issues
DROP TABLE IF EXISTS barangays;
DROP TABLE IF EXISTS municipalities;
DROP TABLE IF EXISTS provinces;
DROP TABLE IF EXISTS regions;
DROP TABLE IF EXISTS transaction_technicians;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS organization_users;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS device_issues;
DROP TABLE IF EXISTS devices;
DROP TABLE IF EXISTS roles;

DROP VIEW IF EXISTS transaction_with_devices;
