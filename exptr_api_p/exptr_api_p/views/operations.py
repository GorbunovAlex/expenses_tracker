from fastapi import APIRouter, HTTPException

from exptr_api_p.entities import operations
from exptr_api_p.entities import response

router = APIRouter()

@router.get("/", response_model=operations.GetOperationsByUserIDResponse)
def get_operations_by_user_id(user_id: int):
    operations = operations.OperationsService.get_operations_by_user_id(user_id)
    if not operations:
        raise HTTPException(status_code=404, detail="Operations not found")
    return operations

@router.get("?id={operation_id}", response_model=response.Response)
def get_operation_by_id(operation_id: int):
    operation = operations.OperationsService.get_operation_by_id(operation_id)
    if not operation:
        raise HTTPException(status_code=404, detail="Operation not found")
    return operation

@router.post("/", response_model=operations.CreateOperationResponse)
def create_operation(operation: operations.OperationRequest):
    return operations.OperationsService.create_operation(operation)

@router.put("/", response_model=operations.UpdateOperationResponse)
def update_operation(operation: operations.Operation):
    return operations.OperationsService.update_operation(operation)

@router.delete("?id={operation_id}")
def delete_operation(operation_id: int):
    return operations.OperationsService.delete_operation(operation_id) 