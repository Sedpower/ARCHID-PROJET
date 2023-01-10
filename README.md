# ARCHID-PROJET

## Comment éxécuter le projet

Executer le fichier `build_demo.bat`. Cela va générer un dossier `build/` à la racine du projet contenant tout les executables `.exe` et fichiers de configurations nécéssaire au fonctionnement de l'application.

Une fois cela fait, il suffit de lancer dans un terminal à part la commande `mosquitto -v` pour mettre en route le broker MQTT. <br>
Il faut également lancer le service `redis-server` en administrateur pour mettre en route la base de donnees NoSql Redis.

Une fois tout cela fait, il suffit de lancer les différents executables `.exe` dans le dossier `build/` pour lancer l'application.

_L'application est totalement personnalisable, en effet vous pouvez si vous le souhaitez changer l'URL du broker MQTT ainsi que la base redis dans le fichier `config.yaml`_

---

## Aéroports :
- NTE
- MAD
- CDG

_Ce sont les valeurs par défaut.
Il est bien sûr tout à fait possible d'ajouter des aéroports à sa guise, il suffit de modifier le config.yaml_

## BDD redis :
- `/<iata>/<natureDonnée>/<année>/<mois>/<jour>/<heure>/<minute>/<seconde> => idCapteur valeur`

## NATURE DONNEE:
- Temperature
- Atmospheric pressure
- Wind speed

## Requêtes API REST :
> Mesure d'un type entre deux date(+heure) pour un aeroport
>
> `GET http://localhost:8080/api/mesures/NTE/2023-01-08-00/2023-01-09-00`
> ```json
> {
>    "temperatures": [
>        {
>            "jour": "JJ/MM/AAAA",
>            "HeureMesure": [
>                {
>                    "heure": "HH:MM:SS",
>                    "Mesure": {
>                        "idCapteur": String,
>                        "Value": float
>                    }
>                },
>                {...}
>            ]
>        },
>        {...}
>    ],
>    "pressions": [
>        {
>            "jour": "JJ/MM/AAAA",
>            "HeureMesure": [
>                {
>                    "heure": "HH:MM:SS",
>                    "Mesure": {
>                        "idCapteur": String,
>                        "Value": float
>                    }
>                },
>                {...}
>            ]
>        },
>        {...}
>    ],
>    "vitesseVents": [
>        {
>            "jour": "JJ/MM/AAAA",
>            "HeureMesure": [
>                {
>                    "heure": "HH:MM:SS",
>                    "Mesure": {
>                        "idCapteur": String,
>                        "Value": float
>                    }
>                },
>                {...}
>            ]
>        },
>        {...}
>    ]
>}
> ```

> Moyenne mesure pour les MESURE pour un aeroport
>
> `GET http://localhost:8080/api/allMesure/NTE/2023-01-10`
> ```json
> {
>     "Name": String, // Correspond à l'IATA
>     "Temperature": String, // Est un float transformé en String et correspond à la moyenne des températures
>     "Pressure": String, // Est un float transformé en String et correspond à la moyenne des pressions
>     "Wind_speed": String // Est un float transformé en String et correspond à la moyenne des vitesses de vent
> }
> ```