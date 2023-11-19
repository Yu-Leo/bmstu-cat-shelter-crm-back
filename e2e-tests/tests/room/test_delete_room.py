import requests
from http import HTTPStatus
import json

import models
import utils


def test_delete_room(base_url):
    url = base_url + 'rooms/'
    room1 = utils.create_room(base_url, models.CreateRoomInput(number="A6", status="empty"))

    del_response = requests.delete(url + str(room1.number))
    assert del_response.status_code == HTTPStatus.NO_CONTENT

    get_response = requests.get(url)
    assert get_response.status_code == HTTPStatus.OK

    response_data = json.loads(get_response.text)
    assert room1.number not in [r["number"] for r in response_data]
