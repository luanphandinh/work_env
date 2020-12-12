package main

const (
	CONFIG_DIR  = "config"
	CONFIG_FILE = "config.json"

	PROFILE_DIR = "profile"
)

const (
	PROFILE_USAGE = `
CLI config helps config your current profile configuration.

Availables command:
	set <VAR_NAME> <VAR_VALUE> 	Set <VAR_NAME> with <VAR_VALUE> to current profile env.
	get <VAR_NAME> 			Get value of <VAR_NAME> from current profile env.
	fix 				Fix up current profile environment, all duplicate keys will be fixed as latest values.
	clean 				Clean up current profile, wipe all environment vars.

`

	CONFIG_USAGE = `
CLI config helps config your CLI configuration

Availables command:
	print 					Print out current CLI Config.

Availables options:
	--current-profile <profile_name>	Set working <profile_name> to CLI.
`
)
