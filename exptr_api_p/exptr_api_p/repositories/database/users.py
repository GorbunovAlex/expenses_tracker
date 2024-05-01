from sqlalchemy import DateTime, Column, Integer, String

from exptr_api_p.core.db import Base

class User(Base):
  __tablename__ = "users"

  id = Column(Integer, primary_key=True, autoincrement=True, unique=True)
  authn_id = Column(String)
  email = Column(String, unique=True, nullable=False)
  password = Column(String)
  created_at = Column(DateTime)
  updated_at = Column(DateTime)