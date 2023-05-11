from typing import Tuple, Union

from fastapi import FastAPI

from pydantic import BaseModel

from distance import distance

from json import load
from os import environ


class Position(BaseModel):
    lat: float
    lng: float


class PositionInput(BaseModel):
    position: Position


app = FastAPI()


def getRequiredEnvVar(env_key: str):
    env_value = environ.get(env_key)
    if env_value is None:
        raise Exception(f"missing required env var {env_key}")
    return env_value


def load_gardens():
    GARDENS_JSON_FILE_PATH = getRequiredEnvVar("GARDENS_JSON_FILE_PATH")
    gardens = []
    with open(GARDENS_JSON_FILE_PATH, "r") as garden_json_file:
        gardens = load(garden_json_file)
    return gardens


# loads all gardens **in memory** when web-service is starting
GARDENS = list(load_gardens())


@app.get("/")
def read_root():
    return {"Hello": "Socarotte"}


@app.post("/v1/products")
async def read_products(position_input: PositionInput, radius: Union[str, None] = None):
    user_position = dict(position_input.position)
    radius_int = int(radius)
    gardens_near_user = []

    def compute_distance(garden):
        garden_distance_from_user = distance(
            start_pos=user_position, dest_pos=garden["position"]
        )
        return {**garden, "distance": garden_distance_from_user}

    def is_inside_radius(garden_with_distance):
        user_distance_from_garden = garden_with_distance["distance"]
        return user_distance_from_garden < radius_int

    gardens_near_user = filter(is_inside_radius, map(compute_distance, GARDENS))
    print(f"PYTHON: products sent to user with position {user_position}")
    return list(gardens_near_user)
