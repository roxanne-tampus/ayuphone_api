INSERT INTO devices (brand, model, release_year) VALUES
-- iPhone
('iPhone', 'iPhone XS', 2018),
('iPhone', 'iPhone XS Max', 2018),
('iPhone', 'iPhone XR', 2018),
('iPhone', 'iPhone 11', 2019),
('iPhone', 'iPhone 11 Pro', 2019),
('iPhone', 'iPhone 11 Pro Max', 2019),
('iPhone', 'iPhone SE (2nd generation)', 2020),
('iPhone', 'iPhone 12', 2020),
('iPhone', 'iPhone 12 mini', 2020),
('iPhone', 'iPhone 12 Pro', 2020),
('iPhone', 'iPhone 12 Pro Max', 2020),
('iPhone', 'iPhone 13', 2021),
('iPhone', 'iPhone 13 mini', 2021),
('iPhone', 'iPhone 13 Pro', 2021),
('iPhone', 'iPhone 13 Pro Max', 2021),
('iPhone', 'iPhone 14', 2022),
('iPhone', 'iPhone 14 Pro', 2022),
('iPhone', 'iPhone 14 Pro Max', 2022),
('iPhone', 'iPhone 15', 2023),
('iPhone', 'iPhone 15 Pro', 2023),
('iPhone', 'iPhone 15 Pro Max', 2023),

-- Samsung
('Samsung', 'Galaxy S9', 2018),
('Samsung', 'Galaxy S9+', 2018),
('Samsung', 'Galaxy Note9', 2018),
('Samsung', 'Galaxy S10', 2019),
('Samsung', 'Galaxy S10+', 2019),
('Samsung', 'Galaxy Note10', 2019),
('Samsung', 'Galaxy S20', 2020),
('Samsung', 'Galaxy S20+', 2020),
('Samsung', 'Galaxy Note20', 2020),
('Samsung', 'Galaxy S21', 2021),
('Samsung', 'Galaxy S21+', 2021),
('Samsung', 'Galaxy S22', 2022),
('Samsung', 'Galaxy S22+', 2022),
('Samsung', 'Galaxy S23', 2023),
('Samsung', 'Galaxy S23+', 2023),

-- Huawei
('Huawei', 'P20', 2018),
('Huawei', 'P20 Pro', 2018),
('Huawei', 'Mate 20', 2018),
('Huawei', 'P30', 2019),
('Huawei', 'P30 Pro', 2019),
('Huawei', 'Mate 30', 2019),
('Huawei', 'P40', 2020),
('Huawei', 'P40 Pro', 2020),
('Huawei', 'Mate 40', 2020),
('Huawei', 'P50', 2021),
('Huawei', 'P50 Pro', 2021),
('Huawei', 'Mate 50', 2022),

-- Vivo
('Vivo', 'V11', 2018),
('Vivo', 'V15', 2019),
('Vivo', 'V17', 2019),
('Vivo', 'V19', 2020),
('Vivo', 'V21', 2021),
('Vivo', 'V23', 2022),
('Vivo', 'V25', 2023),

-- Oppo
('Oppo', 'F9', 2018),
('Oppo', 'F11', 2019),
('Oppo', 'Reno 2', 2019),
('Oppo', 'Reno 3', 2020),
('Oppo', 'Reno 4', 2020),
('Oppo', 'Reno 5', 2021),
('Oppo', 'Reno 6', 2021),
('Oppo', 'Reno 7', 2022),
('Oppo', 'Reno 8', 2023),

-- Realme
('Realme', '2 Pro', 2018),
('Realme', '3 Pro', 2019),
('Realme', 'X2 Pro', 2019),
('Realme', '6 Pro', 2020),
('Realme', '7 Pro', 2020),
('Realme', '8 Pro', 2021),
('Realme', '9 Pro', 2022),
('Realme', '10 Pro', 2023);


INSERT INTO device_issues (issue_description) VALUES
('Screen cracked or shattered'),
('Battery not charging'),
('Overheating'),
('Phone not turning on'),
('Speaker not working'),
('Microphone not working'),
('Camera not functioning'),
('Touchscreen unresponsive'),
('Phone freezing or lagging'),
('Software update failed'),
('Water damage'),
('Charging port not working'),
('Headphone jack not working'),
('Wi-Fi or Bluetooth connectivity issues'),
('Signal reception problems'),
('SIM card not detected'),
('Fingerprint sensor not working'),
('Face ID/Facial recognition not working'),
('Battery draining quickly'),
('Phone restarting randomly'),
('Unable to make or receive calls'),
('Phone stuck in boot loop'),
('Unable to connect to mobile data'),
('Screen flickering'),
('Volume buttons not working'),
('Power button not working'),
('Home button not working'),
('App crashing frequently'),
('Unable to install or update apps'),
('Phone overheating during charging'),
('Proximity sensor not working'),
('No sound during calls'),
('Phone vibrating continuously or not vibrating'),
('Touch ID not working'),
('Phone not syncing with other devices'),
('Phone not recognizing accessories'),
('Low storage space warnings'),
('Corrupted or missing data'),
('Slow charging'),
('Battery swelling');


INSERT INTO regions (name) VALUES
('Central Visayas');


INSERT INTO provinces (name, region_id) VALUES
('Cebu', 1);


-- Insert Municipalities in Cebu Province
INSERT INTO municipalities (name, province_id) VALUES
('Cebu City', 1),
('Lapu-Lapu City', 1),
('Mandaue City', 1),
('Talisay City', 1),
('Carcar City', 1),
('Danao City', 1),
('Naga City', 1),
('Bogo City', 1),
('Toledo City', 1),
('Argao', 1),
('Barili', 1),
('Boljoon', 1),
('Consolacion', 1),
('Cordova', 1),
('Dalaguete', 1),
('Dumanjug', 1),
('Ginatilan', 1),
('Liloan', 1),
('Moalboal', 1),
('Oslob', 1),
('Pilar', 1),
('Pinamungajan', 1),
('Ronda', 1),
('San Fernando', 1),
('San Francisco', 1),
('Santander', 1),
('Sibonga', 1),
('Sogod', 1),
('Tabogon', 1),
('Tuburan', 1),
('Tudela', 1);


-- Insert Barangays for Cebu Province Municipalities
-- Cebu City
INSERT INTO barangays (name, municipality_id) VALUES
('Adlawon', 1),
('Agsungot', 1),
('Apas', 1),
('Bacayan', 1),
('Banilad', 1),
('Basak Pardo', 1),
('Basak San Nicolas', 1),
('Bonbon', 1),
('Budlaan', 1),
('Busay', 1),
('Calamba', 1),
('Capitol Site', 1),
('Carreta', 1),
('Cogon Pardo', 1),
('Guadalupe', 1),
('Kasambagan', 1),
('Labangon', 1),
('Lahug', 1),
('Mabolo', 1),
('Pasil', 1),
('Pung-ol Sibugay', 1),
('Sambag I', 1),
('Sambag II', 1),
('Talamban', 1),
('Tisa', 1);

-- Lapu-Lapu City
INSERT INTO barangays (name, municipality_id) VALUES
('Agus', 2),
('Babag', 2),
('Basak', 2),
('Buaya', 2),
('Calawisan', 2),
('Canjulao', 2),
('Gun-ob', 2),
('Ibo', 2),
('Looc', 2),
('Maribago', 2),
('Marigondon', 2),
('Pajac', 2),
('Pajo', 2),
('Poblacion', 2),
('Subabasbas', 2);

-- Mandaue City
INSERT INTO barangays (name, municipality_id) VALUES
('Alang-alang', 3),
('Bakilid', 3),
('Banilad', 3),
('Basak', 3),
('Cabancalan', 3),
('Casili', 3),
('Cubacub', 3),
('Canduman', 3),
('Guizo', 3),
('Labogon', 3),
('Looc', 3),
('Mantuyong', 3),
('Opao', 3),
('Pakna-an', 3),
('Subangdaku', 3);

-- Talisay City
INSERT INTO barangays (name, municipality_id) VALUES
('Biasong', 4),
('Bulacao', 4),
('Cansojong', 4),
('Dumlog', 4),
('Lawaan I', 4),
('Lawaan II', 4),
('Lawaan III', 4),
('Linao', 4),
('Maghaway', 4),
('Manipis', 4),
('Mohon', 4),
('Poblacion', 4),
('San Isidro', 4),
('Tabunok', 4),
('Tangke', 4);

-- Carcar City
INSERT INTO barangays (name, municipality_id) VALUES
('Bolocboloc', 5),
('Buenavista', 5),
('Calidngan', 5),
('Can-asujan', 5),
('Guadalupe', 5),
('Liburon', 5),
('Napo', 5),
('Ocana', 5),
('Perrelos', 5),
('Poblacion I', 5),
('Poblacion II', 5),
('Poblacion III', 5),
('Tuyom', 5),
('Valladolid', 5);

-- Argao
INSERT INTO barangays (name, municipality_id) VALUES
('Alambijud', 10),
('Anajao', 10),
('Bala-as', 10),
('Balisong', 10),
('Binlod', 10),
('Bulasa', 10),
('Capio-an', 10),
('Conalum', 10),
('Dagatan', 10),
('Langtad', 10),
('Mabasa', 10),
('Mompeller', 10),
('Poblacion', 10),
('Talaytay', 10),
('Tapon', 10),
('Tiguib', 10);

-- Danao City
INSERT INTO barangays (name, municipality_id) VALUES
('Danasan', 6),
('Guinsay', 6),
('Looc', 6),
('Mantija', 6),
('Poblacion', 6),
('Santo Tomas', 6),
('Suba', 6),
('Taytay', 6);

-- Oslob
INSERT INTO barangays (name, municipality_id) VALUES
('Aguinid', 20),
('Bonbon', 20),
('Calumpang', 20),
('Cañang', 20),
('Looc', 20),
('Luka', 20),
('Poblacion', 20),
('Tan-awan', 20),
('Tigbao', 20),
('Tumalog', 20);

-- Moalboal
INSERT INTO barangays (name, municipality_id) VALUES
('Basdiot', 19),
('Bugho', 19),
('Saavedra', 19),
('Tuble', 19),
('Balabagon', 19),
('Poblacion East', 19),
('Poblacion West', 19);
