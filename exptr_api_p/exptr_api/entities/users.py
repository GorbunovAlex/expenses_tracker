import datetime
from pydantic import BaseModel

from exptr_api.entities import response

class User(BaseModel):
  id: int
  authn_id: str
  email: str
  created_at: datetime.datetime
  updated_at: datetime.datetime

class SignUpRequest(BaseModel):
  email: str
  password: str
  
class LoginRequest(BaseModel):
  email: str
  password: str

class LoginResponse(BaseModel):
  token: str
  response: response.Response

class UserSession(BaseModel):
  id: str
  user_id: str
  created_at: datetime.datetime
  token: str