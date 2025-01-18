-- +goose Up
-- +goose StatementBegin
-- business_names_202501
CREATE TABLE IF NOT EXISTS business_names (
    abn                    TEXT UNIQUE NOT NULL,
    name                   TEXT NOT NULL,
    status                 TEXT,
    registered_at          TEXT,
    cancel_at              TEXT,
    renew_at               TEXT,
    state_number           TEXT,
    state_of_registration  TEXT
);

INSERT OR REPLACE INTO
    business_names (abn, name, status, registered_at, cancel_at, renew_at, state_number, state_of_registration)
SELECT
    trim(BN_ABN) AS abn,
    trim(BN_NAME) AS name,
    NULLIF(trim(BN_STATUS), '') AS status,
    NULLIF(trim(BN_REG_DT), '') AS registered_at,
    NULLIF(trim(BN_CANCEL_DT), '') AS cancel_at,
    NULLIF(trim(BN_RENEW_DT), '') AS renew_at,
    NULLIF(trim(BN_STATE_NUM), '') AS state_number,
    NULLIF(trim(BN_STATE_OF_REG), '') AS state_of_registration
FROM
    business_names_csv
WHERE
    abn != '' AND
    name != '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE business_names;
-- +goose StatementEnd
