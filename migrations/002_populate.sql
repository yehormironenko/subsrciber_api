INSERT INTO users (email) VALUES
                              ('user1@example.com'),
                              ('user2@example.com');

INSERT INTO wallet_subscriptions (user_id, wallet_address) VALUES
                                                               (1, '0xAbC123...'),
                                                               (1, '0xEfG456...'),
                                                               (2, '0xH1J789...');

-- Notification preferences are automatically created by trigger when wallet_subscriptions are inserted
-- Update them to have custom test values
UPDATE notification_preferences SET email_notifications = TRUE, websocket_notifications = TRUE
WHERE user_id = 1 AND wallet_address = '0xAbC123...';

UPDATE notification_preferences SET email_notifications = TRUE, websocket_notifications = FALSE
WHERE user_id = 1 AND wallet_address = '0xEfG456...';

UPDATE notification_preferences SET email_notifications = TRUE, websocket_notifications = FALSE
WHERE user_id = 2 AND wallet_address = '0xH1J789...';

INSERT INTO transactions (wallet_address, tx_hash, amount, timestamp) VALUES
                                                                          ('0xAbC123...', '0xTxHash123', 0.5, NOW()),
                                                                          ('0xEfG456...', '0xTxHash456', 1.2, NOW()),
                                                                          ('0xH1J789...', '0xTxHash789', 2.8, NOW());