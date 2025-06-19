UPDATE sen.sensor
SET last_reading_timestamp = $2,
    last_reading_value = $3,
    last_reading_error = $4
WHERE id = $1;
