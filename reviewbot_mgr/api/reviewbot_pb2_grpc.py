# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import reviewbot_pb2 as reviewbot__pb2


class TemplatesStub(object):
    """Template management service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/Templates/Create',
                request_serializer=reviewbot__pb2.MessageTemplate.SerializeToString,
                response_deserializer=reviewbot__pb2.TemplateResponse.FromString,
                )
        self.Update = channel.unary_unary(
                '/Templates/Update',
                request_serializer=reviewbot__pb2.MessageTemplate.SerializeToString,
                response_deserializer=reviewbot__pb2.TemplateResponse.FromString,
                )
        self.Delete = channel.unary_unary(
                '/Templates/Delete',
                request_serializer=reviewbot__pb2.DeleteTemplateRequest.SerializeToString,
                response_deserializer=reviewbot__pb2.TemplateResponse.FromString,
                )
        self.List = channel.unary_unary(
                '/Templates/List',
                request_serializer=reviewbot__pb2.ListTemplatesRequest.SerializeToString,
                response_deserializer=reviewbot__pb2.TemplateResponse.FromString,
                )


class TemplatesServicer(object):
    """Template management service.
    """

    def Create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def List(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TemplatesServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=reviewbot__pb2.MessageTemplate.FromString,
                    response_serializer=reviewbot__pb2.TemplateResponse.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=reviewbot__pb2.MessageTemplate.FromString,
                    response_serializer=reviewbot__pb2.TemplateResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=reviewbot__pb2.DeleteTemplateRequest.FromString,
                    response_serializer=reviewbot__pb2.TemplateResponse.SerializeToString,
            ),
            'List': grpc.unary_unary_rpc_method_handler(
                    servicer.List,
                    request_deserializer=reviewbot__pb2.ListTemplatesRequest.FromString,
                    response_serializer=reviewbot__pb2.TemplateResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Templates', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Templates(object):
    """Template management service.
    """

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Templates/Create',
            reviewbot__pb2.MessageTemplate.SerializeToString,
            reviewbot__pb2.TemplateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Templates/Update',
            reviewbot__pb2.MessageTemplate.SerializeToString,
            reviewbot__pb2.TemplateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Templates/Delete',
            reviewbot__pb2.DeleteTemplateRequest.SerializeToString,
            reviewbot__pb2.TemplateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def List(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Templates/List',
            reviewbot__pb2.ListTemplatesRequest.SerializeToString,
            reviewbot__pb2.TemplateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class MessageStub(object):
    """Send a message to a customer.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Send = channel.unary_unary(
                '/Message/Send',
                request_serializer=reviewbot__pb2.SendMessageRequest.SerializeToString,
                response_deserializer=reviewbot__pb2.SendMessageResponse.FromString,
                )


class MessageServicer(object):
    """Send a message to a customer.
    """

    def Send(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MessageServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Send': grpc.unary_unary_rpc_method_handler(
                    servicer.Send,
                    request_deserializer=reviewbot__pb2.SendMessageRequest.FromString,
                    response_serializer=reviewbot__pb2.SendMessageResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Message', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Message(object):
    """Send a message to a customer.
    """

    @staticmethod
    def Send(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Message/Send',
            reviewbot__pb2.SendMessageRequest.SerializeToString,
            reviewbot__pb2.SendMessageResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)