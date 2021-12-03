# life-utils
Utility tools you'll use for life

## Autocomplete

Bash:

```bash
  $ source <(life-utils completion bash)

  # To load completions for each session, execute once:

  # Linux:
  $ life-utils completion bash > /etc/bash_completion.d/life-utils

  # macOS:
  $ life-utils completion bash > /usr/local/etc/bash_completion.d/life-utils
```

Zsh:

```zsh
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ life-utils completion zsh > "${fpath[1]}/_life-utils"

  # You will need to start a new shell for this setup to take effect.
```

fish:

```sh
  $ life-utils completion fish | source

  # To load completions for each session, execute once:
  $ life-utils completion fish > ~/.config/fish/completions/life-utils.fish
```

PowerShell:

```powershell
  PS> life-utils completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> life-utils completion powershell > life-utils.ps1
  # and source this file from your PowerShell profile.
```

## Development

Adding commands using cobra:

```bash
  $ ~/go/bin/cobra add <name>
```
