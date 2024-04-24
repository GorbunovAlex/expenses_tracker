import datetime
from pydantic import BaseModel

from exptr_api.entities import response

class Operation(BaseModel):
    id: int
    user_id: int
    category_id: int
    amount: int
    currency: str
    name: str
    comment: str
    type: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    

class OperationRequest(Operation):
    user_id: int
    category_id: int
    amount: int
    currency: str
    name: str
    comment: str
    type: str
    updated_at: datetime.datetime

class CreateOperationResponse(BaseModel):
    response: response.Response

class GetOperationsByUserIDResponse:
    def __init__(self, Response, Operations):
        self.Response = Response
        self.Operations = Operations

class UpdateOperationResponse:
    def __init__(self, Response):
        self.Response = Response