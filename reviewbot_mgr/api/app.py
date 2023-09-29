import os
from google.protobuf.json_format import MessageToDict
import grpc
from flask import Flask, request, jsonify
from reviewbot_pb2 import (DeleteTemplateRequest, ListTemplatesRequest,
                           MessageTemplate, SendMessageRequest)
from reviewbot_pb2_grpc import MessageStub, TemplatesStub

# Initialize gRPC client
GRPC_HOST = os.getenv('GRPC_HOST', 'localhost')
GRPC_PORT = os.getenv('GRPC_PORT', '50051')
GRPC_TIMEOUT = 3
channel = grpc.insecure_channel(f"{GRPC_HOST}:{GRPC_PORT}")
template_service = TemplatesStub(channel)
message_service = MessageStub(channel)

# Initialize Flask app
app = Flask(__name__)

@app.route('/')
def index():
    return MessageToDict(template_service.List(ListTemplatesRequest(), timeout=GRPC_TIMEOUT))

@app.route('/health')
def health():
    return jsonify({'status': 'ok'})

@app.route('/message', methods=['POST'])
def message():
    """Send a message to a customer."""
    req = parse_message_request(request.json)
    return MessageToDict(message_service.Send(req, timeout=GRPC_TIMEOUT))

@app.route('/template', methods=['GET', 'POST', 'PUT', 'DELETE'])
def template():
    """Create, update, list, or delete templates."""
    if request.method in ['POST', 'PUT']:
        templ_req = parse_template_request(request.json)
        if templ_req is None:
            return jsonify({'message': 'Invalid request'}), 400

    if request.method == 'POST':
        return MessageToDict(template_service.Create(templ_req, timeout=GRPC_TIMEOUT))
    elif request.method == 'PUT':
        return MessageToDict(template_service.Update(templ_req, timeout=GRPC_TIMEOUT))
    elif request.method == 'DELETE':
        del_req = DeleteTemplateRequest(name=request.args.get('name'))
        return MessageToDict(template_service.Delete(del_req, timeout=GRPC_TIMEOUT))
    else:
        return MessageToDict(template_service.List(ListTemplatesRequest(), timeout=GRPC_TIMEOUT))

def parse_message_request(data: dict):
    """Parse request data to send a message to a customer."""
    if data.get('first_name') is None:
        return None
    if data.get('last_name') is None:
        return None
    if data.get('template_name') is None:
        return None

    return SendMessageRequest(
        first_name=data['first_name'],
        last_name=data['last_name'],
        template_name=data['template_name'],
    )

def parse_template_request(data: dict):
    """Parse request data to create or update a template."""
    if data.get('name') is None:
        return None
    if data.get('content') is None:
        return None
    return MessageTemplate(
        name=data['name'],
        content=data['content'],
        triggers=data.get('triggers', []),
    )

@app.errorhandler(grpc.RpcError)
def handle_grpc_error(e: grpc.RpcError):
    """Handle a gRPC error."""
    if e.code() == grpc.StatusCode.INVALID_ARGUMENT or "Exception serializing request!" in e.details():
        return jsonify({'message': e.details()}), 400
    elif 'has not started a chat' in e.details():
        return jsonify({'message': e.details()}), 403
    elif e.code() == grpc.StatusCode.NOT_FOUND:
        return jsonify({'message': e.details()}), 404
    elif e.code() == grpc.StatusCode.ALREADY_EXISTS:
        return jsonify({'message': e.details()}), 409
    else:
        return jsonify({'message': e.details()}), 500
