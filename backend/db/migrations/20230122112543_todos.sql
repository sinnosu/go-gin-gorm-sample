-- migrate:up
create table todos (
  id integer not null AUTO_INCREMENT,
  name varchar(255) not null ,
  description text,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp,
  primary key (`id`)
) ENGINE=InnoDB;
-- migrate:down
drop table todos;