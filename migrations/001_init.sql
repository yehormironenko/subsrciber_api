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
                                          wallet_address VARCHAR(255) NOT NULL,
                                          email_notifications BOOLEAN DEFAULT TRUE,
                                          websocket_notifications BOOLEAN DEFAULT TRUE,
                                          created_at TIMESTAMP DEFAULT NOW(),
                                          CONSTRAINT fk_wallet_address FOREIGN KEY (user_id, wallet_address)
                                              REFERENCES wallet_subscriptions(user_id, wallet_address) ON DELETE CASCADE,
                                          CONSTRAINT unique_user_wallet_notification UNIQUE (user_id, wallet_address)
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

CREATE OR REPLACE FUNCTION create_notification_preferences_for_wallet()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO notification_preferences (user_id, wallet_address, email_notifications, websocket_notifications)
    VALUES (NEW.user_id, NEW.wallet_address, FALSE, FALSE);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_create_wallet_notification_preferences
    AFTER INSERT ON wallet_subscriptions
    FOR EACH ROW
EXECUTE FUNCTION create_notification_preferences_for_wallet();