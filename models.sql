CREATE TABLE IF NOT EXISTS customers (
    customer_id serial NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    date_of_birth date NOT NULL,
    city VARCHAR(100) NOT NULL,
    zipcode VARCHAR(10) NOT NULL,
    status int NOT NULL, 
    CONSTRAINT pk_customers PRIMARY KEY(customer_id)
);

INSERT INTO customers VALUES (1, 'Martin', '1989-09-08', 'Parana', '3100', 1)  ON CONFLICT DO NOTHING;
INSERT INTO customers VALUES (2, 'Nicole', '2001-01-01', 'Parana', '3100', 1)  ON CONFLICT DO NOTHING;


