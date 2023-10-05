const express = require('express');
const app = express();

app.use(express.json()); // Parse JSON request bodies

app.get("/", async (_, res) => {
  console.log("Infra test , sample log for / route");
  res.status(200).json({"msg":"Hello world! Infra Team!"});
});

app.post("/postRoute", async (req, res) => {
  const requestBody = req.body;
  console.log("Infra test ,Received POST request with body:", requestBody);
  res.status(200).json({"msg": "POST request received successfully! Infra Team!"});
});

const port = process.env.PORT || 8080;
app.listen(port, () => console.log(`Application is now ready, listening to: ${port} randomId: ${Math.random()}`));
