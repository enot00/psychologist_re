package domain

import "time"

type Client struct { //TODO internal model
	Id               int64
	UserName         string
	PhoneNumber      string
	Email            string
	Avatar           string
	OldPassword      string // for PUT request
	Password         string
	RegistrationDate time.Time
}

/*
CREATE TABLE IF NOT EXISTS client
(
    id bigint NOT NULL DEFAULT nextval('client_id_seq'::regclass) PRIMARY KEY,
    user_name character varying(50) NOT NULL UNIQUE,
    phone_number character varying(13) NOT NULL UNIQUE,
    email character varying,
    avatar character varying,
	password character varying NOT NULL,
    registration_date date NOT NULL
)
*/
