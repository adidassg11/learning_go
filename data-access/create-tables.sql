DROP TABLE IF EXISTS album;
CREATE TABLE album(  
    id          INT AUTO_INCREMENT NOT NULL COMMENT 'Primary Key',
    title       VARCHAR(128) NOT NULL COMMENT 'the album title',
    artist      VARCHAR(255) NOT NULL COMMENT 'the artist name', 
    price       DECIMAL(5, 2) NOT NULL COMMENT 'price of the album',
    PRIMARY KEY (`id`),
    KEY `title` (`title`) COMMENT 'adding this just for fun'
) COMMENT 'Album table comment';

INSERT INTO album (
    title, artist, price
)
VALUES
    ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);