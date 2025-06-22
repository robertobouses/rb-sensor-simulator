UPDATE sen.alert_historial
SET resolved_reading_id = $2
WHERE sensor_id = $1
  AND resolved_reading_id IS NULL;
