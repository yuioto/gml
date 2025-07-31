package cmd

import "github.com/urfave/cli/v3"

var Commands = []*cli.Command{
	NewCmd,    // Initialize project
	AddCmd,    // Add mods, resources, etc.
	RemoveCmd, // Remove mods, resources
	UpdateCmd, // Update mods, versions, etc.
	CheckCmd,  // Check dependency integrity, versions, etc.
	BuildCmd,  // Build a complete playable instance (download + assemble)
	RunCmd,    // Launch the game instance
	ExportCmd, // Export integration package
	ImportCmd, // Import integration package
	SearchCmd, // Search for available mods, resources
}
