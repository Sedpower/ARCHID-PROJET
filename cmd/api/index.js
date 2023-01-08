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
                document.getElementById('temp').innerHTML = `Température : ${data.Temperature} °C`;
                document.getElementById('pres').innerHTML = `Pression : ${data.Pressure} hPa`;
                document.getElementById('wind').innerHTML = `Vitesse du vent : ${data.Wind_speed} m/h`;
            } else {
                console.error('Une erreur est survenue');
            }
        };
    });

    let xhr3 = new XMLHttpRequest();
    xhr3.open('GET', 'http://localhost:8080/api/mesures/NTE/2023-01-08-17/2023-01-08-23');
    xhr3.send();
    xhr3.onload = function() {
      if (xhr3.status === 200) {
        let data3 = JSON.parse(xhr3.responseText);
        displayCharts(data3)
      }
    }

    const madButton = document.getElementById('NTE');
    madButton.checked = true;
    madButton.dispatchEvent(new Event('change'));

  } else {
    console.error('Une erreur est survenue');
  }
}

function displayCharts(data) {

  // Récupérer les données de pression, de température et de vitesse du vent
  let pressureData = data.pressions[0].HeureMesure.sort((a,b) => a.heure.localeCompare(b.heure)).map(h => h.Mesure.Value)
  let temperatureData = data.temperatures[0].HeureMesure.sort((a,b) => a.heure.localeCompare(b.heure)).map(h => h.Mesure.Value)
  let windSpeedData = data.vitesseVents[0].HeureMesure.sort((a,b) => a.heure.localeCompare(b.heure)).map(h => h.Mesure.Value)

  // Récupérer les heures de mesure
  let hoursPres = data.pressions[0].HeureMesure.map(h => h.heure);
  let hoursTemp = data.temperatures[0].HeureMesure.map(h => h.heure);
  let hoursWind = data.vitesseVents[0].HeureMesure.map(h => h.heure);

  // Créer les options de graphique pour chaque type de donnée
  const pressureOptions = {
    type: 'line',
    data: {
      labels: hoursPres,
      datasets: [{
        label: 'Pression',
        data: pressureData,
        backgroundColor: 'rgba(255, 99, 132, 0.2)',
        borderColor: 'rgba(255, 99, 132, 1)',
        borderWidth: 1,
        pointRadius: 1
      }]
    },
    options: {
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true
          }
        }]
      }
    }
  };

  const temperatureOptions = {
    type: 'line',
    data: {
      labels: hoursTemp,
      datasets: [{
        label: 'Température',
        data: temperatureData,
        backgroundColor: 'rgba(255, 206, 86, 0.2)',
        borderColor: 'rgba(255, 206, 86, 1)',
        borderWidth: 1,
        pointRadius: 1
      }]
    },
    options: {
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true
          }
        }]
      }
    }
  };

  const windSpeedOptions = {
    type: 'line',
    data: {
      labels: hoursWind,
      datasets: [{
        label: 'Vitesse du vent',
        data: windSpeedData,
        backgroundColor: 'rgba(75, 192, 192, 0.2)',
        borderColor: 'rgba(75, 192, 192, 1)',
        borderWidth: 1,
        pointRadius: 1
      }]
    },
    options: {
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true
          }
        }]
      }
    }
  };
// Créer les graphiques en utilisant Chart.js
const pressureChart = new Chart(document.getElementById('pressure-chart'), pressureOptions);
const temperatureChart = new Chart(document.getElementById('temperature-chart'), temperatureOptions);
const windSpeedChart = new Chart(document.getElementById('wind-speed-chart'), windSpeedOptions);

// Mettre à jour les graphiques toutes les 10 secondes
let refreshIntervalId = setInterval(async function() {
  // Récupérer les nouvelles données de mesure météorologique de l'API
  const response = await fetch(`http://localhost:8080/api/mesures/NTE/2023-01-08-15/2023-01-08-21`);
  const data = await response.json();

  // Récupérer les données de pression, de température et de vitesse du vent
  let pressureData = data.pressions[0].HeureMesure.sort((a,b) => a.heure.localeCompare(b.heure)).map(h => h.Mesure.Value)
  let temperatureData = data.temperatures[0].HeureMesure.sort((a,b) => a.heure.localeCompare(b.heure)).map(h => h.Mesure.Value)
  let windSpeedData = data.vitesseVents[0].HeureMesure.sort((a,b) => a.heure.localeCompare(b.heure)).map(h => h.Mesure.Value)

  // Récupérer les heures de mesure
  let hoursPrest = data.pressions[0].HeureMesure.map(h => h.heure);
  let hoursTempt = data.temperatures[0].HeureMesure.map(h => h.heure);
  let hoursWindt = data.vitesseVents[0].HeureMesure.map(h => h.heure);

  pressureChart.data.datasets[0].data = pressureData
  temperatureChart.data.datasets[0].data = temperatureData
  windSpeedChart.data.datasets[0].data = windSpeedData

  pressureChart.data.labels = hoursPrest;
  temperatureChart.data.labels = hoursTempt;
  windSpeedChart.data.labels = hoursWindt;

  pressureChart.update();
  temperatureChart.update();
  windSpeedChart.update();
}, 10000);

function stopRefresh() {
  clearInterval(refreshIntervalId);
}

}

