-- +goose Up
-- Migration: Create the business_search FTS5 virtual table for fast text search on business names.
-- Note: FTS5 does not support IF NOT EXISTS.

-- +goose StatementBegin
CREATE VIRTUAL TABLE business_search USING fts5(abn, name, state);

-- Populate the FTS table from the main business_names table.
INSERT INTO
    business_search
    SELECT 
        abn,
        name,
        state_of_registration
    FROM
        business_names;
-- +goose StatementEnd

-- +goose Down
-- Migration: Drop the business_search FTS5 table.
-- +goose StatementBegin
DROP TABLE business_search;
-- +goose StatementEnd
