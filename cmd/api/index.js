// Création de l'objet XMLHttpRequest
const xhr = new XMLHttpRequest();

// Configuration de la requête
xhr.open('GET', 'http://localhost:8080/api/aeroports');

// Envoi de la requête
xhr.send();

const fieldSet = document.getElementById("fieldSet")

// Gestion de la réponse
xhr.onload = function() {
  if (xhr.status === 200) {
    // Récupération des données JSON
    const data = JSON.parse(xhr.responseText);

    // Parcours du tableau
    for (const item of data) {
        // Création de l'élément div
        const div = document.createElement('div');
    
        // Création de l'élément input
        const input = document.createElement('input');
        input.setAttribute('type', 'radio');
        input.setAttribute('id', item.name);
        input.setAttribute('name', 'aireport');
        input.setAttribute('value', item.name);
    
        // Création de l'élément label
        const label = document.createElement('label');
        label.setAttribute('for', item.name);
        label.textContent = item.name;
    
        // Ajout des éléments à la page
        div.appendChild(input);
        div.appendChild(label);
        fieldSet.appendChild(div);
    }

    // Récupération des éléments input de type radio
    const radioButtons = document.querySelectorAll('input[type=radio]');
    const calendar = document.getElementById('date')
    var value;

    // Ajout d'un écouteur d'événement change aux éléments input
    radioButtons.forEach(button => button.addEventListener('change', function() {

        value = this.value;
        calendar.removeAttribute('hidden');
        
        document.getElementById('txtAeroportDate').innerHTML = `Moyenne des mesures pour ${value} le : `;

        const today = new Date();
        const formattedDate = today.toISOString().slice(0, 10);
        calendar.value = formattedDate;
        calendar.dispatchEvent(new Event('change'));
        
    }));

    calendar.addEventListener('change', function() {
        // Récupération de la valeur de la date sélectionnée
        const date = this.value;
        
        const xhr2 = new XMLHttpRequest();
        xhr2.open('GET', `http://localhost:8080/api/allMesure/${value}/${date}`);
        xhr2.send();
        xhr2.onload = function() {
            if (xhr2.status === 200) {
                const data = JSON.parse(xhr2.responseText);
                document.getElementById('nomAero').innerHTML = `Nom de l'aéroport : ${data.Name}`;
                document.getElementById('temp').innerHTML = `Température : ${data.Temperature}`;
                document.getElementById('pres').innerHTML = `Pression : ${data.Pressure}`;
                document.getElementById('wind').innerHTML = `Vitesse du vent : ${data.Wind_speed}`;
                console.log(data);
            } else {
                console.error('Une erreur est survenue');
            }
        };
    });

    const madButton = document.getElementById('NTE');
    madButton.checked = true;
    madButton.dispatchEvent(new Event('change'));

  } else {
    console.error('Une erreur est survenue');
  }
}


