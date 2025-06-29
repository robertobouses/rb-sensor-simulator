# rb-sensor-simulator

IoT sensor simulator built in Go. Includes configurable sensor behavior, NATS-based messaging, and data persistence.






## Features

- Periodic simulation of sensor readings, each with its own logic and configuration.
- NATS-based API for:
  - Registering and updating sensor configuration (sampling interval, alert thresholds).
  - Querying the current configuration and last reading of each sensor.
  - Publishing sensor readings.
- Error simulation in readings (e.g., invalid values or sensor failure).
- Alert system with history for out-of-threshold readings.
- Data persistence in PostgreSQL.
- Unit tests for core use cases.






## Technologies

- **Golang** 1.24.3
- **NATS** using [nats.go](https://github.com/nats-io/nats.go)
- **PostgreSQL**
- **Docker** (for infrastructure and services)






## Project Structure

 .env
│   .gitignore
│   docker-compose.yml
│   go.mod
│   go.sum
│   rb-Diagrama Explicativo.png
│   README.md
│
├───cmd
│   │   main.go
│   │
│   ├───migrations
│   │       000001_create_sen_schema.down.sql
│   │       000001_create_sen_schema.up.sql
│   │       000002_create_sensor_table.down.sql
│   │       000002_create_sensor_table.up.sql
│   │       000003_create_sensor_reading_table.down.sql
│   │       000003_create_sensor_reading_table.up.sql
│   │       000004_create_alert_historial_table.down.sql
│   │       000004_create_alert_historial_table.up.sql
│   │
│   └───simulator
│           main.go
│           sensors.csv
│
└───internal
    ├───domain
    │   │   alert_historial.go
    │   │   sensor.go
    │   │   sensor_reading.go
    │   │
    │   └───use_cases
    │           app_service.go
    │           create_sensor.go
    │           get_sensor_config_and_last_readings.go
    │           get_sensor_config_and_last_readings_test.go
    │           mock_repo_test.go
    │           save_sensor_reading.go
    │           save_sensor_reading_test.go
    │           save_sensor_test.go
    │           update_sensor_config.go
    │           update_sensor_config_test.go
    │
    ├───infrastructure
    │   ├───nats
    │   │       create_sensor.go
    │   │       get_sensor_config_and_last_readings.go
    │   │       handler.go
    │   │       process_sensor_reading.go
    │   │       server.go
    │   │       update_sensor_config.go
    │   │
    │   └───repository
    │       │   get_sensor_by_id.go
    │       │   get_sensor_config_and_last_readings.go
    │       │   repository.go
    │       │   save_alert.go
    │       │   save_sensor.go
    │       │   save_sensor_reading.go
    │       │   update_alert_resolved.go
    │       │   update_sensor_config.go
    │       │
    │       └───sql
    │               get_sensor_by_id.sql
    │               get_sensor_last_readings_by_id.sql
    │               save_alert.sql
    │               save_sensor.sql
    │               save_sensor_reading.sql
    │               update_alert_resolved.sql
    │               update_sensor_config.sql
    │
    └───pkg
        └───postgres
                client.go






## How to Run

1. **Start services with Docker:**

   docker-compose up -d




2. **Run database migrations:**

docker run -v ${PWD}/cmd/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" up




3. **Down database migrations:**

docker run -v ${PWD}/cmd/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" down --all




4. **Run the application (in two separate terminals):**

**Terminal 1 – Start the NATS handlers and app logic:**

go run cmd/main.go


**Terminal 2 – Start the sensor simulator:**
cd cmd/simulator
go run main.go




5. **Testing**

**To run all unit tests:**

go test ./...




**To check code coverage for use cases:**

go test -cover ./internal/domain/use_cases