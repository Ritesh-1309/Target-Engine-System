use adserver;

SHOW TABLES;

CREATE TABLE campaigns (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    cta VARCHAR(100),
    status ENUM('ACTIVE', 'INACTIVE') NOT NULL
);

CREATE TABLE targeting_rules (
    campaign_id VARCHAR(50),
    include_apps JSON DEFAULT NULL,
    exclude_apps JSON DEFAULT NULL,
    include_countries JSON DEFAULT NULL,
    exclude_countries JSON DEFAULT NULL,
    include_os JSON DEFAULT NULL,
    exclude_os JSON DEFAULT NULL,
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
);

INSERT INTO campaigns (id, name, image, cta, status) VALUES
('spotify', 'Spotify - Music for everyone', 'https://somelink', 'Download', 'ACTIVE'),
('duolingo', 'Duolingo: Best way to learn', 'https://somelink2', 'Install', 'ACTIVE'),
('subwaysurfer', 'Subway Surfer', 'https://somelink3', 'Play', 'ACTIVE');

INSERT INTO targeting_rules (
    campaign_id, include_apps, exclude_apps,
    include_countries, exclude_countries,
    include_os, exclude_os
) VALUES
('spotify', NULL, NULL, '["us", "canada"]', NULL, NULL, NULL),
('duolingo', NULL, NULL, NULL, '["us"]', '["android", "ios"]', NULL),
('subwaysurfer', '["com.gametion.ludokinggame"]', NULL, NULL, NULL, '["android"]', NULL);

select * from campaigns;
select * from targeting_rules;