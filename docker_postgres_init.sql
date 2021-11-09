CREATE TABLE Customers
(
    id  uuid,
    name text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    age int NOT NULL,
    married bool,
    CONSTRAINT Customers_pkey PRIMARY KEY (id)
    
);

CREATE TABLE cats
(
    id  uuid,
    name text COLLATE pg_catalog."default",
    age int NOT NULL,
    color text COLLATE pg_catalog."default",
    CONSTRAINT cats_pkey PRIMARY KEY (id)
    
);