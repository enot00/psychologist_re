create table users (
                               id serial primary key,
                               phone varchar(255) UNIQUE not null,
                               email varchar(255) default NULL,
                               password varchar(255) NOT NULL,
                               first_name varchar(255) NOT NULL,
                               second_name varchar(255) default NULL,
                               last_name varchar(255) NOT NULL,
                               description varchar(255) default NULL,
                               avatar varchar(255) default NULL,
                               role smallint default 0,
                               created_at timestamp not null default now(),
                               deleted_at timestamp default NULL,
                               CONSTRAINT correct_type CHECK ( role = ANY (Array [0, 1]))
);

create table specializations (
                                 id serial primary key,
                                 name varchar(255) UNIQUE not null
);

create table working_hours (
                               id serial primary key,
                               psychologist_id int NOT NULL ,
                               week_day smallint default NULL,
                               date date default NULL,
                               start_time timestamp NOT NULL ,
                               end_time timestamp NOT NULL,
                               FOREIGN KEY (psychologist_id)
                                   REFERENCES users(id),
                               CONSTRAINT end_gt_start CHECK ( end_time > start_time ),
                               CONSTRAINT correct_week_day CHECK ( week_day = ANY (Array [0, 1, 2, 3, 4, 5, 6])),
                               CONSTRAINT working_hours_check CHECK (( date is not null OR week_day is not  null) AND ( date is null OR week_day is null))
);

insert into users
(phone, email, password, first_name, second_name, last_name, description, avatar, role)
VALUES
    ('+380993575966', 'i.ivanov@gmail.com', '$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Ivan', '', 'Ivanov', '', 'photo1', 0),
    ('+380993575977', 's.stepanov@gmail.com','$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Stepan', '', 'Stepanov', 'I`m the best psychologist in the world', 'photo2', 1),
    ('+380993575988', 'w.widorov@gmail.com','$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Sidor', '', 'Sidorov', '',  'photo3', 0),
    ('+380993575999', '','8d962324559a3a6', 'Didor', 'Sidorovich', 'Sidorov', 'I`m the best psychologist in the world',  'photo4', 1),
    ('+380993575900', 'r.ridorov@gmail.com','$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Sidor', '', 'Sidorov', '',  'photo5', 0),
    ('+380993575911', '','$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Sidor', 'Sidorovich', 'Sidorov', 'I`m the best psychologist in the world',  'photo6', 1),
    ('+380993575922', '','$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Sidor', '', 'Sidorov', '',  'photo7', 0),
    ('+380993575933', '','$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Sidor', 'Sidorovich', 'Sidorov', 'I`m the best psychologist in the world',  'photo8', 1),
    ('+380993575944', 'v.vladimirov@gmail.com','$2a$14$.WqZoDNi0ubDWecBvZ6gzOpgJuFDZqBIhXoZ9Trcm6xz0B6guAXvu', 'Vladimir', '', 'Vladimirov', 'I`m the best psychologist in the world', 'photo9', 1);

insert into specializations (name)
values
    ('Загальна психологiя'),
    ('Соцiальна психологiя'),
    ('Гендерна психологiя'),
    ('Бiологiчна психологiя'),
    ('Юридична психологiя'),
    ('Педагогiчна психологiя');

create table users_specializations
(
    psychologist_id int NOT NULL,
    specialization_id int not null,
    primary key (psychologist_id, specialization_id),
    FOREIGN KEY (psychologist_id)
        REFERENCES users (id),
    FOREIGN KEY (specialization_id)
        REFERENCES specializations (id)
);
insert into users_specializations (psychologist_id, specialization_id)
VALUES
    (2, 1),
    (4, 2),
    (8, 3),
    (6, 4),
    (9, 5),
    (2, 6),
    (9, 1),
    (6, 2),
    (4, 3),
    (6, 6),
    (4, 5);

insert into working_hours (psychologist_id, week_day, start_time, end_time)
VALUES
    (2, 1,'0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z'),
    (2, 2,'0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z'),
    (2, 3,'0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z'),
    (2, 4,'0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z'),
    (2, 5,'0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z');
insert into working_hours (psychologist_id, date, start_time, end_time)
VALUES
    (4,'2022/06/27','0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z'),
    (6,'2022/06/27','0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z'),
    (9,'2022/06/27','0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z'),
    (8,'2022/06/27','0001-01-01T09:00:00Z', '0001-01-01T18:30:00Z');