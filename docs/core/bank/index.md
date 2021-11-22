---
title: bank
parent: core
grand_parent: gospl
has_children: true
---
## gospl core bank

Get bank informations

```
gospl core bank [flags]
```

### Options

```
  -c, --bank-code int     Bank code
  -h, --help              help for bank
  -i, --id string         Bank ID
  -n, --national-id int   Bank national ID number
```

### Options inherited from parent commands

```
      --config string     config file (default is $HOME/.gospl.yaml)
      --licence string    Licence id
      --out-json string   Write results to JSON file
      --password string   Authenticate with password
      --url string        Webservice url (default "https://webservices.nbs.rs")
      --username string   Authenticate with username
      --verbose           Dump request and response
```

### SEE ALSO

* [gospl core](../index.md)	 - Query nbs core services
* [gospl core bank status](gospl_core_bank_status.md)	 - Get existing bank statuses
* [gospl core bank type](gospl_core_bank_type.md)	 - Get existing bank types

###### Auto generated by spf13/cobra on 22-Nov-2021