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
    name TEXT,
    description TEXT,
    status TEXT,
    tags TEXT[],
    slug TEXT,
    story_sequence integer DEFAULT 0,
    last_activity_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- Project Documents
CREATE TABLE IF NOT EXISTS inri_project_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL,
    name TEXT,
    content TEXT,
    tags TEXT[],
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);


-- N:M Project-User relation
CREATE TYPE inri_project_role AS ENUM ('owner', 'admin', 'developer', 'guest');

CREATE TABLE inri_project_users (
    project_id UUID NOT NULL REFERENCES inri_projects(id) ON DELETE CASCADE,
    user_id    UUID NOT NULL REFERENCES iam_users(id) ON DELETE CASCADE,
    role       inri_project_role NOT NULL DEFAULT 'guest',
    is_starred BOOLEAN DEFAULT false NOT NULL,
    joined_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (project_id, user_id)
);

-- Kanban Board
CREATE TABLE inri_kanban_boards (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL UNIQUE REFERENCES inri_projects(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Kanban Board Columns
CREATE TYPE inri_column_type AS ENUM ('normal', 'done', 'blocked');

CREATE TABLE inri_kanban_columns (
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    board_id  UUID NOT NULL REFERENCES inri_kanban_boards(id) ON DELETE CASCADE,
    key       TEXT NOT NULL,
    name      TEXT NOT NULL,
    type      inri_column_type NOT NULL DEFAULT 'normal',
    position  INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (board_id, key),
    UNIQUE (board_id, position)
);

-- Kanban Stories
CREATE TYPE inri_story_kind AS ENUM (
  'bug', 'feature', 'doc', 'task', 'design', 'plan'
);

CREATE TABLE inri_stories (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id  UUID NOT NULL REFERENCES inri_projects(id) ON DELETE CASCADE,
    board_id    UUID REFERENCES inri_kanban_boards(id) ON DELETE SET NULL,
    column_id   UUID REFERENCES inri_kanban_columns(id) ON DELETE SET NULL,
    slug        TEXT NOT NULL,
    title       TEXT NOT NULL,
    description TEXT,
    kind        inri_story_kind NOT NULL,
    assignee_id UUID REFERENCES iam_users(id),
    priority    INTEGER NOT NULL DEFAULT 0,
    size        INTEGER,
    estimate    INTERVAL,
    start_date  TIMESTAMPTZ,
    end_date    TIMESTAMPTZ,
    position    INTEGER NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Labels
CREATE TABLE inri_labels (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id  UUID NOT NULL,
    name        TEXT NOT NULL,
    description TEXT,
    color       TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE inri_story_labels (
    story_id UUID NOT NULL REFERENCES inri_stories(id) ON DELETE CASCADE,
    label_id UUID NOT NULL REFERENCES inri_labels(id) ON DELETE CASCADE,
    PRIMARY KEY (story_id, label_id)
);

-- Kanban Stories
CREATE TYPE inri_activity_type AS ENUM (
  'comment', 'change_column', 'add_label', 'remove_label',
  'estimate_change', 'change_assignee'
);

CREATE TABLE inri_story_activities (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    story_id   UUID NOT NULL REFERENCES inri_stories(id) ON DELETE CASCADE,
    actor_id   UUID NOT NULL REFERENCES iam_users(id),
    type       inri_activity_type NOT NULL,
    payload    JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Kanban Stories
CREATE TABLE inri_story_time_logs (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    story_id    UUID NOT NULL REFERENCES inri_stories(id) ON DELETE CASCADE,
    user_id     UUID NOT NULL REFERENCES iam_users(id) ON DELETE CASCADE,
    duration    INTERVAL NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT positive_duration CHECK (duration > interval '0')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inri_story_time_logs;
DROP TABLE IF EXISTS inri_story_activities;
DROP TABLE IF EXISTS inri_story_labels;
DROP TABLE IF EXISTS inri_labels;
DROP TABLE IF EXISTS inri_stories;
DROP TABLE IF EXISTS inri_kanban_columns;
DROP TABLE IF EXISTS inri_kanban_boards;
DROP TABLE IF EXISTS inri_project_users;
DROP TABLE IF EXISTS inri_project_documents;
DROP TABLE IF EXISTS inri_projects;
DROP TABLE IF EXISTS inri_attachments;
DROP TABLE IF EXISTS inri_invoice;
DROP TABLE IF EXISTS inri_contact_person;
DROP TABLE IF EXISTS inri_clients;
DROP TYPE inri_project_role;
DROP TYPE inri_column_type;
DROP TYPE inri_story_kind;
DROP TYPE inri_activity_type;
-- +goose StatementEnd
