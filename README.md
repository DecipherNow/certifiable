# Certifiable

A command line utility for generating single use certificate authorities and issuing certificates.

## Overview

There are a number of situations in which having an ephemeral certificate authority is sufficient or optimal. Certifiable fills this role by generating a new certificate authority on every run and issuing the requested certificates from that authority. None of the authority's cryptographic key material is persisted to disk so once the run completes no additional certificates can be issued by that certificate authority.

## Usage

### Installation

### Single
In order to generate a single certificate run the following command.

```
certifiable issue single -c localhost -a localhost -a localhost.localdomain -e 86000s
```

This will create the following files in the current directory.

| File             | Description                                       |
|------------------|---------------------------------------------------|
| root.crt         | The PEM encoded root certificate.                 |
| intermediate.crt | The PEM encoded intermediate certificate.         |
| localhost.crt    | The PEM encoded server certificate for localhost. |
| localhost.key    | The PEM encoded server key for localhost.         |

### Batch
In order to generate a multiple certificates run the following command.

```
certifiable issue batch --config ./config.json
```

Where `config.json` contains the following content.

```
[
	{
		"commonName": "localhost",
		"alternativeNames": [
			"localhost",
			"localhost.localdomain"
		],
		"expires": "24h"
	},
	{
		"commonName": "example.com",
		"alternativeNames": [
			"example.com",
			"www.example.com"
		],
		"expires": "1h"
	}
]
```

This will create the following files in the current directory.

| File             | Description                                         |
|------------------|-----------------------------------------------------|
| root.crt         | The PEM encoded root certificate.                   |
| intermediate.crt | The PEM encoded intermediate certificate.           |
| localhost.crt    | The PEM encoded server certificate for localhost.   |
| localhost.key    | The PEM encoded server key for localhost.           |
| example.com.crt  | The PEM encoded server certificate for example.com. |
| example.com.key  | The PEM encoded server key for example.com.         |

## Development

### Prerequisites

### Building

### Testing

### Deploying
