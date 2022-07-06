CREATE TABLE users
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO users (name)
VALUES ('u1'),
       ('u2'),
       ('u3'),
       ('u4');

CREATE TABLE chats
(
    id     SERIAL PRIMARY KEY CHECK ( id > 0 ),
    user_id INT NOT NULL,
    participants INT[2] NOT NULL UNIQUE,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
);

INSERT INTO chats (user_id, participants)
VALUES (1, '{1,2}'),
       (1, '{2,1}'),
       (2, '{2,3}'),
       (3, '{3,4}');


ALTER SEQUENCE chats_id_seq RESTART WITH 3;


CREATE TABLE messages
(
    id      SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    from_id INT NOT NULL,
    text    VARCHAR(255),
    file_path VARCHAR(255) NULL,
    date    TIMESTAMP,
    FOREIGN KEY (chat_id)
        REFERENCES chats (id) ON DELETE CASCADE,
    FOREIGN KEY (from_id)
        REFERENCES users (id) ON DELETE SET NULL
);
