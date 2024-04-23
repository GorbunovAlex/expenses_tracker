from sqlalchemy import DateTime, Column, ForeignKey, Integer, String

from .db import Base

class Operation:
    __tablename__ = "operations"

    id = Column(Integer, primary_key=True, autoincrement=True, unique=True)
    user_id = Column(Integer, ForeignKey("users.id"), nullable=False, unique=False)
    category_id = Column(Integer, ForeignKey("categories.id"), nullable=False, unique=False)
    amount = Column(Integer, nullable=False, unique=False)
    currency = Column(String, nullable=False, unique=False)
    name = Column(String, nullable=False, unique=False)
    comment = Column(String, nullable=True, unique=False)
    type = Column(String, nullable=False, unique=False)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)

    def __init__(self, ID, UserID, CategoryID, Amount, Currency, Name, Comment, Type, CreatedAt, UpdatedAt):
        self.ID = ID
        self.UserID = UserID
        self.CategoryID = CategoryID
        self.Amount = Amount
        self.Currency = Currency
        self.Name = Name
        self.Comment = Comment
        self.Type = Type
        self.CreatedAt = CreatedAt
        self.UpdatedAt = UpdatedAt

class OperationRequest:
    def __init__(self, UserID, CategoryID, Amount, Currency, Name, Comment, Type, CreatedAt, UpdatedAt):
        self.UserID = UserID
        self.CategoryID = CategoryID
        self.Amount = Amount
        self.Currency = Currency
        self.Name = Name
        self.Comment = Comment
        self.Type = Type
        self.CreatedAt = CreatedAt
        self.UpdatedAt = UpdatedAt

class CreateOperationResponse:
    def __init__(self, Response):
        self.Response = Response

class GetOperationsByUserIDResponse:
    def __init__(self, Response, Operations):
        self.Response = Response
        self.Operations = Operations

class UpdateOperationResponse:
    def __init__(self, Response):
        self.Response = Response