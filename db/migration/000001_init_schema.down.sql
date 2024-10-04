ALTER TABLE "line_items" DROP CONSTRAINT "line_items_invoice_number_fkey";
ALTER TABLE "invoices" DROP CONSTRAINT "invoices_customer_id_fkey";
ALTER TABLE "invoices" DROP CONSTRAINT "invoices_vendor_id_fkey";

DROP TABLE IF EXISTS "vendors";
DROP TABLE IF EXISTS "customers";
DROP TABLE IF EXISTS "line_items";
DROP TABLE IF EXISTS "invoices";