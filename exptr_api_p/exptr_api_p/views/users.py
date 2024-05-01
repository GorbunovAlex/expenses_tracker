from fastapi import APIRouter, HTTPException

from exptr_api_p.entities import users

router = APIRouter()

@router.get("?email={email}", response_model=users.User)
def get_user_by_email(email: str):
    user = users.UsersService.get_user_by_email(email)
    if not user:
        raise HTTPException(status_code=404, detail="User not found")
    return user

@router.put("/", response_model=users.User)
def update_user(user: users.User):
    return users.UsersService.update_user(user)

@router.post('/login', response_model=users.LoginRequest)
def login(user: users.LoginRequest):
    return user

@router.post('/signup', response_model=users.SignUpRequest)
def signup(user: users.SignUpRequest):
    return user