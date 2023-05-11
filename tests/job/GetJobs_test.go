package main

import (
	"context"
	pb "scow-slurm-adapter/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetJobs(t *testing.T) {

	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:8972", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewJobServiceClient(conn)

	// Call the Add RPC with test data
	user := []string{"test03", "test02"}
	account := []string{"c_admin", "a_admin"}
	state := []string{"CANCELLED", "PENDDING"}
	req := &pb.GetJobsRequest{
		Fields: []string{}, 
		// Filter: &pb.GetJobsRequest_Filter{Users: user, Accounts: account, States: state, EndTime: &pb.TimeRange{StartTime: &timestamppb.Timestamp{Seconds: 1682066342}, EndTime: &timestamppb.Timestamp{Seconds: 1682586485}}}, PageInfo: &pb.PageInfo{Page: 1, PageSize: 10},
		Filter: &pb.GetJobsRequest_Filter{Users: user, Accounts: account, States: state, EndTime: &pb.TimeRange{StartTime: &timestamppb.Timestamp{Seconds: 1682066342}, EndTime: &timestamppb.Timestamp{Seconds: 1682586485}}}, 

	}
	res, err := client.GetJobs(context.Background(), req)
	if err != nil {
		t.Fatalf("GetJobs failed: %v", err)
	}

	// Check the result, 通过判断错误为nil 来决定是否执行成功
	// assert.Empty(t, err)
	assert.IsType(t, []*pb.JobInfo{}, res.Jobs)
}
