package commands

import (
	"os"

	"github.com/checkmarxDev/ast-cli/internal/wrappers"
	"github.com/spf13/cobra"
)

func NewUtilsCommand(healthCheckWrapper wrappers.HealthCheckWrapper,
	ssiWrapper wrappers.SastMetadataWrapper,
	rmWrapper wrappers.SastRmWrapper,
	logsWrapper wrappers.LogsWrapper,
	queriesWrapper wrappers.QueriesWrapper,
	uploadsWrapper wrappers.UploadsWrapper,
) *cobra.Command {
	scanCmd := &cobra.Command{
		Use:   "utils",
		Short: "AST Utility functions",
	}
	healthCheckCmd := NewHealthCheckCommand(healthCheckWrapper)
	ssiCmd := NewSastMetadataCommand(ssiWrapper)
	rmCmd := NewSastResourcesCommand(rmWrapper)
	queriesCmd := NewQueryCommand(queriesWrapper, uploadsWrapper)
	logsCmd := NewLogsCommand(logsWrapper)
	//
	/// Complete command
	//
	var completionCmd = &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:
	
	Bash:
	
		$ source <(cx completion bash)
	
		# To load completions for each session, execute once:
		# Linux:
		$ cx completion bash > /etc/bash_completion.d/cx
		# macOS:
		$ cx completion bash > /usr/local/etc/bash_completion.d/cx
	
	Zsh:
	
		# If shell completion is not already enabled in your environment,
		# you will need to enable it.  You can execute the following once:
	
		$ echo "autoload -U compinit; compinit" >> ~/.zshrc
	
		# To load completions for each session, execute once:
		$ cx completion zsh > "${fpath[1]}/_cx"
	
		# You will need to start a new shell for this setup to take effect.
	
	fish:
	
		$ cx completion fish | source
	
		# To load completions for each session, execute once:
		$ cx completion fish > ~/.config/fish/completions/cx.fish
	
	PowerShell:
	
		PS> cx completion powershell | Out-String | Invoke-Expression
	
		# To load completions for every new session, run:
		PS> cx completion powershell > cx.ps1
		# and source this file from your PowerShell profile.
	`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
		},
	}
	scanCmd.AddCommand(healthCheckCmd, ssiCmd, rmCmd, queriesCmd, logsCmd, completionCmd)
	return scanCmd
}
