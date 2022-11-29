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


## Aéroports :

- NTE

## BDD redis :

- `/AERO/DATEJOUR/MESURE => valeur;datejour+heure`
- `/AERO/DATEJOURHEURE => datejour+heure (de la premiere mesure)`

## MESURE:
- TEMP
- PRESSURE
- WIND

## Requêtes API REST :

> Mesure d'un type entre deux date(+heure) pour un aeroport
>
> `GET /localhost:port/api/NTE/TEMP/18-11-2022-18:18:10/18-11-2022-18:18:20`
> ```json
> {
>     "NTE": {
>         "TEMP": [
>             {
>                 "TEMP": 26,
>                 "DATE": "18-11-2022-18:18:10"
>             },
>             {
>                 "TEMP": 26,
>                 "DATE": "18-11-2022-18:18:20"
>             }
>         ]
>     }
> }
> ```

> Moyenne mesure pour les MESURE pour un aeroport
>
> `GET /localhost:port/api/NTE/TEMP/`
> ```json
> {
>     "NTE": {
>         "TEMP": 19
>     }
> }
> ```

https://www.soberkoder.com/swagger-go-api-swaggo/
