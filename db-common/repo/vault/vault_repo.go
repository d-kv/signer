package vault

import (
	"context"
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/entity"
	"fmt"
	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type Repo struct {
	client *vault.Client
}

func New(conf config.VaultConfig) (error, *Repo) {
	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress(conf.Address),
		vault.WithRequestTimeout(conf.Timout),
	)
	if err != nil {
		return err, nil
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken(conf.Token); err != nil {
		return err, nil
	}

	return nil, &Repo{client: client}
}

// FindTokenByIntegrationId find token by integrationId
// ctx := context.Background()
// err, repo := vault.New(config.VaultConfig{
// Token:   "my-token",
// Address: "http://127.0.0.1:8200",
// Timout:  30 * time.Second,
// })
// if err != nil {
// log.Fatal(err)
// }
//
// err, token := repo.FindTokenByIntegrationId(ctx, "some_integration_id")
// if err != nil {
// log.Fatal(err)
// }
// log.Printf("Found integration: %s", token)
func (repo Repo) FindTokenByIntegrationId(ctx context.Context, id string) (error, *entity.IntegrationToken) {
	path := fmt.Sprintf("data/integration_token/%s", id)
	integrationToken, err := repo.client.Secrets.KvV2Read(ctx, path, vault.WithMountPath("secret"))
	if err != nil {
		return err, nil
	}
	return nil, &entity.IntegrationToken{
		IntegrationId: integrationToken.Data.Data["integration_id"].(string),
		Token:         integrationToken.Data.Data["token"].(string),
		KeyId:         integrationToken.Data.Data["key_id"].(string),
	}
}

func (repo Repo) SaveIntegrationToken(ctx context.Context, token *entity.IntegrationToken) error {
	path := fmt.Sprintf("data/integration_token/%s", token.IntegrationId)
	_, err := repo.client.Secrets.KvV2Write(ctx, path, schema.KvV2WriteRequest{Data: map[string]interface{}{
		"integration_id": token.IntegrationId,
		"token":          token.Token,
		"key_id":         token.KeyId,
	}}, vault.WithMountPath("secret"))
	return err
}
