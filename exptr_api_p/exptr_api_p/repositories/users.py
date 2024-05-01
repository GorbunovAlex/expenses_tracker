import datetime
from sqlalchemy.orm import Session

from exptr_api_p.entities import users
from exptr_api_p.repositories.database import users as db_users

class UsersRepository:
  def __init__(self, session: Session):
    self.session = session

  def create_user(self, user: users.User):
    db_user = db_users.User(
      email=user.email,
      password=user.password,
      created_at=user.created_at,
      updated_at=user.updated_at
    )
    self.session.add(db_user)
    self.session.commit()
    self.session.refresh(db_user)
    return db_user
  
  def update_user(self, user: users.User):
    db_user = self.session.query(db_users.User).filter(db_users.User.id == user.id).first()
    db_user.email = user.email
    db_user.password = user.password
    db_user.updated_at = user.updated_at
    self.session.commit()
    self.session.refresh(db_user)
    return db_user
  
  def get_user_by_email(self, email: str):
    return self.session.query(db_users.User).filter(db_users.User.email == email).first()
  
  def set_user_session(self, user_id: int, token: str):
    session = self.session.query(db_users.UserSession).filter(db_users.UserSession.user_id == user_id).first()
  
    if session is not None:
      session.updated_at = datetime.datetime.now()
      self.session.commit()
      self.session.refresh(session)
    else:
      db_user_session = db_users.UserSession(
        user_id=user_id,
        created_at=datetime.datetime.now(),
        token=token
      )
      self.session.add(db_user_session)
      self.session.commit()
      self.session.refresh(db_user_session)
  
  def updated_user_session(self, user_id: int, token: str):
    self.session.query(db_users.UserSession).filter(db_users.UserSession.user_id == user_id).update({db_users.UserSession.token: token})
    self.session.commit()
  
  def get_user_session_by_token(self, token: str):
    return self.session.query(db_users.UserSession).filter(db_users.UserSession.token == token).first()
  
  def get_user_session(self, user_id: int):
    return self.session.query(db_users.UserSession).filter(db_users.UserSession.user_id == user_id).first()
  
  def delete_user_session(self, user_id: int):
    self.session.query(db_users.UserSession).filter(db_users.UserSession.user_id == user_id).delete()
    self.session.commit()
  
  def delete_outdated_sessions(self):
    self.session.query(db_users.UserSession).filter(db_users.UserSession.created_at < datetime.datetime.now() - datetime.timedelta(hours=1)).delete()
    self.session.commit()