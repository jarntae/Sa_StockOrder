BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "categories" (
	"category_id"	text,
	"category_name"	text,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	PRIMARY KEY("category_id")
);
CREATE TABLE IF NOT EXISTS "genders" (
	"gender_id"	text,
	"gender_name"	text,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	PRIMARY KEY("gender_id")
);
CREATE TABLE IF NOT EXISTS "positions" (
	"position_id"	text,
	"position_name"	text,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	PRIMARY KEY("position_id")
);
CREATE TABLE IF NOT EXISTS "employees" (
	"employee_id"	text,
	"first_name"	text,
	"last_name"	text,
	"username"	text,
	"password"	text,
	"register_date"	datetime,
	"gender_id"	text,
	"position_id"	text,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	PRIMARY KEY("employee_id"),
	CONSTRAINT "fk_positions_employee" FOREIGN KEY("position_id") REFERENCES "positions"("position_id"),
	CONSTRAINT "fk_genders_employee" FOREIGN KEY("gender_id") REFERENCES "genders"("gender_id")
);
CREATE TABLE IF NOT EXISTS "products" (
	"product_id"	text,
	"product_name"	text,
	"category_id"	text,
	"employee_id"	text,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	PRIMARY KEY("product_id"),
	CONSTRAINT "fk_categories_product" FOREIGN KEY("category_id") REFERENCES "categories"("category_id"),
	CONSTRAINT "fk_employees_product" FOREIGN KEY("employee_id") REFERENCES "employees"("employee_id")
);
CREATE TABLE IF NOT EXISTS "suppliers" (
	"supplier_id"	integer,
	"supplier_name"	text,
	"phone"	text,
	"email"	text,
	"address"	text,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	PRIMARY KEY("supplier_id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "stocks" (
	"stock_id"	integer,
	"quantity"	integer,
	"price"	real,
	"date_in"	datetime,
	"expiration_date"	datetime,
	"product_id"	text,
	"employee_id"	text,
	"supplier_id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	PRIMARY KEY("stock_id" AUTOINCREMENT),
	CONSTRAINT "fk_products_stock" FOREIGN KEY("product_id") REFERENCES "products"("product_id"),
	CONSTRAINT "fk_employees_stock" FOREIGN KEY("employee_id") REFERENCES "employees"("employee_id"),
	CONSTRAINT "fk_suppliers_stock" FOREIGN KEY("supplier_id") REFERENCES "suppliers"("supplier_id")
);
CREATE INDEX IF NOT EXISTS "idx_categories_deleted_at" ON "categories" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_genders_deleted_at" ON "genders" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_positions_deleted_at" ON "positions" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_employees_deleted_at" ON "employees" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_products_deleted_at" ON "products" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_suppliers_deleted_at" ON "suppliers" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_stocks_deleted_at" ON "stocks" (
	"deleted_at"
);
COMMIT;
