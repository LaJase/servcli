## servcli completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(servcli completion zsh)

To load completions for every new session, execute once:

#### Linux:

	servcli completion zsh > "${fpath[1]}/_servcli"

#### macOS:

	servcli completion zsh > $(brew --prefix)/share/zsh/site-functions/_servcli

You will need to start a new shell for this setup to take effect.


```
servcli completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -c, --config string   config file (default is $PWD/config/servcli-config.yaml)
```

### SEE ALSO

* [servcli completion](servcli_completion.md)	 - Generate the autocompletion script for the specified shell

