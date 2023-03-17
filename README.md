# openshiftappslogin

Retrieves a bearer token from an OpenShift CI cluster.

## Usage:

```
openshiftappslogin [flags]

Flags:
--config string     config file (default is $XDG_CONFIG_HOME/openshiftappslogin/config.yaml)
-h, --help              help for openshiftappslogin
-p, --password string   User password
-o, --url string        Token request URL
-u, --username string   User name
```

## config.yaml

Use a config file to setup defaults.

| key      | description             |
|----------|-------------------------|
| url      | OAuth token request URL |
| username | Username                |
| password | PIN + token             |
| secret   | OTP token               |
| prefix   | PIN                     |

Specify `secret` and `prefix` to have your PIN+token password generated for you.