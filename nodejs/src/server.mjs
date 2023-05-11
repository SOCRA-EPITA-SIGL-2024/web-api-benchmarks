import express from "express";
import { distance } from "./distance.mjs";
import { env, exit } from "process";

const GARDENS_JSON_FILE_PATH = env["GARDENS_JSON_FILE_PATH"] || exit(1);
import { createRequire } from "module";
const require = createRequire(import.meta.url);
const GARDENS = require(GARDENS_JSON_FILE_PATH);

const app = express();
const port = 3000;

app.use(express.json());

app.get("/", function (request, response) {
  response.send({ Hello: "Socarotte!" });
});

/**
 * Find all products that are 10km for the position of the user
 */
app.post("/v1/products", function (request, response) {
  try {
    const radius = +request.query.radius;
    const userPosition = request.body.position;
    const productsAroundUsers = GARDENS.map(function (product) {
      const distanceToProduct = distance(userPosition, product.position);
      return {
        ...product,
        distance: distanceToProduct,
      };
    }).filter(function ({ distance }) {
      return distance < radius;
    });
    console.log("NODEJS: products sent to user with position", userPosition);
    response.send(productsAroundUsers);
  } catch (e) {
    console.log(`ERROR with /v1/products: ${e}`);
    response.status = 500;
    response.send({ details: e });
  }
});

app.listen(port, function () {
  console.log(`Server started on localhost:${port}`);
});
