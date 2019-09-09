


create database projectA charset utf8mb4;
use projectA;

create user projectAUser@localhost identified by 'hellokang';
grant all privileges on projectA.* to projectAUser@localhost;
flush privileges;

create table if not exists a_categories (
  id int unsigned auto_increment,
  parent_id int unsigned,
  name varchar(255),
  logo varchar(255),
  description varchar(255),
  sort_order int,
  meta_title varchar(255),
  meta_keywords varchar(255),
  meta_description varchar(255),
  primary key (id),
  index (parent_id),
  index (name),
  index (sort_order)
)engine innodb charset utf8mb4;

insert into a_categories (id, name, parent_id) values (1, '未分类', 0);