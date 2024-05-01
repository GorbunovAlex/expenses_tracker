from exptr_api_p.repositories import categories as CatRepository
from exptr_api_p.entities import categories

class CategoriesService:
  def __init__(self, repository: CatRepository):
    self.repository = repository

  def get_categories(self, user_id: int):
    return self.repository.get_categories(user_id)
  
  def create_category(self, category: categories.CategoryRequest):
    return self.repository.create_category(category)
  
  def update_category(self, category: categories.Category):
    return self.repository.update_category(category)
  
  def delete_category(self, category_id: int):
    return self.repository.delete_category(category_id)