package models

import (
	"time"

	hubtypes "github.com/sentinel-official/hub/types"

	"github.com/sentinel-official/health-check/types"
)

type Record struct {
	Addr                    string          `json:"addr,omitempty" bson:"addr,omitempty"`
	ClientConfig            []byte          `json:"client_config,omitempty" bson:"client_config,omitempty"`
	ConfigExchangeError     string          `json:"config_exchange_error,omitempty" bson:"config_exchange_error,omitempty"`
	ConfigExchangeTimestamp time.Time       `json:"config_exchange_timestamp,omitempty" bson:"config_exchange_timestamp,omitempty"`
	GigabytePrice           int64           `json:"gigabyte_price,omitempty" bson:"gigabyte_price,omitempty"`
	InfoFetchError          string          `json:"info_fetch_error,omitempty" bson:"info_fetch_error,omitempty"`
	InfoFetchTimestamp      time.Time       `json:"info_fetch_timestamp,omitempty" bson:"info_fetch_timestamp,omitempty"`
	LocationFetchError      string          `json:"location_fetch_error,omitempty" bson:"location_fetch_error,omitempty"`
	LocationFetchTimestamp  time.Time       `json:"location_fetch_timestamp,omitempty" bson:"location_fetch_timestamp,omitempty"`
	RemoteURL               string          `json:"remote_url,omitempty" bson:"remote_url,omitempty"`
	ServerConfig            []byte          `json:"server_config,omitempty" bson:"server_config,omitempty"`
	SessionID               uint64          `json:"session_id,omitempty" bson:"session_id,omitempty"`
	Status                  hubtypes.Status `json:"status,omitempty" bson:"status,omitempty"`
	SubscriptionID          uint64          `json:"subscription_id,omitempty" bson:"subscription_id,omitempty"`
	Type                    types.NodeType  `json:"type,omitempty" bson:"type,omitempty"`
}
