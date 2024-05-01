from sqlalchemy import DateTime, Column, ForeignKey, Integer, String

from sqlalchemy.orm import relationship

from exptr_api_p.core.db import Base

class Operation(Base):
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