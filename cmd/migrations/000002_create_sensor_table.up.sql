BEGIN;
CREATE TABLE IF NOT EXISTS sen.sensor (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    sampling_interval INTEGER NOT NULL,
    alert_threshold_min REAL NOT NULL,
    alert_threshold_max REAL NOT NULL,
    unit TEXT NOT NULL
);
COMMIT;
