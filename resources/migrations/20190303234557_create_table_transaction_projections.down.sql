DROP FUNCTION public.event_stream_notify ();
DROP TRIGGER events_transaction_stream_notify ON events_transaction_stream;
DROP TABLE transaction_projections;