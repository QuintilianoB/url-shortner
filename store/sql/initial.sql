create table if not exists urls
(
    id        int auto_increment primary key,
    url       text not null,
    url_short text not null
);

create table if not exists statistics
(
    id     int auto_increment primary key,
    id_url int     not null,
    time   int(64) not null,
    constraint statistics_fk_urls
        foreign key (id_url) references urls (id)
);