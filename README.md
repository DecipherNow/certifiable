# Certifiable

A completely untrustworthy certificate authority and tooling for making development with certificates easy.

## Overview

This repository provides the scripts required to build a self contained certificate authority that provides an HTTP API for creating and retrieving server and user certificates.

## Building

To build and push the certificate authority image run the following command from the root directory of this repository:

    ./publish

## Running

To run the certificate authority simply execute the following:

    docker run -p 4567:4567 deciphernow/certifiable:latest
    
This will start the certificate authority container and expose the HTTP API on port [http://localhost:4567](http://localhost:4567).

## Usage

The following table describes the endpoints exposed by the certificate authority API.

| Method | Endpoint                         | Description                                 |
|--------|----------------------------------|---------------------------------------------|
| GET    | /root                            | List the root certificate files.            |
| GET    | /root/:file                      | Download a root certificate file.           |
| GET    | /intermediate                    | List the intermediate certificate files.    |
| GET    | /intermediate/:file              | Download an intermediate certificate file.  |
| GET    | /servers                         | List the servers with certificates.         |
| GET    | /servers/:server                 | List the certificate files for a server.    |
| POST   | /servers/:server?cn=:common_name | Creates certificates for server.            |
| GET    | /servers/:server/:file           | Download a certificate file for a server.   |
| GET    | /users                           | List the users with certificates.           |
| GET    | /users/:user                     | List the certificate files for a user.      |
| POST   | /users/:user?cn=:common_name     | Creates certificates for user.              |
| GET    | /users/:user/:file               | Download a certificate file for a user.     |

## Discussion and Examples

https://forums.deciphernow.com/t/update-expired-localhost-pki-certificate-steps/625
