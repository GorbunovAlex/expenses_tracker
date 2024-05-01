from fastapi import APIRouter

from .categories import router as categories_router
from .operations import router as operations_router
from .users import router as users_router

# EXAMPLE
# api_router.include_router(users.router, prefix="/users", tags=["users"])

api_router = APIRouter()

api_router.include_router(categories_router, prefix="/categories", tags=["categories"])
api_router.include_router(operations_router, prefix="/operations", tags=["operations"])
api_router.include_router(users_router, prefix="/users", tags=["users"])
