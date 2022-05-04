# compserv-poc

This repository is dedicated to experimenting with different approaches for
implementing a RESTful OSCAL API. There is a proof-of-concept [API
reference](https://github.com/EasyDynamics/oscal-rest) available upstream. We
will use this as the basis for generating the API implementation. Even though
it will change in the future, it should be sufficient to exercise various
approaches for generating API server code.

Each approach will be isolated to a branch.

# OpenAPI Generator

This approach uses [OpenAPI Generator](https://openapi-generator.tech/) to
generate a server implementation using the
[go-gin-server](https://openapi-generator.tech/docs/generators/go-gin-server)
generator.

Run the `generate.sh` script to generate the OSCAL API in a container, writing
the generated files to a volume shared with the host:

```console
$ ./generate.sh
```

Because of the difference between users, you'll need to `chown` the generated
files after the script completes:

```console
$ chown -R $USER:$USER .
```

The `generate.sh` script accepts a file path to an OpenAPI specification, but
by default it will use the [proof-of-concept specification
upstream](https://raw.githubusercontent.com/EasyDynamics/oscal-rest/develop/openapi.yaml).

Issues:

- The golang version in the generated `Dockerfile` is hardcoded to 1.10
