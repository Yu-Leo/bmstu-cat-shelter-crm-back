import requests
from http import HTTPStatus
import json

import models
import utils


def test_get_rooms_list(base_url):
    url = base_url + 'rooms/'
    room1 = utils.create_room(base_url, models.CreateRoomInput(number="A1", status="empty"))
    room2 = utils.create_room(base_url, models.CreateRoomInput(number="A2", status="occupied"))

    response = requests.get(url)
    assert response.status_code == HTTPStatus.OK

    response_data = json.loads(response.text)
    numbers = [r["number"] for r in response_data]
    assert room1.number in numbers
    assert room2.number in numbers
