CREATE TABLE IF NOT EXISTS search_history(
    id SERIAL PRIMARY KEY,
    query TEXT
);

CREATE TABLE IF NOT EXISTS address(
    id SERIAL PRIMARY KEY,
    lat TEXT,
    lon TEXT
);

CREATE TABLE IF NOT EXISTS history_search_address(
    id SERIAL PRIMARY KEY,
    search_id INT,
    address_id INT,
    FOREIGN KEY (search_id) REFERENCES search_history(id),
    FOREIGN KEY (address_id) REFERENCES address(id)
);