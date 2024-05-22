Environment variables, their defaults, and their purposes:
| Environment Variable | Default Value   | Purpose                                                 |
|----------------------|-----------------|---------------------------------------------------------|
| PORT                 | 3000            | The port for the service                                |
| DB_USER              | root            | The username for the database                           |
| DB_PASSWORD          | root            | The password for the database                           |
| DB_HOST              | localhost       | The host for the database                               |
| DB_PORT              | 3306            | The port for the database                               |
| DB_NAME              | provider        | The name of the database                                |
| CONSUL_ADDRESS       | localhost:8500  | The address for the consul service registry             |
| CONSUL_SERVICE_NAME  | provider        | The name of the service in the consul service registry  |
| CONSUL_SERVICE_PATH  | /provider       | The path that defines the traefik router for the service|