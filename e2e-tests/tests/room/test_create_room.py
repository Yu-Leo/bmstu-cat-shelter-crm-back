from dataclasses import asdict
import models
import requests
from http import HTTPStatus


def test_create_room(base_url):
    url = base_url + 'rooms/'
    create_room_input = models.CreateRoomInput(number="A5", status="empty")

    response = requests.post(url, json=asdict(create_room_input))
   
    assert response.status_code == HTTPStatus.CREATED
    assert response.text == f"\"{create_room_input.number}\""

