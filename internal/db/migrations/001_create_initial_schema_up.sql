-- Create roles table
CREATE TABLE IF NOT EXISTS roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Create devices table
CREATE TABLE IF NOT EXISTS devices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    brand VARCHAR(50) NOT NULL,
    model VARCHAR(100) NOT NULL,
    release_year INT NOT NULL,
    UNIQUE (brand, model, release_year)
);

-- Create device_issues table
CREATE TABLE IF NOT EXISTS device_issues (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    issue_description VARCHAR(255) NOT NULL,
    UNIQUE (issue_description)
);

-- Create organizations table
CREATE TABLE IF NOT EXISTS organizations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER IF NOT EXISTS update_organizations_updated_at
AFTER UPDATE ON organizations
FOR EACH ROW
BEGIN
    UPDATE organizations SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(255) UNIQUE,
    password VARCHAR(255) NOT NULL,
    role_id TINYINT NOT NULL DEFAULT 3,
    address VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id),
    CHECK (role_id IN (1, 2, 3, 4)) -- Assuming roles are predefined
);

CREATE TRIGGER IF NOT EXISTS update_users_updated_at
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- Create organization_users table
CREATE TABLE IF NOT EXISTS organization_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT,
    organization_id INT,
    role_id TINYINT NOT NULL DEFAULT 4,
    invited_by INT,
    status TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, organization_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE,
    FOREIGN KEY (invited_by) REFERENCES users(id)
);

CREATE TRIGGER IF NOT EXISTS update_organization_users_updated_at
AFTER UPDATE ON organization_users
FOR EACH ROW
BEGIN
    UPDATE organization_users SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    customer_id INT NOT NULL,
    device_id INT NOT NULL,
    device_issue_id INT,
    note TEXT,
    status TEXT NOT NULL DEFAULT 'pending',
    pickup_address TEXT NOT NULL,
    full_address TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (device_id) REFERENCES devices(id) ON DELETE CASCADE,
    FOREIGN KEY (device_issue_id) REFERENCES device_issues(id) ON DELETE CASCADE
);

CREATE TRIGGER IF NOT EXISTS update_transactions_updated_at
AFTER UPDATE ON transactions
FOR EACH ROW
BEGIN
    UPDATE transactions SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- Create transaction_technicians table
CREATE TABLE IF NOT EXISTS transaction_technicians (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    transaction_id INT NOT NULL,
    technician_id INT NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    unassigned_at TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
    FOREIGN KEY (technician_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TRIGGER IF NOT EXISTS update_transaction_technicians_unassigned_at
AFTER UPDATE ON transaction_technicians
FOR EACH ROW
BEGIN
    UPDATE transaction_technicians SET unassigned_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- FOR PICKUP ADDRESS
-- Table for Regions
CREATE TABLE IF NOT EXISTS regions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL
);

-- Table for Provinces
CREATE TABLE IF NOT EXISTS provinces (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    region_id INT NOT NULL,
    FOREIGN KEY (region_id) REFERENCES regions(id)
);

-- Table for Municipalities
CREATE TABLE IF NOT EXISTS municipalities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    province_id INT NOT NULL,
    FOREIGN KEY (province_id) REFERENCES provinces(id)
);

-- Table for Barangays
CREATE TABLE IF NOT EXISTS barangays (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    municipality_id INT NOT NULL,
    FOREIGN KEY (municipality_id) REFERENCES municipalities(id)
);

CREATE VIEW transaction_with_devices AS
SELECT
    t.id,
    t.status,
    t.pickup_address,
    t.full_address,
    d.brand,
    d.model,
    di.issue_description,
    t.created_at,
    t.updated_at
FROM
    transactions t
LEFT JOIN
    devices d ON t.device_id = d.id
LEFT JOIN
    device_issues di ON t.device_issue_id = di.id;