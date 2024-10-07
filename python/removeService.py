import grpc

import bgremover_pb2 as bgremover__pb2

import bgremover_pb2_grpc as bgremover__pb

from concurrent import futures



class RemoveServicer(bgremover__pb.RemoveServicer):
    def RemoveBG(self, request, context):
        print("Received request")
        return bgremover__pb2.ImageResponse(ProcessedImage=request.Image)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    bgremover__pb.add_RemoveServicer_to_server(RemoveServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

serve()
