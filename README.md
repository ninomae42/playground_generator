# playground_generator

## Description

Playground generator is a tool to generate a database environment config files for testing.

This tool generates the following files:
  - `compose.yaml`: Docker compose file to create database/migration containers
  - `taskfile.yaml`: task-runner file to run bunch of commands

This tool also copies the following directories:
  - `base_codes/`
    - The files/directories in this directory will be copied to the output directory.
    - By default, base configuration codes for database migration scripts and containers are stored. (see `base_code` directory)

## Requirements

- Go 1.18 or later
- Docker, Docker Compose (Optional)
- [Task](https://taskfile.dev/) (Optional)

## Installation

```bash
git clone https://github.com/ninomae42/playground_generator.git
cd playground_generator
go install
```

## Usage

```bash
pggen -n [project_name] -o [output_dir]
```

- project_name
  - The name of the project: default is `playground`
  - This will be used as docker compose containers' name prefix and docker network name
- output_dir
  - The output directory: detault is `./playground`
