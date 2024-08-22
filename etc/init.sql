CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NUll,
    login VARCHAR(100) NOT NULL,
    password BYTEA NOT NUll,
    haspremium BOOLEAN DEFAULT FALSE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE users IS 'Таблица с информацией по пользователям';

CREATE TABLE IF NOT EXISTS tokens (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) NOT NULL,
    token bytea NOT NULL,
    refresh bytea NOT NULL,
    expires_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE tokens IS 'Таблица с информацией по выданным токенам по результату авторизации';

CREATE TABLE IF NOT EXISTS roles (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE roles IS 'Таблица с информацией по ролям пользователей системы';

CREATE TABLE IF NOT EXISTS permissions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE permissions IS 'Таблица с правами пользователей';

COMMENT ON COLUMN permissions.name IS 'Часть маршрута обращения к endpoint-у сервера, на которую выдается право доступа';

CREATE TABLE IF NOT EXISTS users_roles (
    user_id BIGINT REFERENCES users(id) NOT NULL,
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

COMMENT ON TABLE users_roles IS 'Таблица СВЯЗЕЙ пользователей и ролей';

CREATE TABLE IF NOT EXISTS roles_permissions (
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    permission_id BIGINT REFERENCES permissions(id) NOT NULL,
    PRIMARY KEY (role_id, permission_id)
);

COMMENT ON TABLE roles_permissions IS 'Таблица СВЯЗЕЙ ролей и прав пользователей';

CREATE TABLE IF NOT EXISTS products (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL CHECK(price > 0),
    description VARCHAR(500),
    quantity INTEGER NOT NULL DEFAULT 0 CHECK(quantity >= 0),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE products IS 'Таблица с информацией по товарам';

CREATE TABLE IF NOT EXISTS product_categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    discount_category INTEGER CHECK (
        discount_category >= 0 
        AND discount_category < 100
    ) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE products IS 'Таблица с информацией по товарам';

CREATE TABLE IF NOT EXISTS product_discounts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    discount_premium INTEGER CHECK(
        discount_premium >= 0
        AND discount_premium < 100
    ) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE product_discounts IS 'Таблица с информацией по скидкам на товары/категории товаров';

CREATE TABLE IF NOT EXISTS carts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    totalprice_products BIGINT CHECK(totalprice_products >= 0),
    discount INTEGER CHECK(
        discount >= 0
        AND discount < 100
    ),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE carts IS 'Таблица с информацией по корзине пользователя';

CREATE TABLE IF NOT EXISTS orders (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    is_paid BOOLEAN NOT NULL DEFAULT FALSE,
    address VARCHAR(100) NOT NULL,
    totalprice_products BIGINT NOT NUll,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE orders IS 'Таблица с информацией по заказам пользователя';

CREATE TABLE IF NOT EXISTS carts_users (
    user_id BIGINT REFERENCES users(id) NOT NULL,
    cart_id BIGINT REFERENCES carts(id) NOT NULL,
    PRIMARY KEY(user_id, cart_id)
);

COMMENT ON TABLE carts_users IS 'Таблица связей корзины и пользователя';

CREATE TABLE IF NOT EXISTS products_product_discounts (
    product_id BIGINT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    product_discount_id BIGINT REFERENCES product_discounts(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY (product_id, product_discount_id)
);

COMMENT ON TABLE products_product_discounts IS 'Таблица связей продукта и персональной скидки';

CREATE TABLE IF NOT EXISTS products_product_categories (
    product_id BIGINT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    product_categories_id BIGINT REFERENCES product_categories(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY (product_id, product_categories_id)
);

COMMENT ON TABLE products_product_categories IS 'Таблица связей продукта и категории';

CREATE TABLE IF NOT EXISTS carts_products (
    cart_id BIGINT REFERENCES carts(id) ON DELETE CASCADE NOT NULL,
    product_id BIGINT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(cart_id, product_id),
    discount INTEGER CHECK (discount >= 0 AND discount < 100) NOT NULL,
    quantity_product BIGINT CHECK(quantity_product >= 0) NOT NULL,
    total_product_price BIGINT CHECK(total_product_price >= 0) NOT NULL
);

CREATE TABLE IF NOT EXISTS users_orders (
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    order_id BIGINT REFERENCES orders(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY (user_id, order_id)
);

CREATE TABLE IF NOT EXISTS orders_products (
    order_id BIGINT REFERENCES orders(id) ON DELETE CASCADE NOT NULL,
    product_id BIGINT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY (order_id, product_id),
    quantity_product BIGINT CHECK(quantity_product >= 0) NOT NULL,
    total_product_price BIGINT CHECK(total_product_price >= 0) NOT NULL
);

CREATE TABLE IF NOT EXISTS carts_histories (
    cart_id BIGINT REFERENCES carts(id) ON DELETE CASCADE NOT NULL,
    product_id BIGINT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY (cart_id, product_id)
);

INSERT INTO permissions(name) VALUES('GET product');
INSERT INTO permissions(name) VALUES('DELETE product');
INSERT INTO permissions(name) VALUES('GET products');
INSERT INTO permissions(name) VALUES('POST products');
INSERT INTO permissions(name) VALUES('PUT products');
INSERT INTO permissions(name) VALUES('GET user');
INSERT INTO permissions(name) VALUES('POST pay');
INSERT INTO permissions(name) VALUES('PUT cart');
INSERT INTO permissions(name) VALUES('GET cart');
INSERT INTO permissions(name) VALUES('POST order');
INSERT INTO permissions(name) VALUES('GET order');
INSERT INTO permissions(name) VALUES('DELETE order');
INSERT INTO permissions(name) VALUES('GET orders');

INSERT INTO roles(name) VALUES('admin');
INSERT INTO roles(name) VALUES('user');

--admin permissions
INSERT INTO roles_permissions VALUES(1, 1);
INSERT INTO roles_permissions VALUES(1, 2);
INSERT INTO roles_permissions VALUES(1, 3);
INSERT INTO roles_permissions VALUES(1, 4);
INSERT INTO roles_permissions VALUES(1, 5);
INSERT INTO roles_permissions VALUES(1, 6);
INSERT INTO roles_permissions VALUES(1, 7);
INSERT INTO roles_permissions VALUES(1, 8);
INSERT INTO roles_permissions VALUES(1, 9);
INSERT INTO roles_permissions VALUES(1, 10);
INSERT INTO roles_permissions VALUES(1, 11);
INSERT INTO roles_permissions VALUES(1, 12);
INSERT INTO roles_permissions VALUES(1, 13);

--user permissions
INSERT INTO roles_permissions VALUES(2, 1);
INSERT INTO roles_permissions VALUES(2, 3);
INSERT INTO roles_permissions VALUES(2, 6);
INSERT INTO roles_permissions VALUES(2, 7);
INSERT INTO roles_permissions VALUES(2, 8);
INSERT INTO roles_permissions VALUES(2, 9);
INSERT INTO roles_permissions VALUES(2, 10);
INSERT INTO roles_permissions VALUES(2, 11);
INSERT INTO roles_permissions VALUES(2, 12);
INSERT INTO roles_permissions VALUES(2, 13);

--
INSERT INTO users(firstname, lastname, email, login, password) VALUES('admin','admin', 'admin@k.ru', 'admin', 'admin');
INSERT INTO users_roles VALUES(1, 1);

-- test data
INSERT INTO products(name, price) VALUES ('product_1', 100);
INSERT INTO products(name, price) VALUES ('product_2', 200);
INSERT INTO products(name, price) VALUES ('product_4', 300);

INSERT INTO product_discounts(discount_premium) VALUES(10);
INSERT INTO product_discounts(discount_premium) VALUES(20);
INSERT INTO product_discounts(discount_premium) VALUES(30);

INSERT INTO products_product_discounts VALUES(1, 1);
INSERT INTO products_product_discounts VALUES(2, 2);
INSERT INTO products_product_discounts VALUES(3, 3);

INSERT INTO product_categories(name, discount_category) VALUES ('category_1', 5);
INSERT INTO product_categories(name, discount_category) VALUES ('category_2', 6);
INSERT INTO product_categories(name, discount_category) VALUES ('category_3', 7);

INSERT INTO products_product_categories VALUES(1, 1);
INSERT INTO products_product_categories VALUES(2, 2);
INSERT INTO products_product_categories VALUES(3, 3);

--cart
INSERT INTO carts(totalprice_products, discount) VALUES (0, 0);
INSERT INTO carts_users VALUES(1, 1);
