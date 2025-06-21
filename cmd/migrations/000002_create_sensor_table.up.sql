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
INSERT INTO sen.sensor (id, name, type, sampling_interval, alert_threshold_min, alert_threshold_max, unit)
VALUES
  (gen_random_uuid(), 'TempSensor-A', 'temperature', 5, -10.0, 50.0, '°C'),
  (gen_random_uuid(), 'HumiditySensor-B', 'humidity', 10, 20.0, 90.0, '%'),
  (gen_random_uuid(), 'PressureSensor-C', 'pressure', 15, 950.0, 1050.0, 'hPa');

COMMIT;
