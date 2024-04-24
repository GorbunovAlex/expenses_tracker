from sqlalchemy.orm import Session

from exptr_api.entities import categories
from exptr_api.repositories.database import categories as db_categories

class CategoriesRepository:
  def __init__(self, session: Session):
    self.session = session

  def create_category(self, category: categories.CategoryRequest):
    db_category = db_categories.Category(
      user_id=category.user_id,
      name=category.name,
      type=category.type,
      created_at=category.created_at,
      updated_at=category.updated_at,
      color=category.color,
      icon=category.icon
    )
    self.session.add(db_category)
    self.session.commit()
    self.session.refresh(db_category)
    return db_category

  def update_category(self, category: categories.Category):
    db_category = self.session.query(db_categories.Category).filter(db_categories.Category.id == category.id).first()
    db_category.user_id = category.user_id
    db_category.name = category.name
    db_category.type = category.type
    db_category.updated_at = category.updated_at
    db_category.color = category.color
    db_category.icon = category.icon
    self.session.commit()
    self.session.refresh(db_category)
    return db_category

  def get_categories(self, user_id: int):
    return self.session.query(db_categories.Category).filter(db_categories.Category.user_id == user_id)

  def delete_category(self, category_id: int):
    self.session.query(db_categories.Category).filter(db_categories.Category.id == category_id).delete()
    self.session.commit()