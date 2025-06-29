BEGIN;
CREATE TABLE IF NOT EXISTS sen.sensor_reading (
    id UUID PRIMARY KEY,
    sensor_id UUID NOT NULL REFERENCES sen.sensor(id),
    timestamp TIMESTAMPTZ NOT NULL,
    value REAL NOT NULL,
    error TEXT
);
COMMIT;
