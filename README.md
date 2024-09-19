# go-proxy-make

This project allows the creation of HTTP proxies based on YAML configuration files. 
It automatically generates proxy functions that can be used to forward HTTP requests to other services.

## How config a new proxy

### Environment configuration

| Environment Variable | Description                                                                                  | Default Value                      |
|----------------------|----------------------------------------------------------------------------------------------|------------------------------------|
| `CONFIG_DIR`         | Directory where the YAML configuration files are located.                                    | `./setup`                          |
| `OUTPUT_DIR`         | Directory where the generated code will be saved.                                            | `./cmd/api/generated`              |
| `SERVER_PORT`        | Port where the HTTP server will be started.                                                  | `:8081`                            |

### Where to Place the YAML Configuration File

The YAML configuration files for the proxy should be placed in the directory specified by the `CONFIG_DIR` 
environment variable. By default, this directory is `/setup`.

### How to Fill the YAML Configuration File

The YAML files must follow the [OpenAPI 3.0.0](https://swagger.io/specification/) structure. 

The `proxy_mapping` field is used to configure the proxy. It must contain the following fields:

| Field          | Description                                                                 |
|----------------|-----------------------------------------------------------------------------|
| `proxy_host`   | The base URL of the host where the proxy will forward the requests.         |
| `proxy_endpoint` | The specific endpoint on the proxy host to which the requests will be forwarded. |
| `proxy_method` | The HTTP method (e.g., GET, POST, PUT) to be used when forwarding the requests. |
| `field_mappings` | A mapping of fields from the incoming request to the fields expected by the proxy endpoint. |

Below is an example of how a YAML file should be configured to set up the proxies:

```yaml
openapi: 3.0.0
info:
  title: API Example
paths:
  /order:
    post:
      summary: Submit an order
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                orderId:
                  type: string
                name:
                  type: string
                count:
                  type: string
      proxy_mapping:
        proxy_host: https://{my-proxy-api-host}
        proxy_endpoint: /{my-proxy-api-path}
        proxy_method: post
        field_mappings:
          orderId: id
          name: product
          count: quantity
    get:
      proxy_mapping:
        proxy_host: https://{my-proxy-api-host}
        proxy_endpoint: /{my-proxy-api-path}
        proxy_method: get
        field_mappings:
          orderId: id
          name: product
          count: quantity
```

### How to Generate the Proxy Code

After creating the YAML configuration file, you can generate the proxy code by following these steps:

1. Ensure that the YAML configuration files are placed in the directory specified by the `CONFIG_DIR` 
2. environment variable. By default, this directory is `/setup`.
3. Run the following command to generate the proxy code:

   ```sh
   go run cmd/cli/main.go
    ```
   This command will read the `YAML` configuration files and generate the corresponding proxy functions.
4. The generated code will be saved in the directory specified by the `OUTPUT_DIR` environment variable. 
By default, this directory is `/cmd/api/generated`.

Make sure to set the `CONFIG_DIR` and `OUTPUT_DIR` environment variables if you are using custom directories.

### How to Use the Generated Proxy

After generating the proxy code, start the server with the following command:  

```sh
go run cmd/api/main.go
```

The server will start on the port specified by the `SERVER_PORT` environment variable (default is `:8081`). 

You can now make HTTP requests to the proxy endpoints. Below are examples of how to execute calls using curl:  

```sh
curl -X GET http://localhost:8081/order
```

> [!NOTE]
> Attention, this case is just an example, if it is necessary to isolate contexts,
> images can be generated using only the content generated from the specified proxy.
