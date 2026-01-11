ALTER TABLE inri_clients ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_clients ALTER COLUMN description DROP NOT NULL;
ALTER TABLE inri_clients ALTER COLUMN tax_number DROP NOT NULL;
ALTER TABLE inri_clients ALTER COLUMN address DROP NOT NULL;
ALTER TABLE inri_projects ALTER COLUMN description DROP NOT NULL;
ALTER TABLE inri_invoice ALTER COLUMN external_id DROP NOT NULL;
ALTER TABLE inri_invoice ALTER COLUMN external_link DROP NOT NULL;
ALTER TABLE inri_invoice ALTER COLUMN note DROP NOT NULL;
ALTER TABLE inri_invoice RENAME COLUMN note TO description;

