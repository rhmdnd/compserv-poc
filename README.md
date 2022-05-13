# compserv-poc

This repository is dedicated to experimenting with different approaches for
implementing a RESTful OSCAL API. There is a proof-of-concept [API
reference](https://github.com/EasyDynamics/oscal-rest) available upstream. We
will use this as the basis for generating the API implementation. Even though
it will change in the future, it should be sufficient to exercise various
approaches for generating API server code.

Each approach will be isolated to a branch.

## Deploying

You can deploy the infrastructure using `terraform`.

``console
$ cd terraform/
$ terraform apply
```

The database password is kept in AWS Secret Manager and you must log into to
fetch it.
