import grpc

import bgremover_pb2 as bgremover__pb2

import bgremover_pb2_grpc as bgremover__pb

from concurrent import futures

from rembg import remove



class RemoveServicer(bgremover__pb.RemoveServicer):
    def RemoveBG(self, request_iterator, context):

        imageData = b''

        for chunk in request_iterator:
            imageData += chunk.Image

        output = remove(imageData,force_return_bytes=True)

        return bgremover__pb2.ImageResponse(ProcessedImage=output)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    bgremover__pb.add_RemoveServicer_to_server(RemoveServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

serve()
