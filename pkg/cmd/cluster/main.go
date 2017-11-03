// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package cluster

import (
	"fmt"
	"log"
	"net"

	pbempty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"openpitrix.io/openpitrix/pkg/config"
	db "openpitrix.io/openpitrix/pkg/db/cluster"
	pb "openpitrix.io/openpitrix/pkg/service.pb"
)

func Main(cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.ClusterService.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClusterServiceServer(grpcServer, NewClusterServer(&cfg.Database))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type ClusterServer struct {
	db *db.ClusterDatabase
}

func NewClusterServer(cfg *config.Database) *ClusterServer {
	db, err := db.OpenClusterDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &ClusterServer{
		db: db,
	}
}

func (p *ClusterServer) GetCluster(ctx context.Context, args *pb.ClusterId) (reply *pb.Cluster, err error) {
	result, err := p.db.GetCluster(ctx, args.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "GetCluster: %v", err)
	}
	reply = To_proto_Cluster(nil, result)
	return
}

func (p *ClusterServer) GetClusterList(ctx context.Context, args *pb.ClusterListRequest) (reply *pb.ClusterListResponse, err error) {
	if args.PageNumber <= 0 {
		args.PageNumber = 1
	}
	if args.PageSize <= 0 {
		args.PageSize = 10
	}

	result, err := p.db.GetClusterList(ctx)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "GetClusterList: %v", err)
	}

	items := To_proto_ClusterList(result, int(args.PageNumber), int(args.PageSize))
	reply = &pb.ClusterListResponse{
		Items:       items,
		TotalItems:  int32(len(result)),
		TotalPages:  int32((len(result) + int(args.PageSize) - 1) / int(args.PageSize)),
		PageSize:    args.PageSize,
		CurrentPage: int32(len(result)/int(args.PageSize)) + 1,
	}

	return
}

func (p *ClusterServer) CreateCluster(ctx context.Context, args *pb.Cluster) (reply *pbempty.Empty, err error) {
	err = p.db.CreateCluster(ctx, To_database_Cluster(nil, args))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "CreateCluster: %v", err)
	}

	reply = &pbempty.Empty{}
	return
}

func (p *ClusterServer) UpdateCluster(ctx context.Context, args *pb.Cluster) (reply *pbempty.Empty, err error) {
	err = p.db.UpdateCluster(ctx, To_database_Cluster(nil, args))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "UpdateCluster: %v", err)
	}

	reply = &pbempty.Empty{}
	return
}

func (p *ClusterServer) DeleteCluster(ctx context.Context, args *pb.ClusterId) (reply *pbempty.Empty, err error) {
	err = p.db.DeleteCluster(ctx, args.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "DeleteCluster: %v", err)
	}

	reply = &pbempty.Empty{}
	return
}