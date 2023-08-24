# DPM APIs
This repository contains the codebase of the DPM APIs that are used within
our [mobile application](https://github.com/desmos-labs/dpm).

## Development

In order to run an instance of this APIs, you will need to provide the following environment variables:

| Name                      | Description                                                      | Required | Default   |
|---------------------------|------------------------------------------------------------------|----------|-----------| 
| `SERVER_ADDRESS`          | Address where the server will be listening for connections       | No       | `0.0.0.0` |
| `SERVER_PORT`             | Port where the server will be listening for connections          | No       | `3000`    |
| `CAERUS_GRPC_ADDRESS`     | Address of Caerus instance to use                                | Yes      | -         |
| `CAERUS_GRPC_IS_INSECURE` | Tells whether the connection to Caerus should be insecure or not | No       | `false`   |
| `LOG_LEVEL`               | Log level to use                                                 | No       | `info`    |

## Available endpoints

### Deep Links

#### Create generic address deep link
This endpoint allows to create a deep link that can be used to open the DPM application and allow the user to select
what action to take on the given address.

Endpoint

```
GET /deep-links/{address}?chain_type=<chain_type>
```

Params:

* the `chain_type` param represents the chain for which the link should be generated (either `testnet` or `mainnet`)

Example response body

```json
{
  "deep_link": "https://desmos.app.link/..."
}
```

#### View profile
This endpoint allows to create a deep link that can be used to open the DPM application and allow the user to view
the profile of the given address.

Endpoint

```
GET /deep-links/{address}/view_profile?chain_type=<chain_type>
```

Params:

* the `chain_type` param represents the chain for which the link should be generated (either `testnet` or `mainnet`)

Example response body

```json
{
  "deep_link": "https://desmos.app.link/..."
}
```

#### Send tokens
This endpoint allows to create a deep link that can be used to open the DPM application and allow the user to send
tokens to the given address.

Endpoint

```
GET /deep-links/{address}/send?amount=<amount>&chain_type=<chain_type>
```

Params:

* the `amount` param represents the optional amount of tokens to send. If provided it must be a valid Cosmos coins
  amount encoded in the string format (i.e. `10udaric`)
* the `chain_type` param represents the chain for which the link should be generated (either `testnet` or `mainnet`)