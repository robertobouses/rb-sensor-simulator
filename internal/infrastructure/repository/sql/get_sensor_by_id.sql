SELECT 
    id,
    name,
    type,
    sampling_interval,
    alert_threshold,
    unit,
    enabled
FROM sen.sensor WHERE id = $1
