import datetime

from sqlalchemy import DateTime, Column, ForeignKey, Integer, String
from sqlalchemy.orm import relationship

from .db import Base

class Category(Base):
    __tablename__ = "categories"

    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, ForeignKey("users.id"))
    name = Column(String)
    type = Column(String)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)
    color = Column(String)
    icon = Column(String)

    def __init__(self, id: str, user_id: str, name: str, type: str, created_at: datetime.datetime, updated_at: datetime.datetime, color: str, icon: str):
        self.id = id
        self.user_id = user_id
        self.name = name
        self.type = type
        self.created_at = created_at
        self.updated_at = updated_at
        self.color = color
        self.icon = icon

class CategoryRequest:
    def __init__(self, user_id: str, name: str, type: str, created_at: datetime.datetime, updated_at: datetime.datetime, color: str, icon: str):
        self.user_id = user_id
        self.name = name
        self.type = type
        self.created_at = created_at
        self.updated_at = updated_at
        self.color = color
        self.icon = icon

class CategoryResponse:
    def __init__(self, response: response.Response):
        self.response = response

class GetCategoriesResponse:
    def __init__(self, response: response.Response, categories: List[Category]):
        self.response = response
        self.categories = categories