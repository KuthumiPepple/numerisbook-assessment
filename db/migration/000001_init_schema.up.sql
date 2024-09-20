CREATE TABLE "invoices" (
  "invoice_number" bigserial PRIMARY KEY,
  "customer_id" bigint NOT NULL,
  "vendor_id" bigint NOT NULL,
  "issue_date" timestamptz NOT NULL DEFAULT (now()),
  "due_date" timestamptz NOT NULL DEFAULT (now() + interval '30 days'),
  "status" varchar NOT NULL DEFAULT 'draft',
  "subtotal" bigint NOT NULL DEFAULT 0,
  "discount_rate" bigint NOT NULL DEFAULT 0,
  "discount" bigint NOT NULL DEFAULT 0,
  "total_amount" bigint NOT NULL DEFAULT 0,
  "billing_currency" varchar NOT NULL DEFAULT 'USD',
  "note" varchar NOT NULL DEFAULT 'Thank you for your patronage'
);

CREATE TABLE "line_items" (
  "id" bigserial PRIMARY KEY,
  "invoice_number" bigint NOT NULL,
  "description" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "unit_price" bigint NOT NULL,
  "total_price" bigint NOT NULL
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "address" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "vendors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "address" varchar NOT NULL,
  "email" varchar NOT NULL,
  "bank_account_name" varchar NOT NULL,
  "bank_account_no" bigint NOT NULL,
  "bank_name" varchar NOT NULL
);

ALTER TABLE "invoices" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "line_items" ADD FOREIGN KEY ("invoice_number") REFERENCES "invoices" ("invoice_number");