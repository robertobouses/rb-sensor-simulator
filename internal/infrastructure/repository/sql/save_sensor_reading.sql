INSERT INTO sen.sensor_reading (
    id,
    sensor_id,
    timestamp,
    value,
    error
)VALUES (gen_random_uuid(), $1, $2, $3, $4);
