# !/usr/bin/env python
# -*- coding: utf-8 -*-

# @FileName: client.py
# @Time    : 2024/4/28  17:47
# @Author  : zcc

import grpc
import test_pb2
import test_pb2_grpc

def run_client():
    with grpc.insecure_channel('localhost:50052') as channel:
        stub = test_pb2_grpc.ExampleServiceStub(channel)
        response = stub.SendMessage(test_pb2.Request(message='/home/fzuer/Downloads/code/707cebd56c51a51b743582122ff39302.jpg'))
        print(f"Response from server: {response.reply}")

if __name__ == '__main__':
    run_client()

