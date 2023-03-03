package swarm

import (
	"context"

	s "github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

type SwarmClient interface {
	client.SwarmAPIClient
}

type SwarmManager interface {
	InitMaster(advertiseAddr string) (string, error)
	InitWorker(token string) error
}

func (m *swarmManager) InitMaster(advertiseAddr string) (string, error) {
	r := s.InitRequest{AdvertiseAddr: advertiseAddr}
	return m.cli.SwarmInit(m.ctx, r)
}

func (m *swarmManager) InitWorker(token string) error {
	r := s.JoinRequest{JoinToken: token}
	return m.cli.SwarmJoin(m.ctx, r)
}

type swarmManager struct {
	cli SwarmClient
	ctx context.Context
}

func NewSwarmManager(ctx context.Context, client *client.Client) SwarmManager {
	return &swarmManager{
		cli: client,
		ctx: ctx,
	}
}
