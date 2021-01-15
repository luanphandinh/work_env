package main

const (
	configDir  = "config"
	configFile = "config.json"
	porfileDir = "profile"
)

const (
	envUsage = `
CLI config helps config your current profile configuration.

Availables command:
	set <VAR_NAME> <VAR_VALUE> 	Set <VAR_NAME> with <VAR_VALUE> to current profile env.
	print 				Get value from cli env.
	fix 				Fix up current profile environment, all duplicate keys will be fixed as latest values.
	clean 				Clean up current profile, wipe all environment vars.

`

	configUsage = `
CLI config helps config your CLI configuration

Availables command:
	print 					Print out current CLI Config.

Availables options:
	--current-profile <profile_name>	Set working <profile_name> to CLI.
`
)
