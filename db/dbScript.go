package db

var script = `
CREATE TABLE IF NOT EXISTS  categories
(
    category_id   serial PRIMARY KEY,
    category_name text                              NOT NULL
);

/*CREATE TABLE IF NOT EXISTS manufacturers
(
    manufacturer_id INTEGER PRIMARY KEY serial NOT NULL,
    manufacturer_name TEXT NOT NULL
);*/

CREATE TABLE IF NOT EXISTS products
(
    product_id   serial PRIMARY KEY  ,
    product_name text                              NOT NULL,
    description text NOT NULL,
    /*manufacturer_id INTEGER  NOT NULL,*/
    category_id  INTEGER,
    img_src text NOT NULL,
    diameter text Not NULL,
    height integer not null
/*,
    FOREIGN KEY ([manufacturer_id]) REFERENCES "manufacturers" ([manufacturer_id]) ON DELETE NO ACTION ON UPDATE NO ACTION*/
);

CREATE TABLE IF NOT EXISTS price_change
(
    product_id        INTEGER NOT NULL REFERENCES products (product_id) ON DELETE CASCADE,
    date_price_change integer NOT NULL,
    new_price         REAL    NOT NULL,
    CONSTRAINT PK_PRICE_CHANGE PRIMARY KEY (product_id, date_price_change)
);

CREATE TABLE IF NOT EXISTS stores
(
    store_id   serial PRIMARY KEY ,
    store_name text                              NOT NULL
);

/* Поставки */
CREATE TABLE IF NOT EXISTS deliveries
(
    product_id    INTEGER NOT NULL REFERENCES products (product_id),
    store_id      INTEGER NOT NULL REFERENCES stores (store_id),
    delivery_date INTEGER NOT NULL,
    product_count INTEGER NOT NULL
);

/*Клиенты*/
CREATE TABLE IF NOT EXISTS customers
(
    customer_id       serial PRIMARY KEY,
    customer_name     text                              NOT NULL,
    customer_password TEXT                              NOT NULL,
    customer_city    TEXT,
    customer_phone    TEXT UNIQUE                       NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY,
    created INTEGER NOT NULL,
    customer_id INTEGER NOT NULL REFERENCES customers (customer_id)
);

/* Покупки */
CREATE TABLE IF NOT EXISTS purchases
(
    purchase_id   serial PRIMARY KEY,
    customer_id   INTEGER                           NOT NULL REFERENCES customers (customer_id),
    store_id      INTEGER                           NOT NULL REFERENCES stores (store_id),
    purchase_date INTEGER                           NOT NULL
);

CREATE TABLE IF NOT EXISTS purchase_items
(
    purchase_id   INTEGER NOT NULL REFERENCES purchases (purchase_id),
    product_id    INTEGER NOT NULL REFERENCES products (product_id),
    product_count INTEGER NOT NULL,
    product_price REAL    NOT NULL,
    CONSTRAINT PK_PURCHASE_ITEMS PRIMARY KEY (purchase_id, product_id)
);

CREATE TABLE IF NOT EXISTS Files
(_id serial PRIMARY KEY,
 FileName TEXT NOT NULL,
 ImageData BYTEA)


`
