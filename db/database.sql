CREATE TABLE IF NOT EXISTS  categories
(
    category_id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    category_name text                              NOT NULL
);

/*CREATE TABLE IF NOT EXISTS manufacturers
(
    manufacturer_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    manufacturer_name TEXT NOT NULL
);*/

CREATE TABLE IF NOT EXISTS products
(
    product_id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    product_name text                              NOT NULL,
    description text NOT NULL,
    /*manufacturer_id INTEGER  NOT NULL,*/
    category_id  INTEGER                           ,
    img_src text NOT NULL,
    diameter text Not NULL,
    height integer not null,

    FOREIGN KEY ([category_id]) REFERENCES "categories" ([category_id]) ON DELETE NO ACTION ON UPDATE NO ACTION/*,
    FOREIGN KEY ([manufacturer_id]) REFERENCES "manufacturers" ([manufacturer_id]) ON DELETE NO ACTION ON UPDATE NO ACTION*/
);

CREATE TABLE IF NOT EXISTS price_change
(
    product_id        INTEGER NOT NULL,
    date_price_change integer NOT NULL,
    new_price         REAL    NOT NULL,
    CONSTRAINT PK_PRICE_CHANGE PRIMARY KEY (product_id, date_price_change),
    FOREIGN KEY ([product_id]) REFERENCES "products" ([product_id]) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS stores
(
    store_id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    store_name text                              NOT NULL
);

/* Поставки */
CREATE TABLE IF NOT EXISTS deliveries
(
    product_id    INTEGER NOT NULL,
    store_id      INTEGER NOT NULL,
    delivery_date INTEGER NOT NULL,
    product_count INTEGER NOT NULL,
    FOREIGN KEY ([product_id]) REFERENCES "products" ([product_id]) ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY ([store_id]) REFERENCES "stores" ([store_id]) ON DELETE NO ACTION ON UPDATE NO ACTION
);

/*Клиенты*/
CREATE TABLE IF NOT EXISTS customers
(
    customer_id       INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    customer_name     text                              NOT NULL,
    customer_password TEXT                              NOT NULL,
    customer_city    TEXT,
    customer_phone    TEXT UNIQUE                       NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY NOT NULL,
    created INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    FOREIGN KEY ([customer_id]) REFERENCES "customers" ([customer_id]) ON DELETE NO ACTION ON UPDATE NO ACTION
);

/* Покупки */
CREATE TABLE IF NOT EXISTS purchases
(
    purchase_id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    customer_id   INTEGER                           NOT NULL REFERENCES customers(customer_id),
    store_id      INTEGER                           NOT NULL REFERENCES stores(store_id),
    purchase_date INTEGER                           NOT NULL,
);

CREATE TABLE IF NOT EXISTS purchase_items
(
    purchase_id   INTEGER NOT NULL,
    product_id    INTEGER NOT NULL,
    product_count INTEGER NOT NULL,
    product_price REAL    NOT NULL,
    CONSTRAINT PK_PURCHASE_ITEMS PRIMARY KEY (purchase_id, product_id),
    FOREIGN KEY ([product_id]) REFERENCES "products" ([product_id]) ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY ([purchase_id]) REFERENCES "purchases" ([purchase_id]) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS Files
(_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
 FileName TEXT NOT NULL,
 ImageData BLOB)

