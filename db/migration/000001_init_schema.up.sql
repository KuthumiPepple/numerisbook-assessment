CREATE TABLE "invoices" (
  "invoice_number" bigserial PRIMARY KEY,
  "customer_name" varchar NOT NULL,
  "customer_email" varchar NOT NULL,
  "customer_phone" varchar NOT NULL,
  "customer_address" varchar NOT NULL,
  "sender_name" varchar NOT NULL,
  "sender_email" varchar NOT NULL,
  "sender_phone" varchar NOT NULL,
  "sender_address" varchar NOT NULL,
  "issue_date" timestamptz NOT NULL DEFAULT (now()),
  "due_date" timestamptz NOT NULL DEFAULT (now() + interval '30 days'),
  "status" varchar NOT NULL DEFAULT 'draft',
  "subtotal" varchar NOT NULL,
  "discount_rate" varchar NOT NULL,
  "discount" varchar NOT NULL,
  "total_amount" varchar NOT NULL,
  "payment_info" varchar NOT NULL,
  "billing_currency" varchar NOT NULL DEFAULT 'USD',
  "note" varchar NOT NULL DEFAULT 'Thank you for your patronage'
);

CREATE TABLE "line_items" (
  "id" bigserial PRIMARY KEY,
  "invoice_number" bigint NOT NULL,
  "description" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "unit_price" varchar NOT NULL,
  "total_price" varchar NOT NULL
);

CREATE INDEX ON "invoices" ("status");

CREATE INDEX ON "line_items" ("invoice_number");

ALTER TABLE "line_items" ADD FOREIGN KEY ("invoice_number") REFERENCES "invoices" ("invoice_number");