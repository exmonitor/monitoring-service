package icmp

import (
	"encoding/json"
	"fmt"
	"github.com/exmonitor/exclient/database"
	"github.com/exmonitor/exclient/database/spec/service"
	"github.com/exmonitor/exlogger"
	"github.com/pkg/errors"
	"time"
)

/*
Example metadata:
{
	"id": 2,
	"target": "101.102.103.104",
	"timeout": 2,
}
*/

type RawCheck struct {
	Id      int    `json:"id"`
	Target  string `json:"target"`
	Timeout int    `json:"timeout"`
}

func ParseCheck(service *service.Service, dbClient database.ClientInterface, logger *exlogger.Logger) (*Check, error) {
	var rawCheck RawCheck
	err := json.Unmarshal([]byte(service.Metadata), &rawCheck)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to parse ICMP json metadata for check id %d", service.ID))
	} else {
		logger.LogDebug("Successfully parsed ICMP json metadata for check id %d", service.ID)
	}

	checkConfig := CheckConfig{
		Id:            service.ID,
		FailThreshold: service.FailThreshold,
		Interval:      service.Interval,
		Target:        rawCheck.Target,
		Timeout:       time.Second * time.Duration(rawCheck.Timeout),
		Logger:        logger,
		DBClient:      dbClient,
	}

	return NewCheck(checkConfig)
}
