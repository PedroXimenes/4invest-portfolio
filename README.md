Link do Figma para o protótipo do sistema:
https://www.figma.com/file/BqSxCwiZsd1EFJuC7s6ZAc/4-invest?type=design&node-id=0%3A1&mode=design&t=TOOeJJii8n6vBgpI-1

database fourinvest
user admin4invest
password 4invest

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT,
    password TEXT,
    age INTEGER,
    name TEXT,
    nationality VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE portfolio (
    portfolio_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    name VARCHAR(255)
);

CREATE TABLE assets (
    portfolio_id INTEGER REFERENCES portfolio(portfolio_id),
    type VARCHAR(255),
    symbol VARCHAR(255),
    quantity INTEGER,
    max_value_limit REAL,
    ideal_percentage REAL
);

senha banco no gcp: Vah1>pdg{fp,%cvm

fourinvestadmin: g:Q=($U('h^+en6T