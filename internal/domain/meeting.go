package domain

import "time"

type Meeting struct { //TODO internal model
	Id             int64
	PsychologistId int64
	ClientId       int64
	MeetingDate    time.Time
	StartTime      float64
	EndTime        float64
	Status         string
}

/*
CREATE TABLE IF NOT EXISTS meeting
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    psychologist_id bigint NOT NULL,
    client_id bigint NOT NULL,
    meeting_date date NOT NULL,
    start_time double precision,
    end_time double precision,
    status character varying(30) NOT NULL DEFAULT 'not completed',
	CONSTRAINT status_permissible_value CHECK (status = ANY (ARRAY['not completed', 'successfully completed', 'unsuccessfully completed'])),
    CONSTRAINT meeting_client_id_fkey FOREIGN KEY (client_id)
        REFERENCES client (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT meeting_psychologist_id_fkey FOREIGN KEY (psychologist_id)
        REFERENCES psychologist (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
)
*/
