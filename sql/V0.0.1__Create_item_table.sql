CREATE TABLE items
(
  id bigserial primary key,
  name text not null unique,
  description text null,
  count int not null
);



insert into items(name,description, count)
values (
    'Sword of stones',
    'Not that practical really',
    1
);