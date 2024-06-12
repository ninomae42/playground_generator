# playground_generator

## Description

Playground generator is a tool to generate a database environment config files for testing.

This tool generates the following files:
  - `compose.yaml`: Docker compose file to create database/migration containers
  - `taskfile.yaml`: task-runner file to run bunch of commands
  - `base_codes`: Base codes for database migration scripts and containers

## Requirements
- Go 1.18 or later
- Docker, Docker Compose (Optional)
- Task (Optional)

## Usage

```bash
pggen -n [project_name] -o [output_dir]
```

- project_name
  - The name of the project: default is `playground`
  - This will be used as docker compose containers' name prefix and docker network name
- output_dir
  - The output directory: detault is `./playground`
