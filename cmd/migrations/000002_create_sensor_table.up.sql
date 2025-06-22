BEGIN;

CREATE TABLE IF NOT EXISTS sen.sensor (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    sampling_interval INTEGER NOT NULL,
    alert_threshold_min REAL NOT NULL,
    alert_threshold_max REAL NOT NULL,
    unit TEXT NOT NULL,
    status TEXT NOT NULL
);

COMMIT;
