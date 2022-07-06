CREATE TABLE IF NOT EXISTS client
(
    id serial NOT NULL PRIMARY KEY,
    user_name character varying(50) NOT NULL UNIQUE,
    email character varying,
    avatar character varying,
    registration_date date NOT NULL,
    password character varying NOT NULL,
    phone_number character varying(13) NOT NULL UNIQUE
);

INSERT INTO client(user_name, email, avatar, registration_date, password, phone_number)
VALUES 
	('client1', 'client1@mail.net', 'https://mytest-avatars.s3.eu-north-1.amazonaws.com/ava1.png', NOW(), '$2a$14$DiFvOQ.XZNPMULrzSG0KXe3N1DZUiLwdqzaAylJUn4cJGjjaeHAs2', '+380982162334'),
	('client2', 'client2@mail.net', 'https://mytest-avatars.s3.eu-north-1.amazonaws.com/ava2.png', NOW(), '$2a$14$DiFvOQ.XZNPMULrzSG0KXe3N1DZUiLwdqzaAylJUn4cJGjjaeHAs2', '+380982162399'),
	('client3', 'client3@mail.net', 'https://mytest-avatars.s3.eu-north-1.amazonaws.com/ava3.webp', NOW(), '$2a$14$DiFvOQ.XZNPMULrzSG0KXe3N1DZUiLwdqzaAylJUn4cJGjjaeHAs2', '+380662362376');