package server_context

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gogs.io/gogs/internal/context"
	"gogs.io/gogs/internal/db"
)

func Session2User(c *context.Context) {
	uid, ok := c.Session.Get("uid").(int64)
	c.JSONSuccess(map[string]interface{}{
		"ok":  ok,
		"uid": uid,
	})
}

func itemFromIDs(itemType string, args []string) (payload interface{}, err error) {
	ids := make([]interface{}, len(args))

	switch itemType {
	case "ssh_key", "email", "repository":
		// Single `int64` id
		if len(args) != 1 {
			err = fmt.Errorf("Invalid number of args for context")
			return
		}

		id, parse_err := strconv.ParseInt(args[0], 10, 64)
		if parse_err != nil {
			err = fmt.Errorf("Unable to parse context ids: %v", args)
			return
		}
		ids[0] = id
	case "attachment":
		// Single `string` id
		if len(args) != 1 {
			err = fmt.Errorf("Invalid number of args for context")
			return
		}
		ids[0] = args[0]
	case "app_token":
		// Double `int64` ids
		if len(args) != 2 {
			err = fmt.Errorf("Invalid number of args for context")
			return
		}

		id1, parse1_err := strconv.ParseInt(args[0], 10, 64)
		id2, parse2_err := strconv.ParseInt(args[1], 10, 64)

		if parse1_err != nil || parse2_err != nil {
			err = fmt.Errorf("Unable to parse context ids: %v", args)
			return
		}
		ids[0] = id1
		ids[1] = id2
	case "repo_webhook":
		// Two `strings` and a `int64` id
		if len(args) != 3 {
			err = fmt.Errorf("Invalid number of args for context")
			return
		}

		id, parse_err := strconv.ParseInt(args[2], 10, 64)
		if parse_err != nil {
			err = fmt.Errorf("Unable to parse context ids: %v", args)
			return
		}

		ids[0] = args[0]
		ids[1] = args[1]
		ids[2] = id
	default:
		err = fmt.Errorf("Unknown item type: %s", itemType)
		return
	}

	switch itemType {
	case "app_token":
		payload, err = db.AccessTokens.GetByID(ids[0].(int64), ids[1].(int64))
	case "attachment":
		payload, err = db.GetAttachmentByUUID(ids[0].(string))
	case "email":
		payload, err = db.GetEmailByID(ids[0].(int64))
	case "repository":
		payload, err = db.GetRepositoryByID(ids[0].(int64))
	case "ssh_key":
		payload, err = db.GetPublicKeyByID(ids[0].(int64))
	case "repo_webhook":
		var user *db.User
		user, err = db.GetUserByName(ids[0].(string))
		if err != nil {
			return
		}

		var repo *db.Repository
		repo, err = db.GetRepositoryByName(user.ID, ids[1].(string))
		if err != nil {
			return
		}

		payload, err = db.GetWebhookOfRepoByID(repo.ID, ids[2].(int64))
	}

	return
}

func GetContext(c *context.Context) {
	itemType := c.Params(":itemType")
	args := strings.Split(c.Params("*"), "/")

	payload, err := itemFromIDs(itemType, args)

	// An error occurred somewhere, relay that error onward
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Success!
	c.JSONSuccess(payload)
}