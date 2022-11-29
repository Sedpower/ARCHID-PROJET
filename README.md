# ARCHID-PROJET

```
Architecture attendue

project-root/
    ├── cmd/
    │   ├── pub
    │   │   ├── pub.go
    │   │   └── pub.exe
    │   └── sub
    │       ├── sub.go
    │       └── sub.exe
    ├── internal/
    └── ...
```


Aeroport : 

NTE

BDD redis :

/AERO/DATEJOUR/MESURE => valeur;datejour+heure
/AERO/DATEJOURHEURE => datejour+heure (de la premiere mesure)

MESURE:
TEMP
PRESSURE
WIND

Requête api rest :

mesure d'un type entre deux date(+heure) pour un aeroport

GET /localhost:port/api/NTE/TEMP/18-11-2022-18:18:10/18-11-2022-18:18:20
```
{
    "NTE": {
        "TEMP": [
            {
                "TEMP": 26,
                "DATE": "18-11-2022-18:18:10"
            },
            {
                "TEMP": 26,
                "DATE": "18-11-2022-18:18:20"
            }
        ]
    }
}
```

moyenne mesure pour les MESURE pour un aeroport

GET /localhost:port/api/NTE/TEMP/
```
{
    "NTE": {
        "TEMP": 19
    }
}
```
