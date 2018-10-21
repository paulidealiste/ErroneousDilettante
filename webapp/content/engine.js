 //Elements
 let lbtn = u('#wildbutton');
 let rbtn = u('#rantbutton');
 let eron = u('#DilettanteName');
 let victim = u('#chosenvictim');
 //Interactions
 let chooseNameli = (e) => {
   let elm = u(e.target);
   u('div.reprobationRow').removeClass('uk-background-secondary uk-light');
   elm.parent().addClass('uk-background-secondary uk-light');
   victim.text(elm.text());
 };
 //Functions
 let renderResponse = (person) => {
   const nameli = '<div class="reprobationRow uk-padding-small" style="border-bottom: 1px solid grey"><span style="cursor: pointer;" class="uk-width-1-1">' + person.Name + ' ' + person.Surname + '</span></div>';
   eron.append(nameli)
     .find('span')
     .off('click')
     .handle('click', chooseNameli);
 };
 //Handlers
 lbtn.handle('click', async e => {
   const res = await fetch('/erroneus', {
     method: 'POST'
   });
   const data = await res.json();
   renderResponse(data);
 });