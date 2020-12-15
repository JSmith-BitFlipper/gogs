package cmd

import (
	"errors"
	"github.com/urfave/cli"
	"gogs.io/gogs/internal/route"
	rpc_server "gogs.io/gogs/internal/rpc/server"
	log "unknwon.dev/clog/v2"
)

var RepoRPC = cli.Command{
	Name:        "repo_rpc",
	Usage:       "INTERNAL: Start up the rpc server dealing with repository logic",
	Description: `For internal use only`,
	Action:      runRepoRPC,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runRepoRPC(c *cli.Context) error {
	// Make sure all of the logs have been processed at the end
	defer log.Stop()

	// Initilize the database
	route.InitOnlyDB(c.String("config"))

	wg, err := rpc_server.StartInternalRPC()
	if err != nil {
		log.Fatal("Error starting Repo RPC service %v", err)
		return err
	} else {
		log.Info("Repo RPC server started")
	}

	if wg == nil {
		return errors.New("WaitGroup not setup by Repo RPC server")
	}

	// Wait for the goroutine server to terminate
	wg.Wait()
	return nil
}
