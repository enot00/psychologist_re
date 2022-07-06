CREATE TABLE IF NOT EXISTS meeting
(
    id serial NOT NULL PRIMARY KEY,
    psychologist_id bigint NOT NULL,
    client_id bigint NOT NULL,
    meeting_date date NOT NULL,
    start_time float,
    end_time float,
    status character varying(30) NOT NULL DEFAULT 'not completed',
    CONSTRAINT meeting_client_id_fkey FOREIGN KEY (client_id)
        REFERENCES client (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT meeting_psychologist_id_fkey FOREIGN KEY (psychologist_id)
        REFERENCES psychologist (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT status_permissible_value CHECK (status = ANY (ARRAY['not completed', 'successfully completed', 'unsuccessfully completed']))
);

INSERT INTO meeting(psychologist_id, client_id, meeting_date, start_time, end_time)
VALUES
    (1, 1, '2022-06-10', 10.00, 11.00),
 	(1, 2, '2022-06-21', 11.30, 12.30),
 	(1, 3, '2022-06-22', 12.00, 13.00);

INSERT INTO meeting(psychologist_id, client_id, meeting_date, start_time, end_time, status)
VALUES (1, 1, '2022-06-23', 10.00, 11.00, 'successfully completed');    