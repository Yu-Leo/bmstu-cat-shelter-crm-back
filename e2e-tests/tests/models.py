from dataclasses import dataclass

@dataclass
class Room:
    number: str
    status: str

@dataclass
class CreateRoomInput:
    number: str
    status: str

    
@dataclass
class UpdateRoomInput:
    number: str
    status: str

