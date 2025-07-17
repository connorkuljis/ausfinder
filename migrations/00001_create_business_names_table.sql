-- +goose Up
-- Migration: Create the business_names table and import data from a staging CSV table.
-- abn is the primary key for uniqueness and performance.

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS business_names (
    abn                    TEXT PRIMARY KEY NOT NULL,
    name                   TEXT NOT NULL,
    status                 TEXT,
    registered_at          TEXT,
    cancel_at              TEXT,
    state_number           TEXT,
    state_of_registration  TEXT
);

-- Import data from the staging CSV table (business_names_csv).
INSERT OR REPLACE INTO
    business_names (abn, name, status, registered_at, cancel_at, state_number, state_of_registration)
SELECT
    trim(BN_ABN) AS abn,
    trim(BN_NAME) AS name,
    NULLIF(trim(BN_STATUS), '') AS status,
    NULLIF(trim(BN_REG_DT), '') AS registered_at,
    NULLIF(trim(BN_CANCEL_DT), '') AS cancel_at,
    NULLIF(trim(BN_STATE_NUM), '') AS state_number,
    NULLIF(trim(BN_STATE_OF_REG), '') AS state_of_registration
FROM
    business_names_csv
WHERE
    abn != '' AND
    name != '';

-- Drop the staging CSV table after import.
DROP TABLE business_names_csv;
-- +goose StatementEnd

-- +goose Down
-- Migration: Drop the business_names table.
-- +goose StatementBegin
DROP TABLE business_names;
-- +goose StatementEnd
