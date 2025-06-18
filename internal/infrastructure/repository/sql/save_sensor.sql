INSERT INTO sen.sensor (
    name,
    type,
    sampling_interval,
    alert_threshold,
    unit,
    enabled,
    last_reading_timestamp,
    last_reading_value,
    last_reading_error
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;
