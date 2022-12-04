const express = require('express')
const Quotes = require("randomquote-api");

const app = express()
const port = 3000

app.get('/health', (req, res) => {
  res.status(200).send("v1.0.0")
})

app.get('/quotes', (req, res) => {
  const randomquote = Quotes.randomQuote();
  res.status(200).send(randomquote)
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})