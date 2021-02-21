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

func itemFromItemID(itemType string, id int64) (payload interface{}, err error) {
	switch itemType {
	case "ssh_key":
		payload, err = db.GetPublicKeyByID(id)
	case "email":
		payload, err = db.GetEmailByID(id)
	case "repository":
		payload, err = db.GetRepositoryByID(id)
	default:
		err = fmt.Errorf("Unknown item type: %s", itemType)
	}
	return
}

func itemFromUserItemID(itemType string, userID, id int64) (payload interface{}, err error) {
	switch itemType {
	case "app_token":
		payload, err = db.AccessTokens.GetByID(userID, id)
	default:
		err = fmt.Errorf("Unknown item type: %s", itemType)
	}

	return
}

func GetContext(c *context.Context) {
	itemType := c.Params(":itemType")
	args := strings.Split(c.Params("*"), "/")

	var payload interface{}
	var err error

	// The `args` should be an ID or a string identifier
	if len(args) == 1 {
		// Try parsing an `id` first
		if id, parse_err := strconv.ParseInt(args[0], 10, 64); parse_err == nil {
			payload, err = itemFromItemID(itemType, id)
		} else {
			// TODO: string identifier
		}
	} else if len(args) == 2 {
		// The `args` should be pair userID/ID
		userID, parse1_err := strconv.ParseInt(args[0], 10, 64)
		id, parse2_err := strconv.ParseInt(args[1], 10, 64)

		if parse1_err != nil || parse2_err != nil {
			err = fmt.Errorf("Unable to parse userID/id context combo: %v/%v", args[0], args[1])
		}

		payload, err = itemFromUserItemID(itemType, userID, id)
	} else {
		err = fmt.Errorf("Unknown context to retrieve: %v", args)
	}

	// An error occurred somewhere, relay that error onward
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Success!
	c.JSONSuccess(payload)
}
