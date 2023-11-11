import pytest

from config import settings


@pytest.fixture(scope='session')
def base_url():
    HOST = settings['SERVER_HOST']
    PORT = settings['SERVER_PORT']

    url = f'http://{HOST}:{PORT}/'
    return url
