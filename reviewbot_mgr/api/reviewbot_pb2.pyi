from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MessageTemplate(_message.Message):
    __slots__ = ["name", "content", "triggers"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    TRIGGERS_FIELD_NUMBER: _ClassVar[int]
    name: str
    content: str
    triggers: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, name: _Optional[str] = ..., content: _Optional[str] = ..., triggers: _Optional[_Iterable[str]] = ...) -> None: ...

class DeleteTemplateRequest(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class ListTemplatesRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class TemplateResponse(_message.Message):
    __slots__ = ["templates"]
    TEMPLATES_FIELD_NUMBER: _ClassVar[int]
    templates: _containers.RepeatedCompositeFieldContainer[MessageTemplate]
    def __init__(self, templates: _Optional[_Iterable[_Union[MessageTemplate, _Mapping]]] = ...) -> None: ...

class SendMessageRequest(_message.Message):
    __slots__ = ["first_name", "last_name", "template_name"]
    FIRST_NAME_FIELD_NUMBER: _ClassVar[int]
    LAST_NAME_FIELD_NUMBER: _ClassVar[int]
    TEMPLATE_NAME_FIELD_NUMBER: _ClassVar[int]
    first_name: str
    last_name: str
    template_name: str
    def __init__(self, first_name: _Optional[str] = ..., last_name: _Optional[str] = ..., template_name: _Optional[str] = ...) -> None: ...

class SendMessageResponse(_message.Message):
    __slots__ = ["message"]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...
