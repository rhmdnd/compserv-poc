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

```console
$ cd terraform/
$ terraform apply
```

The database password is kept in AWS Secret Manager and you must log into to
fetch it.

# OpenAPI Generator

This approach uses [OpenAPI Generator](https://openapi-generator.tech/) to
generate a server implementation using the
[go-server](https://openapi-generator.tech/docs/generators/go-server/)
generator.

## Prerequisites

The scripts to generate the server implementation use the OpenAPI generator
CLI, which is a jar file. The Vagrantfile in this repository creates a VirtuaBox
instance for the dependencies. The generated code is synced to the host using a
sync folder, using the `virtualbox` sync type in Vagrant. The scripts in
`tools/` expect you have
[VirtualBox](https://www.vagrantup.com/docs/providers/virtualbox) and
[Vagrant](https://www.vagrantup.com/docs/installation) installed.

## Usage

Start the VM:

```console
$ vagrant up; vagrant ssh
```

Update the package index and install necessary packages:

```console
$ bash /vagrant/tools/install.sh
```

Generate the code:

```console
$ bash /vagrant/tools/generate.sh
```
