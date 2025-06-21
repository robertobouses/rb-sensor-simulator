INSERT INTO sen.sensor_reading (
    id,
    sensor_id,
    timestamp,
    value
)VALUES (gen_random_uuid(), $1, $2, $3);
