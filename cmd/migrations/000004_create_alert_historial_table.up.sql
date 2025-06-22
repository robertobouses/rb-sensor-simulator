BEGIN;

CREATE TABLE IF NOT EXISTS sen.alert_historial (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sensor_id UUID NOT NULL REFERENCES sen.sensor(id),
    alert_type TEXT NOT NULL, 
    warning_reading_id UUID NOT NULL REFERENCES sen.sensor_reading(id),
    resolved_reading_id UUID REFERENCES sen.sensor_reading(id),
    created_at TIMESTAMP DEFAULT now()
);

COMMIT;
