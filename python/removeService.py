import grpc

import bgremover_pb2 as bgremover__pb2

import bgremover_pb2_grpc as bgremover__pb

from concurrent import futures

from rembg.bg import remove,new_session




class RemoveServicer(bgremover__pb.RemoveServicer):
    def RemoveBG(self, request_iterator, context):
        model_name = "u2netp"
        session = new_session(model_name)
        imageData = b''

        for chunk in request_iterator:
            imageData += chunk.Image

        output = remove(imageData,session=session,force_return_bytes=True)

        return bgremover__pb2.ImageResponse(ProcessedImage=output)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    bgremover__pb.add_RemoveServicer_to_server(RemoveServicer(), server)
    server.add_insecure_port('0.0.0.0:50051')
    server.start()
    server.wait_for_termination()

serve()
