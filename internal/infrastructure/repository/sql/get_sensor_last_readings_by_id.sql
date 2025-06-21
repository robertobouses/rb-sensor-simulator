SELECT 
    timestamp,
    value
FROM sen.sensor_reading
WHERE sensor_id = $1
ORDER BY timestamp DESC
LIMIT $2
