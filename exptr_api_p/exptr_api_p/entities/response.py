from pydantic import BaseModel

class Response(BaseModel):
  status: str
  error: str

  def OK(self):
    return Response(status="OK", error="")
  
  def ERROR(self, error):
    return Response(status="ERROR", error=error)