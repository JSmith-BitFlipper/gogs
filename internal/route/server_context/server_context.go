package server_context

import (
	"fmt"
	"net/http"

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

func ItemFromItemID(c *context.Context) {
	itemType := c.Params(":itemType")
	id := c.ParamsInt64(":id")

	var payload interface{}
	var err error

	switch itemType {
	case "ssh_key":
		payload, err = db.GetPublicKeyByID(id)
	case "email":
		payload, err = db.GetEmailByID(id)
	default:
		err = fmt.Errorf("Unknown item type: %s", itemType)
	}

	// An error occurred somewhere, relay that error onward
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Success!
	c.JSONSuccess(payload)
}
