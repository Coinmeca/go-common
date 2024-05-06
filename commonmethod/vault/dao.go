package model

import (
	"coinmeca-trader/conf"
	"context"

	"github.com/coinmeca/go-common/commonmethod/vault"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/coinmeca/go-common/commonlog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type VaultDB struct {
	config *conf.Config

	client       *mongo.Client
	colHistory   *mongo.Collection
	colTokenInfo *mongo.Collection
	colVaultInfo *mongo.Collection

	start chan struct{}
}

func NewVaultDB(config *conf.Config, root *Repositories) (IRepository, error) {
	r := &VaultDB{
		config: config,
		start:  make(chan struct{}),
	}

	var err error
	credential := options.Credential{
		Username: r.config.Repositories["vaultDB"]["username"].(string),
		Password: r.config.Repositories["vaultDB"]["pass"].(string),
	}

	clientOptions := options.Client().ApplyURI(config.Repositories["vaultDB"]["datasource"].(string)).SetAuth(credential)
	if r.client, err = mongo.Connect(context.Background(), clientOptions); err != nil {
		return nil, err
	}

	if err = r.client.Ping(context.Background(), nil); err == nil {
		db := r.client.Database(config.Repositories["vaultDB"]["db"].(string))
		r.colHistory = db.Collection("history")
		r.colTokenInfo = db.Collection("token_info")
		r.colVaultInfo = db.Collection("vault_info")
	} else {
		return nil, err
	}

	commonlog.Logger.Debug("load repository",
		zap.String("vaultDB", r.config.Common.ServiceId),
	)
	return r, nil
}

func (v *VaultDB) Start() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		close(v.start)
		return
	}()
}

// TODO: add deposit, withdraw event history
func (v *VaultDB) SaveEventHistory(recent vault.Recent) error {
	filter := bson.M{"txHash": recent.TxHash}
	update := bson.M{
		"$set": bson.M{
			"time":     recent.Time,
			"type":     recent.Type,
			"user":     recent.User,
			"token":    recent.Token,
			"amount":   recent.Amount,
			"meca":     recent.Meca,
			"share":    recent.Share,
			"txHash":   recent.TxHash,
			"updateAt": recent.UpdateAt,
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := v.colHistory.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (v *VaultDB) UpdateGetAllData(vaults []vault.OutputVault) error {
	for _, data := range vaults {
		convertedData := bson.M{
			"key":      data.Key,
			"address":  data.Address.Hex(),
			"name":     data.Name,
			"symbol":   data.Symbol,
			"decimals": data.Decimals,
			"locked":   data.Locked.String(),
			"exchange": data.Exchange.String(),
			"rate":     data.Rate.String(),
			"weight":   data.Weight.String(),
			"need":     data.Need.String(),
		}

		_, err := v.colTokenInfo.UpdateOne(
			context.Background(),
			bson.M{"address": data.Address.Hex()},
			bson.M{"$set": convertedData},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *VaultDB) GetVault(address string) (vault.Vault, error) {
	filter := bson.M{"address": address}

	var result vault.VAult
	err := f.colTokenInfo.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}
