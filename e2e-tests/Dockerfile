FROM python:3.10.4

ENV PYTHONUNBUFFERED 1

WORKDIR /usr/src/app

# Requirements
RUN pip install --upgrade pip
COPY ./requirements.txt ./requirements.txt
RUN pip install -r requirements.txt

COPY . /usr/src/app/
