// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nori-io/nori-common/meta"
	"github.com/nori-io/nori-interfaces/interfaces"

	//"github.com/nori-io/nori-interfaces/interfaces"
	"google.golang.org/grpc"

	pb "github.com/bruteforce1414/testGRPC/metainfo"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.MetaInfoServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetMetaInfo(ctx context.Context, req *pb.MetaDataRequest) (*pb.MetaDataReply, error) {

	for _, v := range exampleData {

		fmt.Println("v.Dependencies", v.Dependencies)
		if (req.Id == string(v.ID.ID)) && (req.Version == string(v.ID.Version)) {
			return &pb.MetaDataReply{
				MetaID: &pb.ID{
					PluginID: &pb.PluginID{
						MetaId:               v.ID.String(),
						XXX_NoUnkeyedLiteral: struct{}{},
						XXX_unrecognized:     nil,
						XXX_sizecache:        0,
					},
					Version:              v.ID.Version,
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				Author: &pb.Author{
					Name:                 v.Author.Name,
					URI:                  v.Author.URI,
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				DependenciesArray: &pb.Dependencies{
					MetaDependency:       nil,
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				Description:          nil,
				Core:                 nil,
				Interface:            nil,
				License:              nil,
				Links:                nil,
				Repository:           nil,
				Tags:                 nil,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}, nil

		}
	}

	log.Printf("Received: %v", req)
	return &pb.MetaDataReply{MetaID: nil}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMetaInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var exampleData = []meta.Data{
	meta.Data{
		ID: meta.ID{
			ID:      "nori/session",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io/",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{
			interfaces.CacheInterface.Dependency(),
		},
		Description: meta.Description{
			Name:        "Nori Session",
			Description: "Nori: Session Interface",
		},
		//Interface: interfaces.SessionInterface,
		License: []meta.License{{
			Title: "",
			Type:  "GPLv3",
			URI:   "https://www.gnu.org/licenses/",
		},
		},
		Tags: []string{"session"},
	},

	meta.Data{
		ID: meta.ID{
			ID:      "nori/session2",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io/",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{
			//interfaces.CacheInterface.Dependency(),
		},
		Description: meta.Description{
			Name:        "Nori Session",
			Description: "Nori: Session Interface",
		},
		//Interface: interfaces.SessionInterface,
		License: []meta.License{{
			Title: "",
			Type:  "GPLv3",
			URI:   "https://www.gnu.org/licenses/",
		},
		},
		Tags: []string{"session"},
	},
}
