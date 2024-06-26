package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	dstore_client "github.com/brotherlogic/dstore/client"
	githubridgeclient "github.com/brotherlogic/githubridge/client"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/tasklist/proto"
)

// Server main server type
type Server struct {
	*goserver.GoServer
	dclient  *dstore_client.DStoreClient
	ghclient githubridgeclient.GithubridgeClient
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
	}
	s.dclient = &dstore_client.DStoreClient{Gs: s.GoServer}
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	password, err := os.ReadFile(fmt.Sprintf("%v/.ghb", dirname))
	if err == nil {
		client, err := githubridgeclient.GetClientExternal(string(password))
		if err != nil {
			panic(err)
		}
		s.ghclient = client
	}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterTaskListServiceServer(server, s)
	pbgh.RegisterGithubSubscriberServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{
		&pbg.State{Key: "magic", Value: int64(12345)},
	}
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.PrepServer("tasklist")
	server.Register = server

	err := server.RegisterServerV2(false)
	if err != nil {
		return
	}

	go func() {
		for {
			ctx, cancel := utils.ManualContext("tasklist-init", time.Minute)
			_, err := server.ValidateTaskLists(ctx, &pb.ValidateTaskListsRequest{})
			server.CtxLog(ctx, fmt.Sprintf("Validated task list: %v", err))
			cancel()
			time.Sleep(time.Hour)
		}
	}()

	fmt.Printf("%v", server.Serve())
}
