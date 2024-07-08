package rpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
	"tiktokrpc/cmd/api/pkg/errmsg"
	pb "tiktokrpc/cmd/user/pkg/ai/rpc/client"
	"time"
)

func GerVector(path string) ([]float32, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewExampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.Request{Message: path}
	res, err := client.SendMessage(ctx, req)
	if err != nil {
		return nil, errmsg.DatabaseError.WithMessage(err.Error())
	}
	return convertStringToFloatArray(res.Reply), nil
}

func convertStringToFloatArray(input string) []float32 {
	// 去掉字符串两边的方括号
	input = strings.Trim(input, "[]")
	// 使用逗号分割字符串
	parts := strings.Split(input, ",")
	// 创建一个 float32 切片来存储结果
	result := make([]float32, len(parts))
	// 将每个字符串元素转换为 float32
	for i, part := range parts {
		// 去除字符串两边的空格
		part = strings.TrimSpace(part)
		num, err := strconv.ParseFloat(part, 32)
		if err != nil {
			// 如果解析失败，返回 nil
			return nil
		}
		result[i] = float32(num)
	}
	return result
}
