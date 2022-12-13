CREATE TABLE ipconfigs (
   ip varchar(255) not null primary key,
   hostfqdn varchar(255) not null,
   active char(1) not null
);
/*
migrate -path db/migration -database "mysql://shiprocket_api:DDwgLAA3WH2a2k1h@tcp(127.0.0.1:3309)/wms" up
*/