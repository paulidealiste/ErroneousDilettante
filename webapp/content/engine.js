class ErroneusEngine {
  constructor() {
  }
  request = async (artists) => {
    const req = await fetch('/erroneus', {
      method: 'POST'
    });
    const data = await req.json();
    artists.push(data);
  }
  pegged = async () => {
    const req = await fetch('/pegged')
    const data = await req.json();
    const free = data.map(dpeg => dpeg.Content);
    return free;
  }
  peg = async (pegged) => {
    const rdat = JSON.stringify(pegged);
    const req = await fetch('/pstore', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: rdat
    });
    console.log(req);
  }
  clearpegged = async () => {
    const req = await fetch('/pclear')
    console.log(req);
  }
}