BEGIN;
CREATE TABLE IF NOT EXISTS sen.sensor (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    sampling_interval INTEGER NOT NULL,
    alert_threshold REAL NOT NULL,
    unit TEXT NOT NULL,
    enabled BOOLEAN NOT NULL
);
COMMIT;
