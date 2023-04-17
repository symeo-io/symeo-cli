<h1 align="center">
<a href="https://app.symeo.io/">
  <img width="300" src="https://s3.eu-west-3.amazonaws.com/symeo.io-assets/symeo-logo.png" alt="symeo">
</a>
</h1>
<p align="center">
  <p align="center">Secret management as code. Easy. Centralized. Secured.</p>
</p>


<h4 align="center">
  <a href="https://app.symeo.io/">SaaS</a> |
  <a href="https://symeo.io/">Website</a> |
  <a href="https://docs.symeo.io/">Docs</a>
</h4>

<h4 align="center">
  <a href="https://github.com/medusajs/medusa/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/license-Apache-blue.svg" />
  </a>
 <a href="https://circleci.com/gh/symeo-io/symeo-cli">
    <img src="https://circleci.com/gh/symeo-io/symeo-cli.svg?style=svg"/>
 </a>
</h4>

# Symeo CLI

The Symeo CLI made for interacting with your Symeo secrets and configuration.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Symeo CLI options](#symeo-cli-commands)

## Install

Run

```bash
sudo curl https://raw.githubusercontent.com/symeo-io/symeo-cli/main/scripts/install.sh | bash
```

## Usage

### Define your configuration contract

Create a `symeo.config.yml` file in the root of your project, and define the structure and types of your application configuration.

For example:
```yaml
database:
  host:
    type: string
  port:
    type: integer
  username:
    type: string
  password:
    type: string
    secret: true
    regex: ^[a-zA-Z0-9]+$
```

- You can nest properties to any depth level
- Supported types are `boolean`, `string`, `integer` and `float`
- Properties can be flagged with `optional: true`, or `secret: true`
- For type `string`, you can add a regex expression that the value will have to match

If you prefer, you can also directly list environment variable-like variables:

```yaml
DATABASE_HOST:
  type: string
DATABASE_PORT:
  type: integer
DATABASE_USERNAME:
  type: string
DATABASE_PASSWORD:
  type: string
  secret: true
  regex: ^[a-zA-Z0-9]+$
```

### Create your local configuration file

Create a `symeo.local.yml` file in the root of your project, defining the values matching your configuration contract.

For example:
```yaml
database:
  host: "localhost"
  port: 5432
  username: "postgres"
  password: "XPJc5qAbQcn77GWg"
```

### Wrap your application startup with the symeo command

To inject your configuration as environment variables, wrap your command like the following:

```bash
symeo-cli start -- your_command
```

This will:
- read your contract and your local file
- Check your values are compliant with the contract
- convert your values into environment variables
- inject them in the given command process

For example, for the previous local file, the following values will be injected:

```bash
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=XPJc5qAbQcn77GWg
```

The contract values are converted in *screaming snake case*. For example, `database.migrationFilePath` will become `DATABASE_MIRGATION_FILE_PATH`.

### Start application with configuration from Symeo platform

After creating an environment and its api key in the [Symeo platform](https://app.symeo.io/), use this command to inject environment values:

```bash
symeo-cli start --api-key $YOUR_ENVIRONMENT_API_KEY -- your_command
```

This will:
- read your contract
- Fetch your values from the Symeo platform
- Check your values are compliant with the contract
- convert your values into environment variables
- inject them in the given command process

### Check your configuration is valid

In your CI or CD pipeline, run:

```shell
symeo-cli validate --api-key $YOUR_ENVIRONMENT_API_KEY
```

Which will check if the values filled in the Symeo platform comply with your contract.

## Symeo CLI commands

### symeo-cli start

Run your command with your configuration values as environment variables, either read from a local file or fetched from the Symeo platform.

#### `-c, --contract-file`

The path to your configuration contract file. Default is `symeo.config.yml`.

#### `-f, --values-file`

The path to your local values file. Default is `symeo.local.yml`.

#### `-k, --api-key`

The environment api key to use to fetch values from Symeo platform. If empty, values will be fetched from local value file (`symeo.local.yml` by default). If specified, parameter `-f, --values-file` is ignored.

#### `-a, --api-url`

The api endpoint used to fetch your configuration with the api key. Default is `https://api.symeo.io/api/v1/values`.


### symeo-cli validate

Check that with your configuration values, either read from a local file or fetched from the Symeo platform, match your contract.

#### `-c, --contract-file`

The path to your configuration contract file. Default is `symeo.config.yml`.

#### `-f, --values-file`

The path to your local values file. Default is `symeo.local.yml`.

#### `-k, --api-key`

The environment api key to use to fetch values from Symeo platform. If empty, values will be fetched from local value file (`symeo.local.yml` by default). If specified, parameter `-f, --values-file` is ignored.

#### `-a, --api-url`

The api endpoint used to fetch your configuration with the api key. Default is `https://api.symeo.io/api/v1/values`.
