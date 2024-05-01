from fastapi import APIRouter, HTTPException

from exptr_api_p.entities import categories
from exptr_api_p.services import categories as categories_service

router = APIRouter()

@router.get("/", response_model=categories.GetCategoriesResponse)
def get_categories():
    categories = categories_service.CategoriesService.get_categories(categories_service.CategoriesService, 1)
    if not categories:
        raise HTTPException(status_code=404, detail="Categories not found")
    return categories

@router.post("/", response_model=categories.CategoryResponse)
def create_category(category: categories.CategoryRequest):
    return categories.CategoriesService.create_category(category)

@router.put("/", response_model=categories.CategoryResponse)
def update_category(category: categories.Category):
    return categories.CategoriesService.update_category(category)

@router.delete("/", response_model=categories.CategoryResponse)
def delete_category(category_id: int):
    return categories.CategoriesService.delete_category(category_id)