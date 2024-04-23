from sqlalchemy import DateTime, Column, Integer, String

from .db import Base

class User:
  __tablename__ = "users"

  id = Column(Integer, primary_key=True, autoincrement=True, unique=True)
  authn_id = Column(String)
  email = Column(String, unique=True, nullable=False)
  password = Column(String)
  created_at = Column(DateTime)
  updated_at = Column(DateTime)

  def __init__(self, ID, AuthnID, Email, Password, CreatedAt, UpdatedAt):
    self.ID = ID
    self.AuthnID = AuthnID
    self.Email = Email
    self.Password = Password
    self.CreatedAt = CreatedAt
    self.UpdatedAt = UpdatedAt