BEGIN;
CREATE TABLE IF NOT EXISTS sen.sensor_reading (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sensor_id UUID NOT NULL REFERENCES sen.sensor(id),
    timestamp TIMESTAMPTZ NOT NULL,
    value REAL NOT NULL
);
COMMIT;
