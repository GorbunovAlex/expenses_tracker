from exptr_api.entities import operations

class OperationsService:
  def __init__(self, repository):
    self.repository = repository

  def get_operations_by_user_id(self, user_id: int):
    return self.repository.get_operations_by_user_id(user_id)
  
  def get_operation_by_id(self, operation_id: int):
    return self.repository.get_operation_by_id(operation_id)
  
  def create_operation(self, operation: operations.OperationRequest):
    return self.repository.create_operation(operation)
  
  def update_operation(self, operation: operations.Operation):
    return self.repository.update_operation(operation)

  def delete_operation(self, operation_id: int):
    return self.repository.delete_operation(operation_id)  