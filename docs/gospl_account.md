## gospl account

Get account info

### Synopsis

Search for company account info using one or more of the following fields:
company name, national identification number, tax identification number, bank code, account number,
control number or city.


```
gospl account [flags]
```

### Examples

```

gospl account --name somename

```

### Options

```
      --account-number int   Search for account number
      --bank-code int        Search for bank code
      --city string          Search for city
      --control-number int   Search for control number
  -h, --help                 help for account
  -n, --name string          Search for company name
      --national-id int      Search for company national id
      --out-json string      Write results to JSON file
  -t, --tax-id string        Search for company tax id
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

* [gospl](gospl.md)	 - NBS web service command line client
* [gospl account status](gospl_account_status.md)	 - Get existing account statuses
* [gospl account type](gospl_account_type.md)	 - Get existing account types

###### Auto generated by spf13/cobra on 22-Nov-2021