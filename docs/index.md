---
layout: default title: Home permalink: / nav_order: 1
---

# GoSpl

GoSpl is command line client / proxy service for NBS webservices

## Installation

Download archive from [releases page](https://github.com/kockicica/gospl/releases) and unpack it somewhere along the
path.

## Usage

Get basic help using

```shell
gospl help
```

For detailed informations see [CLI usage help](gospl.md)

## Configuration

To use NBS web services you should have authorization data: username, password and licence id. Those parameters should
be specified as gospl flags:

|Flag name | Description | Default value |
|--- | --- | --- |
|username | your username | none |
|password | your password | none |
|licence | your licence id | none |
|url | NBS webservices base url | https://webservices.nbs.rs |

That said, if you want to query for exchange curencies, for example, command line should be something like:

```shell
gospl --username <username> --password <password> --licence <licence-id> core currency
```

If everything goes well, you should get response like this one:

```text
NAME                            CURRENCY CODE   CCN     CCA     COUNTRY CODE NUM        COUNTRY CODE ALPHA      COUNTRY NAME
Lek                             8               008     ALL     008                     ALL                     Albania
Algerian Dinar                  12              012     DZD     012                     DZD                     Algeria
Andorran Peseta                 20              020     ADP     020                     ADP                     Andorra
Argentine Peso                  32              032     ARS     032                     ARS                     Argentina
...
```

To avoid setting those flags at every run you can put them in configuration file and tell gospl where that location file
is. Default name for configuration file is ```.gospl.yaml``` and it should be placed in your ```$HOME``` folder.

Configuration file is simple yaml and its content should be:

```yaml
username: <put-your-username-here>
password: <put-your-password-here>
licence: <put-your-licence-id-here>
```

Command line may be a bit simpler:

```shell
gospl core currency
```

and the response should be the same except first line where location of currently used configuration file should be
displayed:

```html
Using config file: /home/gospler/.gospl.yaml
NAME                            CURRENCY CODE   CCN     CCA     COUNTRY CODE NUM        COUNTRY CODE ALPHA      COUNTRY NAME
Lek                             8               008     ALL     008                     ALL                     Albania
Algerian Dinar                  12              012     DZD     012                     DZD                     Algeria
Andorran Peseta                 20              020     ADP     020                     ADP                     Andorra
...
```

If you dont want to use default configuration file location you can specify location of the config file you would like to use as ```config``` flag:

```shell
gospl --config /path/to/config.yaml core currency
```


