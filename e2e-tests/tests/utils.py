import requests
from dataclasses import asdict

import models


def create_room(base_url, create_room_input: models.CreateRoomInput) -> models.Room:
    url = base_url + 'rooms/'
    requests.post(url, json=asdict(create_room_input))

    return models.Room(number=create_room_input.number,
                       status=create_room_input.status)
