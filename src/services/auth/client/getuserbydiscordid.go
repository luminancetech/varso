package client

import (
	"fmt"

	"github.com/luminancetech/varso/src/common/api"
	"github.com/luminancetech/varso/src/services/auth/client/models"
)

func (c *Client) GetUserByDiscordId(discordId string) (*models.GetUserByDiscordIdResponse, *api.Error) {
	return api.SendRequest[any, models.GetUserByDiscordIdResponse](c.Client, c.AuthHeader, "GET", fmt.Sprintf("auth/user/admininspect/%s", discordId), nil)
}
