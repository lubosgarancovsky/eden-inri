ALTER TABLE inri_clients ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_clients ALTER COLUMN description DROP NOT NULL;
ALTER TABLE inri_clients ALTER COLUMN tax_number DROP NOT NULL;
ALTER TABLE inri_clients ALTER COLUMN address DROP NOT NULL;
ALTER TABLE inri_projects ALTER COLUMN description DROP NOT NULL;
ALTER TABLE inri_invoice ALTER COLUMN external_id DROP NOT NULL;
ALTER TABLE inri_invoice ALTER COLUMN external_link DROP NOT NULL;
ALTER TABLE inri_invoice ALTER COLUMN note DROP NOT NULL;
ALTER TABLE inri_invoice RENAME COLUMN note TO description;

ALTER TABLE inri_attachments ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_contact_person ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_kanban_boards ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_kanban_boards ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
ALTER TABLE inri_kanban_columns ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_labels ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
ALTER TABLE inri_labels ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_projects ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_story_activities ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_stories ADD COLUMN deleted_at TIMESTAMPTZ;

ALTER TABLE inri_invoice ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE inri_contact_person ALTER COLUMN email DROP NOT NULL;
ALTER TABLE inri_contact_person ALTER COLUMN phone DROP NOT NULL;
ALTER TABLE inri_project_documents ADD COLUMN deleted_at TIMESTAMPTZ;

ALTER TABLE inri_project_documents ADD COLUMN created_by uuid references iam_users(id) ON DELETE SET NULL;
ALTER TABLE inri_projects ADD COLUMN repository_url TEXT;
ALTER TABLE inri_invoice ADD COLUMN internal_id TEXT unique NOT NULL DEFAULT gen_random_uuid();