package cmd

import (
	"errors"

	"github.com/urfave/cli"
	rpc_server "gogs.io/gogs/internal/rpc/server"
	rpc_shared "gogs.io/gogs/internal/rpc/shared"
	log "unknwon.dev/clog/v2"
)

var WebauthnRPC = cli.Command{
	Name:        "webauthn_rpc",
	Usage:       "INTERNAL: Start up the rpc server dealing with webauthn database logic",
	Description: `For internal use only`,
	Action:      runWebauthnRPC,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runWebauthnRPC(c *cli.Context) error {
	// Make sure all of the logs have been processed at the end
	defer log.Stop()

	err := rpc_server.InitServer(c.String("config"))
	if err != nil {
		log.Fatal("Error initializing Webauthn RPC service %v", err)
		return err
	}

	wg, err := rpc_server.StartInternalRPC(new(rpc_server.Webauthn), rpc_shared.WEBAUTHN_RPC_PORT)
	if err != nil {
		log.Fatal("Error starting Webauthn RPC service %v", err)
		return err
	} else {
		log.Info("Webauthn RPC listening on port %d", rpc_shared.WEBAUTHN_RPC_PORT)
	}

	if wg == nil {
		return errors.New("WaitGroup not setup by Webauthn RPC server")
	}

	// Wait for the goroutine server to terminate
	wg.Wait()
	return nil
}
