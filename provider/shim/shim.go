package shim

import (
	"github.com/StatusCakeDev/terraform-provider-statuscake/v2/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider(version string) *schema.Provider {
	return provider.Provider()
}