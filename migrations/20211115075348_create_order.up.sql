CREATE TABLE orders (
    id serial not null primary key,
    orderUID varchar not null,
    entr varchar not null,
    totalprice int not null,
    customerid varchar not null,
    tracknumber varchar not null,
    deliveryservice varchar not null
);