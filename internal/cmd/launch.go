package cmd

import (
	"github.com/urfave/cli"
	"os"
)

var Launch = cli.Command{
	Name:  "launch",
	Usage: "Launch all of the services in processess",
	Description: `Gogs launcher is the only thing you need to run,
and it takes care of all the other things for you`,
	Action: runLauncher,
	Flags: []cli.Flag{
		stringFlag("port, p", "3000", "Temporary port number of web server to prevent conflict"),
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runLauncher(c *cli.Context) (err error) {
	var args []string
	var p *os.Process
	var processes []*os.Process

	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

	// Start up the RPC servers
	args = []string{"./gogs", "repo_rpc"}
	p, err = os.StartProcess(args[0], args, &procAttr)

	if err != nil {
		return err
	}

	processes = append(processes, p)

	// Pass the `port` and `config` args through to the web server
	args = []string{"./gogs", "web", "--port", c.String("port"), "--config", c.String("config")}
	p, err = os.StartProcess(args[0], args, &procAttr)

	if err != nil {
		return err
	}

	processes = append(processes, p)

	// TODO: In theory, this is just a starter process that should disapear after
	// setting everything up. I keep this here to capture Ctrl-C kill
	// using for loop
	for _, p = range processes {
		p.Wait()
	}
	return nil
}
