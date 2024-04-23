class Operation:
    def __init__(self, ID, UserID, CategoryID, Amount, Currency, Name, Comment, Type, CreatedAt, UpdatedAt):
        self.ID = ID
        self.UserID = UserID
        self.CategoryID = CategoryID
        self.Amount = Amount
        self.Currency = Currency
        self.Name = Name
        self.Comment = Comment
        self.Type = Type
        self.CreatedAt = CreatedAt
        self.UpdatedAt = UpdatedAt

class OperationRequest:
    def __init__(self, UserID, CategoryID, Amount, Currency, Name, Comment, Type, CreatedAt, UpdatedAt):
        self.UserID = UserID
        self.CategoryID = CategoryID
        self.Amount = Amount
        self.Currency = Currency
        self.Name = Name
        self.Comment = Comment
        self.Type = Type
        self.CreatedAt = CreatedAt
        self.UpdatedAt = UpdatedAt

class CreateOperationResponse:
    def __init__(self, Response):
        self.Response = Response

class GetOperationsByUserIDResponse:
    def __init__(self, Response, Operations):
        self.Response = Response
        self.Operations = Operations

class UpdateOperationResponse:
    def __init__(self, Response):
        self.Response = Response