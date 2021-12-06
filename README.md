# Canivete
Multi-tools you'll use for life.


## Origin

"Canivete" is the Portuguese word for [pocket knife](https://pt.wikipedia.org/wiki/Canivete).

Ths inspiration for the name came from multi-tools knifes like the Swiss Army knifes.


## Available commands

| Group | Name | Description  |
|---|---|---|
| datetime | fromunix | Converts a Unix timestamp to human friendly format |
| finance | compoundinterests | Calculates compound interests |
| internet | medium2md | Converts a [Medium](https://medium.com) post to markdown |
| programming | uuid | Generates UUIDs |

## Getting help

General help:

```zsh
$ canivete --help
```

Help for a group of commands:

```zsh
$ canivete <group name> --help
```

E.g.

```zsh
$ canivete finance --help
```


## How to install the autocomplete

Bash:

```bash
$ source <(canivete completion bash)

# To load completions for each session, execute once:

# Linux:
$ canivete completion bash > /etc/bash_completion.d/canivete

# macOS:
$ canivete completion bash > /usr/local/etc/bash_completion.d/canivete
```

Zsh:

```zsh
# If shell completion is not already enabled in your environment,
# you will need to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ canivete completion zsh > "${fpath[1]}/_canivete"

# You will need to start a new shell for this setup to take effect.
```

fish:

```sh
$ canivete completion fish | source

# To load completions for each session, execute once:
$ canivete completion fish > ~/.config/fish/completions/canivete.fish
```

PowerShell:

```powershell
PS> canivete completion powershell | Out-String | Invoke-Expression

# To load completions for every new session, run:
PS> canivete completion powershell > canivete.ps1
# and source this file from your PowerShell profile.
```

## Development

### Semantic Commit Messages

See how a minor change to your commit message style can make you a better programmer.

Format: `<type>(<scope>): <subject>`

`<scope>` is optional

#### Example

```
feat: add hat wobble
^--^  ^------------^
|     |
|     +-> Summary in present tense.
|
+-------> Type: chore, docs, feat, fix, refactor, style, or test.
```

More Examples:

- `feat`: (new feature for the user, not a new feature for build script)
- `fix`: (bug fix for the user, not a fix to a build script)
- `docs`: (changes to the documentation)
- `style`: (formatting, missing semi colons, etc; no production code change)
- `refactor`: (refactoring production code, eg. renaming a variable)
- `test`: (adding missing tests, refactoring tests; no production code change)
- `chore`: (updating grunt tasks etc; no production code change)

References:

- https://www.conventionalcommits.org/
- https://seesparkbox.com/foundry/semantic_commit_messages
- http://karma-runner.github.io/1.0/dev/git-commit-msg.html

### Adding commands using cobra

```bash
  $ ~/go/bin/cobra add <name>
```
