from exptr_api_p.entities import users

class UsersService:
  def __init__(self, repository):
    self.repository = repository

  def get_user_by_email(self, email: str):
    return self.repository.get_user_by_email(email)
  
  def create_user(self, user: users.User):
    return self.repository.create_user(user)
  
  def update_user(self, user: users.User):
    return self.repository.update_user(user)
  
  def set_user_session(self, user_id: int, token: str):
    return self.repository.set_user_session(user_id, token)
  
  def updated_user_session(self, user_id: int, token: str):
    return self.repository.updated_user_session(user_id, token)
  
  def get_user_session_by_token(self, token: str):
    return self.repository.get_user_session_by_token(token)
  
  def get_user_session(self, user_id: int):
    return self.repository.get_user_session(user_id)
  
  def delete_user_session(self, user_id: int):
    return self.repository.delete_user_session(user_id)
  
  def delete_outdated_sessions(self):
    return self.repository.delete_outdated_sessions()