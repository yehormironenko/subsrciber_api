INSERT INTO users (email) VALUES
                              ('user1@example.com'),
                              ('user2@example.com');

INSERT INTO wallet_subscriptions (user_id, wallet_address) VALUES
                                                               (1, '0xAbC123...'),
                                                               (1, '0xEfG456...'),
                                                               (2, '0xH1J789...');

INSERT INTO notification_preferences (user_id, email_notifications, websocket_notifications) VALUES
                                                                                                 (1, TRUE, TRUE),
                                                                                                 (2, TRUE, FALSE);

INSERT INTO transactions (wallet_address, tx_hash, amount, timestamp) VALUES
                                                                          ('0xAbC123...', '0xTxHash123', 0.5, NOW()),
                                                                          ('0xEfG456...', '0xTxHash456', 1.2, NOW()),
                                                                          ('0xH1J789...', '0xTxHash789', 2.8, NOW());