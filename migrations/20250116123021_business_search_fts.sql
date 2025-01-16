-- +goose Up
-- +goose StatementBegin
CREATE VIRTUAL TABLE business_search USING fts5(abn, name, state);

INSERT INTO
    business_search
    SELECT 
        abn,
        name,
        state_of_reg
    FROM
        business_names;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE business_search;
-- +goose StatementEnd
