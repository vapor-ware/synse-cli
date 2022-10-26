package view

import (
	"context"
	"github.com/vapor-ware/synse-cli/pkg/client"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"time"
)

type SynsePoller struct {
	*client.APIClient
	ctx             context.Context
	refreshInterval time.Duration
}

type Response struct {
	Data []*synse.V3Reading
	Err  error
}

func NewSynsePoller(ctx context.Context, client *client.APIClient, interval time.Duration) *SynsePoller {
	return &SynsePoller{
		ctx:             ctx,
		APIClient:       client,
		refreshInterval: interval,
	}
}

func (p *SynsePoller) Update() chan Response {
	tick := time.NewTicker(time.Second * refreshInterval).C
	ch := make(chan Response)
	defer close(ch)
	resp := Response{}
	for {
		select {
		case <-tick:
			data, err := p.APIClient.Readings()
			if err != nil {
				resp.Err = err
				ch <- resp
				break
			}
			resp.Data, resp.Err = data, nil
			return ch
		}
	}
}
