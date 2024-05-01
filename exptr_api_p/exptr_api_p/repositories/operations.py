from sqlalchemy.orm import Session

from exptr_api_p.entities import operations
from exptr_api_p.repositories.database import operations as db_operations

class OperationsRepository:
  def __init__(self, session: Session):
    self.session = session

  def create_operation(self, operation: operations.OperationRequest):
    db_operation = db_operations.Operation(
      user_id=operation.user_id,
      category_id=operation.category_id,
      amount=operation.amount,
      currency=operation.currency,
      name=operation.name,
      comment=operation.comment,
      type=operation.type,
      created_at=operation.created_at,
      updated_at=operation.updated_at
    )
    self.session.add(db_operation)
    self.session.commit()
    self.session.refresh(db_operation)
  
  def update_operation(self, operation: operations.Operation):
    db_operation = self.session.query(db_operations.Operation).filter(db_operations.Operation.id == operation.id).first()
    db_operation.user_id = operation.user_id
    db_operation.category_id = operation.category_id
    db_operation.amount = operation.amount
    db_operation.currency = operation.currency
    db_operation.name = operation.name
    db_operation.comment = operation.comment
    db_operation.type = operation.type
    db_operation.updated_at = operation.updated_at
    self.session.commit()
    self.session.refresh(db_operation)
  
  def get_operation_by_id(self, operation_id: int):
    return self.session.query(db_operations.Operation).filter(db_operations.Operation.id == operation_id).first()
  
  def get_operations_by_user_id(self, user_id: int):
    return self.session.query(db_operations.Operation).filter(db_operations.Operation.user_id == user_id)
  
  def delete_operation(self, operation_id: int):
    self.session.query(db_operations.Operation).filter(db_operations.Operation.id == operation_id).delete()
    self.session.commit()