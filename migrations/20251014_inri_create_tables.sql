-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS inri_clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    client_type TEXT,
    contract_type TEXT,
    name TEXT,
    tax_number TEXT,
    address TEXT,
    tags TEXT[],
    hour_rate DOUBLE PRECISION,
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
    );


CREATE TABLE IF NOT EXISTS inri_contact_person (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    client_id UUID NOT NULL,
    name TEXT,
    email TEXT,
    phone TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
    );

CREATE TABLE IF NOT EXISTS inri_invoice (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    note TEXT,
    external_id TEXT,
    user_id UUID NOT NULL,
    client_id UUID NOT NULL,
    total DOUBLE PRECISION,
    billable_hours DOUBLE PRECISION,
    issued_at TIMESTAMPTZ,
    due_at TIMESTAMPTZ,
    delivered_at TIMESTAMPTZ,
    paid_at TIMESTAMPTZ,
    is_canceled BOOLEAN DEFAULT FALSE,
    external_link TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS inri_attachments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    model TEXT,
    model_id TEXT,
    original_name TEXT,
    mime_type TEXT,
    size INT,
    server_name TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
    );

-- Projects
CREATE TABLE IF NOT EXISTS inri_projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name TEXT,
    description TEXT,
    status TEXT,
    tags TEXT[],
    slug TEXT,
    is_starred BOOLEAN DEFAULT false NOT NULL,
    last_activity_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- Project Documents
CREATE TABLE IF NOT EXISTS inri_project_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    project_id UUID NOT NULL,
    name TEXT,
    content TEXT,
    tags TEXT[],
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inri_project_documents;
DROP TABLE IF EXISTS inri_projects;
DROP TABLE IF EXISTS inri_attachments;
DROP TABLE IF EXISTS inri_invoice;
DROP TABLE IF EXISTS inri_contact_person;
DROP TABLE IF EXISTS inri_clients;
-- +goose StatementEnd
