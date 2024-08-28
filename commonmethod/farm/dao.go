package farm

import (
	"github.com/coinmeca/go-common/commonmethod/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recent struct {
	ChainId  string               `json:"chainId" bson:"chainId"`
	Address  string               `json:"address" bson:"address"`
	Time     int64                `json:"time" bson:"time"`
	Type     int                  `json:"type" bson:"type"`
	User     string               `json:"owner" bson:"owner"`
	Amount   primitive.Decimal128 `json:"amount" bson:"amount"`
	Share    primitive.Decimal128 `json:"share" bson:"share"`
	TxHash   string               `json:"txHash" bson:"txHash"`
	UpdateAt string               `json:"updateAt" bson:"updateAt"`
}

type Volume struct {
	Amount primitive.Decimal128 `json:"amount" bson:"amount"`
	Value  primitive.Decimal128 `json:"value" bson:"value"`
}

type Value struct {
	Stake primitive.Decimal128 `json:"stake" bson:"stake"`
	Earn  primitive.Decimal128 `json:"earn" bson:"earn"`
}

type Chart struct {
	ChainId  string               `json:"chainId" bson:"chainId"`
	Address  string               `json:"address" bson:"address"`
	Time     int64                `json:"time" bson:"time"`
	Staked   Volume               `json:"staked" bson:"staked"`
	Interest Volume               `json:"interest" bson:"interest"`
	Value    Value                `json:"value" bson:"value"`
	Apr      primitive.Decimal128 `json:"apr" bson:"apr"`
	Earned   primitive.Decimal128 `json:"earned" bson:"earned"`
	Total    primitive.Decimal128 `json:"total" bson:"total"`
}

type Last struct {
	ChainId            string               `json:"chainId" bson:"chainId"`
	Address            string               `json:"address" bson:"address"`
	Staked             primitive.Decimal128 `json:"staked" bson:"staked"`
	StakedChange       primitive.Decimal128 `json:"stakedChange" bson:"stakedChange"`
	ValueStaked        primitive.Decimal128 `json:"valueStaked" bson:"valueStaked"`
	ValueStakedChange  primitive.Decimal128 `json:"valueStakedChange" bson:"valueStakedChange"`
	Interest           primitive.Decimal128 `json:"interest" bson:"interest"`
	InterestChange     primitive.Decimal128 `json:"interestChange" bson:"interestChange"`
	Staking24h         primitive.Decimal128 `json:"staking24h" bson:"staking24h"`
	Staking24hChange   primitive.Decimal128 `json:"staking24hChange" bson:"staking24hChange"`
	UnStaking24h       primitive.Decimal128 `json:"unStaking24h" bson:"unStaking24h"`
	UnStaking24hChange primitive.Decimal128 `json:"unStaking24hChange" bson:"unStaking24hChange"`
	Interest24h        primitive.Decimal128 `json:"interest24h" bson:"interest24h"`
	Interest24hChange  primitive.Decimal128 `json:"interest24hChange" bson:"interest24hChange"`
	ValueLocked        primitive.Decimal128 `json:"valueLocked" bson:"valueLocked"`
	Chart              Chart                `json:"chart" bson:"chart"`
	Recent             Recent               `json:"recent" bson:"recent"`
}

type Farm struct {
	ChainId            string               `json:"chainId" bson:"chainId"`
	Address            string               `json:"address" bson:"address"`
	Id                 int64                `json:"id" bson:"id"`
	Main               string               `json:"main" bson:"main"`
	Master             string               `json:"master" bson:"master"`
	Name               string               `json:"name" bson:"name"`
	Symbol             string               `json:"symbol" bson:"symbol"`
	Decimals           int64                `json:"decimals" bson:"decimals"`
	Stake              token.Token          `json:"stake" bson:"stake"`
	Earn               token.Token          `json:"earn" bson:"earn"`
	Start              int64                `json:"start" bson:"start"`
	Period             int64                `json:"period" bson:"period"`
	Duration           int64                `json:"duration" bson:"duration"`
	Goal               int64                `json:"goal" bson:"goal"`
	Staked             primitive.Decimal128 `json:"staked" bson:"staked"`
	StakedChange       primitive.Decimal128 `json:"stakedChange" bson:"stakedChange"`
	ValueStaked        primitive.Decimal128 `json:"valueStaked" bson:"valueStaked"`
	ValueStakedChange  primitive.Decimal128 `json:"valueStakedChange" bson:"valueStakedChange"`
	Interest           primitive.Decimal128 `json:"interest" bson:"interest"`
	InterestChange     primitive.Decimal128 `json:"interestChange" bson:"interestChange"`
	Staking24h         primitive.Decimal128 `json:"staking24h" bson:"staking24h"`
	Staking24hChange   primitive.Decimal128 `json:"staking24hChange" bson:"staking24hChange"`
	UnStaking24h       primitive.Decimal128 `json:"unStaking24h" bson:"unStaking24h"`
	UnStaking24hChange primitive.Decimal128 `json:"unStaking24hChange" bson:"unStaking24hChange"`
	Interest24h        primitive.Decimal128 `json:"interest24h" bson:"interest24h"`
	Interest24hChange  primitive.Decimal128 `json:"interest24hChange" bson:"interest24hChange"`
	ValueLocked        primitive.Decimal128 `json:"valueLocked" bson:"valueLocked"`
	Total              primitive.Decimal128 `json:"total" bson:"total"`
	Claimable          primitive.Decimal128 `json:"claimable" bson:"claimable"`
	Apr                primitive.Decimal128 `json:"apr" bson:"apr"`
	AprChange          primitive.Decimal128 `json:"aprChange" bson:"aprChange"`
	Chart              []Chart              `json:"chart" bson:"chart"`
	Recents            []Recent             `json:"recents" bson:"recents"`
	Last               Last                 `json:"last" bson:"bson"`
}
