create table if not exists tokens_blacklist
(
    id         serial primary key,
    token      text        not null,
    expires_at timestamptz not null
);

create table if not exists refresh_tokens
(
    id         serial primary key,
    user_id    int         not null,
    token      text        not null,
    issued_at timestamptz not null,
    expires_at timestamptz not null,
    foreign key (user_id)
        references users (id) on delete cascade
);
