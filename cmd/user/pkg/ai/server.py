import grpc
from concurrent import futures
import test_pb2
import test_pb2_grpc
import train

class ExampleServicer(test_pb2_grpc.ExampleServiceServicer):
    def SendMessage(self, request, context):
        return test_pb2.Response(reply=train.extract(request.message))

def run_server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    test_pb2_grpc.add_ExampleServiceServicer_to_server(ExampleServicer(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    run_server()