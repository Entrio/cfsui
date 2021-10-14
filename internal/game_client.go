package internal

import (
	"context"
	"fmt"
	gen "github.com/Entrio/cfsui/internal/proto"
	"github.com/rs/zerolog/log"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:43567"
	defaultName = "UnknownEther"
)

type (
	GameClient struct {
		pubClient  gen.CFSPublicClient
		logChan    chan string
		updateTime *time.Ticker
	}
)

func NewGameClient() (*GameClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to remote server")
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	gc := &GameClient{
		pubClient:  gen.NewCFSPublicClient(conn),
		logChan:    make(chan string, 10),
		updateTime: time.NewTicker(time.Second * 2),
	}

	res, err := gc.pubClient.GetServerInfo(ctx, &gen.ServerInfoRequest{})

	if err != nil {
		return nil, err
	}

	gc.sendLog(fmt.Sprintf("Connected to %s", res.GetName()))

	return gc, nil
}

func (gc *GameClient) startup() {
	go gc.update()
}

func (gc *GameClient) update() {
	for {
		select {
		case <-gc.updateTime.C:
			gc.fetchFarms()
		}
	}
}

func (gc *GameClient) AddFarm() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	_, err := gc.pubClient.CreateFarm(ctx, &gen.Empty{})
	if err != nil {
		gc.sendLog(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	gc.sendLog("Farm request sent")
}

func (gc *GameClient) fetchFarms() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := gc.pubClient.GetFarms(ctx, &gen.Empty{})

	if err != nil {
		gc.sendLog(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	gc.sendLog(fmt.Sprintf("Fetched %d farms", len(res.Farms)))
}

func (gc *GameClient) sendLog(data string) {
	gc.logChan <- data
}
