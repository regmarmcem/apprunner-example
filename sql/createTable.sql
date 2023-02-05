create table if not exists tasks (
    task_id integer unsigned auto_increment primary key,
    title varchar(100) not null,
    details text not null,
    created_at datetime
);