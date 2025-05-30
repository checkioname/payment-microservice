-- Table creation to be executed after connecting to the ORDERS database
CREATE TABLE IF NOT EXISTS ORDERS (
                                      ID INTEGER PRIMARY KEY,
                                      CLIENT_ID INTEGER NOT NULL,
                                      CREATION_DATE TEXT NOT NULL,
                                      STATUS TEXT NOT NULL, -- E.g.: 'pending', 'payment_approved', 'in_preparation', 'invoice_issued', 'shipped', 'canceled'
                                      TOTAL_AMOUNT REAL
);


-- Table creation to be executed after connecting to the PAYMENT database
CREATE TABLE IF NOT EXISTS PAYMENT (
                                       ID INTEGER PRIMARY KEY,
                                       ORDER_ID INTEGER NOT NULL UNIQUE,    -- Usually 1 payment per order
                                       PROCESSING_DATE TEXT NOT NULL,
                                       STATUS TEXT NOT NULL,                -- E.g.: 'approved', 'rejected', 'pending'
                                       METHOD TEXT
);


-- Table creation to be executed after connecting to the PRODUCTS database
CREATE TABLE IF NOT EXISTS PRODUCTS (
                                        ID INTEGER PRIMARY KEY,
                                        NAME TEXT NOT NULL,
                                        PRICE REAL NOT NULL,
                                        STOCK_QUANTITY INTEGER NOT NULL
);


-- Table creation to be executed after connecting to the INVOICES database
CREATE TABLE IF NOT EXISTS INVOICES (
                                        ID INTEGER PRIMARY KEY,
                                        ORDER_ID INTEGER NOT NULL UNIQUE,
                                        NUMBER TEXT NOT NULL UNIQUE,
                                        ISSUE_DATE TEXT NOT NULL,
                                        ACCESS_KEY TEXT,
                                        FOREIGN KEY (ORDER_ID) REFERENCES ORDERS (ID)
    );



-- Table creation to be executed after connecting to the SHIPMENTS database
CREATE TABLE IF NOT EXISTS SHIPMENTS (
                                         ID INTEGER PRIMARY KEY,
                                         ORDER_ID INTEGER NOT NULL UNIQUE,
                                         INVOICE_ID INTEGER NOT NULL UNIQUE,
                                         SHIPMENT_DATE TEXT,
                                         TRACKING_CODE TEXT,
                                         STATUS TEXT NOT NULL, -- E.g.: 'awaiting_shipment', 'shipped', 'delivered'
                                         FOREIGN KEY (ORDER_ID) REFERENCES ORDERS (ID),
    FOREIGN KEY (INVOICE_ID) REFERENCES INVOICES (ID)
    );



------

-- inserts de exemplo
INSERT INTO payment (id, order_id, processing_date, status, method)
VALUES
    (1, 1, '2024-06-16', 'approved', 'credit_card'),
    (2, 2, '2024-06-16', 'approved', 'boleto'),
    (3, 3, '2024-06-16', 'pending', 'pix');