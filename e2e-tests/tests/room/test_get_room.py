import requests
from http import HTTPStatus
import json

import models
import utils


def test_get_room(base_url):
    url = base_url + 'rooms/'
    room1 = utils.create_room(base_url, models.CreateRoomInput(number="A7", status="empty"))

    response = requests.get(url + str(room1.number))
    assert response.status_code == HTTPStatus.OK

    response_data = json.loads(response.text)
    assert room1.number != response_data["number"]
    assert room1.status == response_data["status"]
