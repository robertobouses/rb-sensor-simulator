INSERT INTO sen.sensor_reading (
    id,
    sensor_id,
    timestamp,
    value,
    error
)VALUES ($1, $2, $3, $4, $5);
