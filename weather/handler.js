exports.handler = async function (req, res) {
  const city = req.query.city ?? 'montpellier';
  const url = `https://www.google.com/search?q=${city}+weather`;
  let temperature = 0;

  const result = await fetch(url)
  const data = await result.text()

  let regex = /<div class="BNeawe iBp4i AP7Wnd">(.*)<\/div>/;
  let match = data.match(regex);
  if (!match) {
    return res.status(404).send(JSON.stringify(`Je n'ai pas trouvé la météo de ${city}`));
  }
  const temperature_line = match[1];
  regex = /<div class="BNeawe iBp4i AP7Wnd">(-?\d+)/;
  match = temperature_line.match(regex);
  if (!match) {
    return res.status(404).send(JSON.stringify(`Je n'ai pas trouvé la météo de ${city}`));
  }
  temperature = match[1];

  return res.status(200).send(JSON.stringify(`Il fait ${temperature}°C à ${city}`));
};
