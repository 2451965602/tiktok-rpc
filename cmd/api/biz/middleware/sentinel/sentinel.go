package sentinel

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"tiktokrpc/cmd/api/pkg/constants"
	"tiktokrpc/cmd/api/pkg/errmsg"
)

func Init() error {

	err := sentinel.InitDefault()
	if err != nil {
		panic(err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "default",
			Threshold:              constants.SentinelThreshold,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       constants.SentinelStatIntervalInMs,
		},
	})
	if err != nil {
		return errmsg.ServiceError.WithMessage(err.Error())
	}

	return nil
}
