CREATE DATABASE blockchain_notifications;
\c blockchain_notifications;
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE wallet_subscriptions (
                                      id SERIAL PRIMARY KEY,
                                      user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                      wallet_address VARCHAR(255) NOT NULL,
                                      created_at TIMESTAMP DEFAULT NOW(),
                                      UNIQUE(user_id, wallet_address)
);

CREATE TABLE notification_preferences (
                                          id SERIAL PRIMARY KEY,
                                          user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                          email_notifications BOOLEAN DEFAULT TRUE,
                                          websocket_notifications BOOLEAN DEFAULT TRUE,
                                          created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
                              id SERIAL PRIMARY KEY,
                              wallet_address VARCHAR(255) NOT NULL,
                              tx_hash VARCHAR(255) UNIQUE NOT NULL,
                              amount NUMERIC(18,8),
                              timestamp TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_wallet_address ON wallet_subscriptions(wallet_address);
CREATE INDEX idx_tx_hash ON transactions(tx_hash);