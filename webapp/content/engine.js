class ErroneusEngine {
  constructor() {
  }
  request = async (artists) => {
    const res = await fetch('/erroneus', {
      method: 'POST'
    });
    const data = await res.json();
    artists.push(data);
  }
}