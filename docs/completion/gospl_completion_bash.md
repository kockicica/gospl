---
title: bash
parent: completion
grand_parent: gospl
---

## gospl completion bash

generate the autocompletion script for bash

### Synopsis


Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:
$ source <(gospl completion bash)

To load completions for every new session, execute once:
Linux:
  $ gospl completion bash > /etc/bash_completion.d/gospl
MacOS:
  $ gospl completion bash > /usr/local/etc/bash_completion.d/gospl

You will need to start a new shell for this setup to take effect.
  

```
gospl completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
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

* [gospl completion](index.md)	 - generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 22-Nov-2021