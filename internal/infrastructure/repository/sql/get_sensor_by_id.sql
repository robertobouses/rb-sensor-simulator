SELECT 
    id,
    name,
    type,
    sampling_interval,
    alert_threshold_min,
    alert_threshold_max,
    unit
FROM sen.sensor WHERE id = $1
