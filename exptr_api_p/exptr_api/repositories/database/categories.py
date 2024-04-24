from sqlalchemy import DateTime, Column, ForeignKey, Integer, String

from exptr_api.core.db import Base

class Category(Base):
    __tablename__ = "categories"

    id = Column(Integer, primary_key=True, autoincrement=True, unique=True)
    user_id = Column(Integer, ForeignKey("users.id"), nullable=True, unique=False)
    name = Column(String, nullable=False, unique=False)
    type = Column(String, nullable=False, unique=False)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)
    color = Column(String)
    icon = Column(String)