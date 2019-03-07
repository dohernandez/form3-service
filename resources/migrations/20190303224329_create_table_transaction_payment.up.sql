CREATE TABLE transaction_payment (
    id UUID PRIMARY KEY NOT NULL,
    version SMALLINT NOT NULL,
    organisation_id UUID NOT NULL,
    attributes JSONB NOT NULL DEFAULT ('{}')
);