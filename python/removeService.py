import grpc

import bg_remover_pb2 as bgremover__pb2

import bg_remover_pb2_grpc as bgremover__pb




class RemoveServiceServicer(bgremover__pb.RemoveServiceServicer):
    def removeBG(self,request,context):
        print("Received request")
        print(request.image)
        print(request.image.shape)
        print(request.image.dtype)
        print(request.image.size)
        print(request.image.nbytes)
        print(request.image.tobytes())
        print("Received request")
        return bgremover__pb2.ImageResponse(processedImage=request.image)
    

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    bgremover__pb2_grpc.add_RemoveServiceServicer_to_server(RemoveServiceServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()
