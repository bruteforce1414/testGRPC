// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"github.com/nori-io/nori-common/meta"
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
func (s *server) GetMetaInfo(ctx context.Context, req *pb.MetaDataRequest) (*pb.ArrayMetaDataReply, error) {
	var arrayPlugins pb.ArrayMetaDataReply

	for _, v := range exampleData {
		var deps []*pb.Dependency

		for _, d := range v.Dependencies {
			deps = append(deps, &pb.Dependency{
				ID: &pb.PluginID{
					MetaId:               string(d.ID),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				Constraint: d.Constraint,
				Interface: &pb.Interface{
					Interface:            string(d.Interface),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})

		}

		var license *pb.License

		for _, l := range v.License {
			license = &pb.License{
				Title:                l.Title,
				Type:                 l.Type,
				URI:                  l.URI,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

		}

		var links *pb.Links

		for _, l := range v.Links {
			links = &pb.Links{
				Link: []*pb.Link{&pb.Link{
					Title:                l.Title,
					URL:                  l.URL,
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				}},
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

		}

		repository := &pb.Repository{
			Type:                 v.Repository.Type,
			URI:                  v.Repository.URI,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}

		if (req.Id == string(v.ID.ID)) && (req.Version == string(v.ID.Version)) {
			arrayPlugins.ArrayMetaDataReply = append(arrayPlugins.ArrayMetaDataReply, &pb.MetaDataReply{
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
					MetaDependency:       deps,
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				Description: &pb.Description{
					Name:                 v.Description.Name,
					Description:          v.Description.Description,
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				Core: &pb.Core{
					VersionConstraint:    v.Core.VersionConstraint,
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				Interface: &pb.Interface{
					Interface:            string(v.Interface),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				License:              license,
				Links:                links,
				Repository:           repository,
				Tags:                 v.Tags,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})

		}
	}

	log.Printf("Received: %v", req)
	return &arrayPlugins, nil
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
		Dependencies: []meta.Dependency{},
		Description: meta.Description{
			Name:        "Nori Session",
			Description: "Nori: Session Interface",
		},
		Interface: meta.NewInterface("nori/empty", "0.0.1"),
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
		Dependencies: []meta.Dependency{},
		Description: meta.Description{
			Name:        "Nori Session",
			Description: "Nori: Session Interface",
		},
		Interface: meta.NewInterface("nori/empty", "0.0.2"),
		License: []meta.License{{
			Title: "",
			Type:  "GPLv3",
			URI:   "https://www.gnu.org/licenses/",
		},
		},
		Tags: []string{"session"},
	},
}
