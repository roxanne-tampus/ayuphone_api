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


INSERT INTO municipalities (name, province_id) VALUES
('Cebu City', 1),  
('Lapu-Lapu City', 1),  
('Mandaue City', 1),  

-- Component Cities
('Talisay City', 1),
('Carcar City', 1),
('Bogo City', 1),
('Naga City', 1),
('Toledo City', 1),
('Danao City', 1),

-- Municipalities
('Alcantara', 1),
('Alcoy', 1),
('Alegria', 1),
('Aloguinsan', 1),
('Argao', 1),
('Asturias', 1),
('Badian', 1),
('Balamban', 1),
('Barili', 1),
('Boljoon', 1),
('Borbon', 1),
('Catmon', 1),
('Compostela', 1),
('Consolacion', 1),
('Cordova', 1),
('Dalaguete', 1),
('Dumanjug', 1),
('Ginatilan', 1),
('Liloan', 1),
('Malabuyoc', 1),
('Medellin', 1),
('Minglanilla', 1),
('Moalboal', 1),
('Oslob', 1),
('Pilar', 1),
('Pinamungajan', 1),
('Poro', 1),
('Ronda', 1),
('Samboan', 1),
('San Fernando', 1),
('San Francisco', 1),
('Santander', 1),
('Sibonga', 1),
('Sogod', 1),
('Tabogon', 1),
('Tabuelan', 1),
('Tuburan', 1),
('Tudela', 1);


-- Cebu City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Adlawon', 1), ('Agsungot', 1), ('Apas', 1), ('Bacayan', 1), ('Banilad', 1),
('Basak Pardo', 1), ('Basak San Nicolas', 1), ('Bonbon', 1), ('Budlaan', 1), ('Busay', 1),
('Calamba', 1), ('Capitol Site', 1), ('Carreta', 1), ('Cogon Pardo', 1), ('Guadalupe', 1),
('Kasambagan', 1), ('Labangon', 1), ('Lahug', 1), ('Mabolo', 1), ('Pasil', 1),
('Pung-ol Sibugay', 1), ('Sambag I', 1), ('Sambag II', 1), ('Talamban', 1), ('Tisa', 1);

-- Lapu-Lapu City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Agus', 2), ('Babag', 2), ('Basak', 2), ('Buaya', 2), ('Calawisan', 2),
('Canjulao', 2), ('Gun-ob', 2), ('Ibo', 2), ('Looc', 2), ('Mactan', 2),
('Maribago', 2), ('Marigondon', 2), ('Pajac', 2), ('Pajo', 2), ('Poblacion', 2),
('Subabasbas', 2);

-- Mandaue City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Alang-alang', 3), ('Bakilid', 3), ('Banilad', 3), ('Basak', 3), ('Cabancalan', 3),
('Casili', 3), ('Cubacub', 3), ('Canduman', 3), ('Guizo', 3), ('Labogon', 3),
('Looc', 3), ('Mantuyong', 3), ('Opao', 3), ('Pakna-an', 3), ('Subangdaku', 3);

-- Talisay City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Biasong', 4), ('Bulacao', 4), ('Cansojong', 4), ('Dumlog', 4), ('Lawaan I', 4),
('Lawaan II', 4), ('Lawaan III', 4), ('Linao', 4), ('Maghaway', 4), ('Manipis', 4),
('Mohon', 4), ('Poblacion', 4), ('San Isidro', 4), ('Tabunok', 4), ('Tangke', 4);

-- Carcar City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Bolocboloc', 5), ('Buenavista', 5), ('Calidngan', 5), ('Can-asujan', 5), ('Guadalupe', 5),
('Liburon', 5), ('Napo', 5), ('Ocana', 5), ('Perrelos', 5), ('Poblacion I', 5),
('Poblacion II', 5), ('Poblacion III', 5), ('Tuyom', 5), ('Valladolid', 5);

-- Bogo City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Anonang Norte', 6), ('Anonang Sur', 6), ('Banban', 6), ('Binabag', 6), ('Cayang', 6),
('Dakit', 6), ('Don Pedro', 6), ('Gairan', 6), ('Guadalupe', 6), ('La Paz', 6),
('Libertad', 6), ('Lourdes', 6), ('Malingin', 6), ('Marangog', 6), ('Poblacion', 6);

-- Naga City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Alambijud', 7), ('Balirong', 7), ('Cantao-an', 7), ('Cogon', 7), ('Inoburan', 7),
('Jagobiao', 7), ('Lanas', 7), ('Mainit', 7), ('Pangdan', 7), ('Poblacion', 7),
('Tagjaguimit', 7), ('Tangke', 7), ('Tina-an', 7), ('Tuyan', 7);

-- Toledo City Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Awihao', 8), ('Bagakay', 8), ('Bato', 8), ('Bulongan', 8), ('Cabitoonan', 8),
('Canlumampao', 8), ('Cantabaco', 8), ('Capitan Claudio', 8), ('Daanlungsod', 8), ('Don Andres Soriano', 8),
('Ibo', 8), ('Juan Climaco Sr', 8), ('Loay', 8), ('Maalat', 8), ('Media Once', 8),
('Poog', 8), ('Pulangbato', 8), ('Putingbato', 8), ('Sangi', 8), ('Talavera', 8),
('Tubod', 8);

-- Danao City
INSERT INTO barangays (name, municipality_id) VALUES
('Danasan', 9),
('Guinsay', 9),
('Looc', 9),
('Mantija', 9),
('Poblacion', 9),
('Santo Tomas', 9),
('Suba', 9),
('Taytay', 9);



-- Alcantara Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Cabil-isan', 10), ('Camboang', 10), ('Candabong', 10), ('Lawaan', 10),
('Palanas', 10), ('Poblacion', 10), ('Salagmaya', 10), ('San Isidro', 10);

-- Alcoy Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Atabay', 11), ('Bulasa', 11), ('Daan-Lungsod', 11), ('Guiwang', 11), ('Nug-as', 11),
('Pasol', 11), ('Poblacion', 11), ('Pugalo', 11);

-- Alegria Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Compostela', 12), ('Guiwanon', 12), ('Lepanto', 12), ('Madridejos', 12),
('Montpeller', 12), ('Poblacion', 12), ('Santa Filomena', 12), ('Valencia', 12);

-- Aloguinsan Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Angilan', 13), ('Bojo', 13), ('Bonbon', 13), ('Esperanza', 13), ('Kantabugon', 13),
('Olango', 13), ('Poblacion', 13), ('Riverside', 13), ('Rosario', 13), ('Saksak', 13);

-- Argao Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Alambijud', 14), ('Anajao', 14), ('Bala-as', 14), ('Balisong', 14), ('Binlod', 14),
('Bulasa', 14), ('Capio-an', 14), ('Conalum', 14), ('Dagatan', 14), ('Langtad', 14),
('Mabasa', 14), ('Mompeller', 14), ('Poblacion', 14), ('Talaytay', 14), ('Tapon', 14),
('Tiguib', 14);

-- Asturias Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Agbanga', 15), ('Agtugop', 15), ('Bago', 15), ('Banban', 15), ('Baye', 15),
('Bog-o', 15), ('Kabayangan', 15), ('Kambuhawe', 15), ('Lunas', 15), ('Manguiao', 15),
('Owak', 15), ('Poblacion', 15), ('Saksak', 15), ('Tubigagmanok', 15);

-- Badian Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Basak', 16), ('Basiao', 16), ('Bato', 16), ('Bugas', 16), ('Candiis', 16),
('Doldol', 16), ('Ginablan', 16), ('Malhiao', 16), ('Matutinao', 16), ('Patong', 16),
('Poblacion', 16), ('Santicon', 16), ('Sawang', 16), ('Talayong', 16), ('Tigbao', 16);

-- Balamban Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Abucayan', 17), ('Aliwanay', 17), ('Arpili', 17), ('Baliwagan', 17), ('Buanoy', 17),
('Cabagdalan', 17), ('Cambuhawe', 17), ('Cambuhawe Norte', 17), ('Cantuclan', 17), ('Ga-as', 17),
('Matun-og', 17), ('Nangka', 17), ('Pondol', 17), ('Sinsin', 17), ('Sunog', 17), ('Vito', 17);

-- Barili Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Aloguinsan', 18), ('Anopog', 18), ('Bagacay', 18), ('Balabagon', 18), ('Boloc-boloc', 18),
('Cagay', 18), ('Gunting', 18), ('Hilantagaan', 18), ('Japitan', 18), ('Kalunasan', 18),
('Kalunasan Norte', 18), ('Kalunasan Sur', 18), ('Malolos', 18), ('Mataybagon', 18), ('Poblacion', 18);

-- Boljoon Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Alegria', 19), 
('Bantayan', 19), 
('Boljoon Proper', 19), 
('Bunga', 19), 
('Hagnaya', 19),
('Libang', 19), 
('Longon', 19), 
('Malabuyoc', 19), 
('San Juan', 19), 
('Santo Niño', 19), 
('Tawigan', 19);

-- Borbon Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Bagacay', 20), 
('Bantayan', 20), 
('Borbon Proper', 20), 
('Cawag', 20), 
('Kansun', 20),
('Lungsod', 20), 
('Maguikay', 20), 
('Poblacion', 20), 
('San Isidro', 20);

-- Catmon Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Agus', 21), 
('Bagacay', 21), 
('Catmon Proper', 21), 
('Gamay', 21), 
('San Jose', 21),
('Santo Niño', 21), 
('Tabunok', 21), 
('Tuguis', 21);

-- Compostela Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Aloja', 22), 
('Baclayon', 22), 
('Compostela Proper', 22), 
('Dila-dila', 22), 
('Guadalupe', 22),
('Lugay', 22), 
('Poblacion', 22), 
('San Isidro', 22);

-- Consolacion Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Bagalnga', 23), 
('Buanoy', 23), 
('Consolacion Proper', 23), 
('Dapdap', 23), 
('Luyongbon', 23),
('Luyong Baybay', 23), 
('Poblacion', 23), 
('Tayud', 23);

-- Cordova Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Alegria', 24), 
('Cordova Proper', 24), 
('Gabi', 24), 
('Marigondon', 24), 
('Poblacion', 24),
('Tayud', 24);

-- Dalaguete Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Alegria', 25), 
('Dalaguete Proper', 25), 
('Lambug', 25), 
('Maloray', 25), 
('Samboan', 25),
('San Roque', 25), 
('Tungkil', 25);

-- Dumanjug Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Badian', 26), 
('Dumanjug Proper', 26), 
('Lunok', 26), 
('Mananga', 26), 
('San Isidro', 26),
('Santo Niño', 26);

-- Ginatilan Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Anapog', 27), 
('Bantayan', 27), 
('Ginatilan Proper', 27), 
('Mabini', 27), 
('San Jose', 27),
('Santo Niño', 27);

-- Liloan Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Binaliw', 28), 
('Bunga', 28), 
('Burgos', 28), 
('Buwang', 28), 
('Cahumayan', 28),
('Dapdap', 28), 
('Magsaysay', 28), 
('Poblacion', 28), 
('San Vicente', 28);

-- Malabuyoc Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Bunga', 29), 
('Malabuyoc Proper', 29), 
('San Juan', 29), 
('Santa Teresa', 29);

-- Medellin Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Bacong', 30), 
('Bontay', 30), 
('Calumboyan', 30), 
('Caputatan', 30), 
('Daanbantayan', 30),
('Langtad', 30), 
('Medellin Proper', 30), 
('San Jose', 30), 
('Santa Rosa', 30);

-- Minglanilla Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Cadulawan', 31), 
('Calajo-an', 31), 
('Cambuhawe', 31), 
('Camp 7', 31), 
('Cuanos', 31), 
('Guindaruhan', 31), 
('Kabalaasnan', 31), 
('Linao-Lipata', 31), 
('Pakigne', 31), 
('Poblacion Ward I', 31), 
('Poblacion Ward II', 31), 
('Poblacion Ward III', 31), 
('Poblacion Ward IV', 31), 
('Tungkil', 31), 
('Tulay', 31);

-- Moalboal Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Basdiot', 32), 
('Bugho', 32), 
('Saavedra', 32), 
('Tuble', 32), 
('Poblacion East', 32),
('Poblacion West', 32);

-- Oslob Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Aguinid', 33), 
('Bonbon', 33), 
('Calumpang', 33), 
('Cañang', 33), 
('Looc', 33),
('Luka', 33), 
('Poblacion', 33), 
('Tan-awan', 33), 
('Tigbao', 33), 
('Tumalog', 33);

-- Pilar Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Cansuhay', 34), 
('Dapdap', 34), 
('Esperanza', 34), 
('Lanao', 34), 
('Mabini', 34),
('Moabog', 34), 
('Montesuerte', 34), 
('Pitos', 34), 
('Poblacion', 34), 
('San Isidro', 34);

-- Pinamungajan Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Anopog', 35), 
('Binabag', 35), 
('Busay', 35), 
('Butong', 35), 
('Duangan', 35),
('Lamac', 35), 
('Lutopan', 35), 
('Poblacion', 35), 
('Sacsac', 35), 
('Simbahan', 35),
('Tajao', 35), 
('Tangub', 35), 
('Tungkil', 35);

-- Poro Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Adela', 36), 
('Altavista', 36), 
('Cagcagan', 36), 
('Cansabusab', 36), 
('Esperanza', 36),
('Macapili', 36), 
('Mercedes', 36), 
('Poblacion', 36), 
('Punay', 36), 
('San Jose', 36),
('Santo Niño', 36), 
('Teguis', 36);

-- Ronda Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Can-abuhon', 37), 
('Canduling', 37), 
('Cansayahon', 37), 
('Langin', 37), 
('Liboo', 37),
('Malalay', 37), 
('Poblacion', 37), 
('Santa Cruz', 37), 
('Sapu', 37), 
('Tampi', 37);

-- Samboan Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Basak', 38), 
('Bonbon', 38), 
('Calatagan', 38), 
('Colase', 38), 
('Jumangpas', 38),
('Poblacion', 38), 
('San Sebastian', 38), 
('Suba', 38);

-- San Fernando Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Balud', 39), 
('Basak', 39), 
('Bugho', 39), 
('Greenhills', 39), 
('Ilaya', 39),
('Lantawan', 39), 
('Liburon', 39), 
('Magsico', 39), 
('Panadtaran', 39), 
('Poblacion', 39),
('Sambag', 39), 
('San Isidro', 39), 
('Sinsin', 39), 
('Tabionan', 39), 
('Talisay', 39),
('Tonggo', 39), 
('Tubod', 39), 
('Tungkop', 39);

-- San Francisco Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Cagcagan', 40), 
('Consuelo', 40), 
('Esperanza', 40), 
('Haguilanan', 40), 
('Montealegre', 40),
('Northern Poblacion', 40), 
('San Isidro', 40), 
('San Roque', 40), 
('Santa Cruz', 40),
('Southern Poblacion', 40), 
('Union', 40);

-- Santander Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Bacong', 41), 
('Bolocboloc', 41), 
('Dapdap', 41), 
('Gumamela', 41), 
('Libertad', 41),
('Magsaysay', 41), 
('Malabago', 41), 
('Poblacion', 41), 
('San Isidro', 41), 
('San Juan', 41),
('San Miguel', 41), 
('Talisay', 41), 
('Tigbao', 41);

-- Sibonga Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Abunton', 42), 
('Alang-alang', 42), 
('Bato', 42), 
('Biga', 42), 
('Bonbon', 42),
('Budbud', 42), 
('Gumbayan', 42), 
('Mahuay', 42), 
('Poblacion', 42), 
('San Felipe', 42),
('San Jose', 42), 
('Santo Niño', 42), 
('Tabili', 42), 
('Tibasak', 42), 
('Tinaan', 42),
('Tulay', 42);

-- Sogod Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Alang-alang', 43), 
('Bagalupa', 43), 
('Balayniwang', 43), 
('Bantayan', 43), 
('Canghumol', 43),
('Cangmanghag', 43), 
('Catmon', 43), 
('Consolacion', 43), 
('Dolores', 43), 
('Galingon', 43),
('Guadalupe', 43), 
('Malilina', 43), 
('Poblacion', 43), 
('San Jose', 43), 
('San Roque', 43),
('Tugas', 43), 
('Tungkil', 43);

-- Tabogon Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Ban-ban', 44), 
('Bunyon', 44), 
('Gimagaan', 44), 
('Guadalupe', 44), 
('Lusaran', 44),
('Manghuli', 44), 
('Poblacion', 44), 
('San Juan', 44), 
('San Vicente', 44), 
('Santo Niño', 44),
('Tubod', 44);

-- Tabuelan Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Agbangan', 45), 
('Binas-angan', 45), 
('Bunyan', 45), 
('Calape', 45), 
('Gingon', 45),
('Magsaysay', 45), 
('Mantalongon', 45), 
('Poblacion', 45), 
('San Juan', 45), 
('Santo Niño', 45),
('Tabuelan Proper', 45);

-- Tuburan Barangays
INSERT INTO barangays (name, municipality_id) VALUES
('Baclayon', 46), 
('Balayniwang', 46), 
('Binaliw', 46), 
('Bunga', 46), 
('Cabay', 46),
('Cagawasan', 46), 
('Cahumayan', 46), 
('Cansalonoy', 46), 
('Catunggan', 46), 
('Cawag', 46),
('Gua-an', 46), 
('Poblacion', 46), 
('San Juan', 46), 
('San Vicente', 46);

INSERT INTO users (first_name, last_name, email, password, role_id, address, phone_number) VALUES
('Ayuphone', 'Super Admin', 'ayuphoneph@gmail.com', '$2a$14$N3S0JKnY2UBkk9pk0mX3B./IvQo1abuEQPYXnEv2Xnlq.lVbxn00y', 1, 'Cebu City', '09123456789');

INSERT INTO users (first_name, last_name, email, password, role_id, address, phone_number) VALUES
('Ayuphone', 'Admin', 'ayuphoneph_admin@gmail.com', '$2a$14$N3S0JKnY2UBkk9pk0mX3B./IvQo1abuEQPYXnEv2Xnlq.lVbxn00y', 2, 'Cebu City', '09123456781');

INSERT INTO users (first_name, last_name, email, password, role_id, address, phone_number) VALUES
('Ayuphone', 'Customer', 'ayuphoneph_customer@gmail.com', '$2a$14$N3S0JKnY2UBkk9pk0mX3B./IvQo1abuEQPYXnEv2Xnlq.lVbxn00y', 3, 'Cebu City', '09123456782');

INSERT INTO users (first_name, last_name, email, password, role_id, address, phone_number) VALUES
('Ayuphone', 'Technician', 'ayuphoneph_tech@gmail.com', '$2a$14$N3S0JKnY2UBkk9pk0mX3B./IvQo1abuEQPYXnEv2Xnlq.lVbxn00y', 4, 'Cebu City', '09123456783');

-- Insert roles
INSERT INTO roles (name) VALUES
('superadmin'),
('admin'),
('customer'),
('technician');