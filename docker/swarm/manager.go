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
	InitWorker(token string, remoteAddr string) error
	GetToken() (string, error)
}

func (m *swarmManager) InitMaster(advertiseAddr string) (string, error) {
	r := s.InitRequest{
		AdvertiseAddr: advertiseAddr,
		ListenAddr:    "0.0.0.0:2377",
	}
	return m.cli.SwarmInit(m.ctx, r)
}

func (m *swarmManager) InitWorker(token string, remoteAddr string) error {
	r := s.JoinRequest{
		JoinToken:   token,
		ListenAddr:  "0.0.0.0:2377",
		RemoteAddrs: []string{remoteAddr},
	}
	return m.cli.SwarmJoin(m.ctx, r)
}

func (m *swarmManager) GetToken() (string, error) {
	s, err := m.cli.SwarmInspect(m.ctx)
	if err != nil {
		return "", err
	}
	return s.JoinTokens.Manager, nil
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
