---
has_children: true  
title: core  
parent: gospl  
permalink: "/cmd/core"
---

## gospl core

Query nbs core services

### Options

```
  -h, --help              help for core
      --out-json string   Write results to JSON file
```

### Options inherited from parent commands

```
      --config string     config file (default is $HOME/.gospl.yaml)
      --licence string    Licence id
      --password string   Authenticate with password
      --url string        Webservice url (default "https://webservices.nbs.rs")
      --username string   Authenticate with username
      --verbose           Dump request and response
```

### SEE ALSO

* [gospl](../gospl.md)     - NBS web service command line client
* [gospl core bank](bank/index.md)     - Get bank informations
* [gospl core currency](gospl_core_currency.md)     - Get existing currencies
* [gospl core status](gospl_core_status.md)	 - Get company status
* [gospl core type](gospl_core_type.md)	 - Get company type

###### Auto generated by spf13/cobra on 6-Dec-2021

