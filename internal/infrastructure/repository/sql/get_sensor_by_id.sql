SELECT 
    id,
    name,
    type,
    sampling_interval,
    alert_threshold,
    unit,
    enabled,
    last_reading_timestamp,
    last_reading_value,
    last_reading_error
FROM sen.sensor WHERE id = $1
