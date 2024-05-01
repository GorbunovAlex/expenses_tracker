import datetime

from pydantic import BaseModel

from exptr_api_p.entities import response


class Category(BaseModel):
    id: int
    user_id: int
    name: str
    type: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    color: str
    icon: str
    

class CategoryRequest(Category):
    user_id: int
    name: str
    type: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    color: str
    icon: str

class CategoryResponse(BaseModel):
    response: response.Response
    
class GetCategoriesResponse(BaseModel):
    response: response.Response
    categories: list[Category]