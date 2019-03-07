CREATE FUNCTION public.event_stream_notify ()
    RETURNS TRIGGER
LANGUAGE plpgsql AS $$
DECLARE
    channel text := TG_ARGV[0];
BEGIN
    PERFORM (
        WITH payload AS
        (
            SELECT NEW.no, NEW.event_name, NEW.metadata -> '_aggregate_id' AS aggregate_id
        )
        SELECT pg_notify(channel, row_to_json(payload)::text) FROM payload
        );
    RETURN NULL;
END;
$$;

DO LANGUAGE plpgsql $$
BEGIN
    IF NOT EXISTS (
       SELECT * FROM information_schema.triggers
       WHERE
           event_object_schema = 'public' AND
           event_object_table = 'events_transaction_stream' AND
           trigger_schema = 'public' AND
           trigger_name = 'events_transaction_stream_notify'
    )
    THEN
       CREATE TRIGGER events_transaction_stream_notify
           AFTER INSERT
           ON events_transaction_stream
           FOR EACH ROW
       EXECUTE PROCEDURE event_stream_notify('transaction_stream');
    END IF;
END;
$$;

CREATE TABLE transaction_projections (
    no SERIAL,
    name VARCHAR(150) UNIQUE NOT NULL,
    position BIGINT NOT NULL DEFAULT 0,
    state JSONB NOT NULL DEFAULT ('{}'),
    locked BOOLEAN NOT NULL DEFAULT (FALSE),

    PRIMARY KEY (no)
);