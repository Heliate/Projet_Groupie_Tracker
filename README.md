# PokeBooster

## Présentation du projet

PokeBooster est une application web éducative développée dans le cadre d'un projet étudiant, permettant de simuler l'ouverture de boosters de cartes Pokémon virtuels. L'application offre également un catalogue complet des cartes Pokémon organisé par set, avec des fonctionnalités de recherche, de filtrage et un système de favoris.

### Fonctionnalités principales

- Simulation d'ouverture de boosters Pokémon
- Catalogue complet des cartes Pokémon par set
- Recherche de cartes par nom
- Filtrage des cartes par type, rareté, faiblesse et HP
- Système de favoris pour sauvegarder vos cartes préférées
- Interface responsive et conviviale

### Objectifs pédagogiques

Ce projet a été conçu pour mettre en pratique les compétences suivantes :
- Développement web backend avec Go
- Utilisation des templates HTML
- Interaction avec une API externe
- Gestion des données utilisateurs
- Structure MVC (Model-View-Controller)

## Installation et lancement du projet

### Prérequis

- Go 1.18 ou supérieur
- Connexion internet (pour accéder à l'API TCGdex)

### Installation

1. Clonez le dépôt GitHub :
```bash
git clone https://github.com/Heliate/Projet_Groupie_Tracker.git
cd .\TP_API_Pokemon\
```

2. Assurez-vous que la structure du projet est correcte avec les dossiers suivants :
   - `assets/` - Contient les fichiers CSS, JavaScript et images
   - `controllers/` - Contient les contrôleurs de l'application
   - `routes/` - Contient les définitions des routes
   - `Services/` - Contient les services d'accès aux API
   - `templates/` - Contient les templates HTML
   - `utils/` - Contient les utilitaires

3. Créez le dossier `data` à la racine du projet pour stocker les données utilisateur :
```bash
mkdir data
```

### Lancement

Pour lancer l'application, exécutez :
```bash
go run main.go
```

L'application sera accessible à l'adresse : `http://localhost:8000`

## Structure de l'API

### Routes implémentées

| Route | Méthode | Description |
|-------|---------|-------------|
| `/` | GET | Page d'accueil avec le carousel de boosters |
| `/cards` | GET | Affiche le catalogue de cartes (avec pagination) |
| `/cards?set=[set_id]` | GET | Affiche les cartes d'un set spécifique |
| `/cards?search=[term]` | GET | Recherche de cartes par nom |
| `/cards?page=[number]` | GET | Navigation entre les pages de cartes |
| `/favoris` | GET | Affiche les cartes favorites de l'utilisateur |
| `/favoris/add` | POST | Ajoute une carte aux favoris |
| `/favoris/remove` | POST | Retire une carte des favoris |
| `/cgu` | GET | Affiche les conditions générales d'utilisation |

### Paramètres de filtrage

Pour le filtrage des cartes dans la route `/cards`, les paramètres suivants sont disponibles :

| Paramètre | Description | Exemple |
|-----------|-------------|---------|
| `type_filter` | Filtre par type de Pokémon | `?type_filter=Feu&type_filter=Eau` |
| `rarity_filter` | Filtre par rareté | `?rarity_filter=Rare&rarity_filter=Commune` |
| `weakness_filter` | Filtre par faiblesse | `?weakness_filter=Électrique` |
| `sort_filter` | Tri par HP (0: croissant, 1: décroissant) | `?sort_filter=1` |

## API externe utilisée

L'application utilise l'API TCGdex pour récupérer les informations sur les cartes Pokémon.

### Endpoints exploités

| Endpoint | Description |
|----------|-------------|
| `https://api.tcgdex.net/v2/fr/cards/[set_id]-[number]` | Récupère les détails d'une carte spécifique d'un set |
| `https://api.tcgdex.net/v2/fr/cards/[card_id]` | Récupère les détails d'une carte par son ID |

### Structure des données

Les cartes Pokémon sont structurées avec les champs suivants :
- `ID` : Identifiant unique de la carte
- `Image` : URL de l'image de la carte
- `Name` : Nom du Pokémon
- `Rarity` : Rareté de la carte
- `Hp` : Points de vie du Pokémon
- `Types` : Types du Pokémon (Feu, Eau, etc.)
- `Weaknesses` : Faiblesses du Pokémon

## Gestion des données

Les favoris des utilisateurs sont stockés dans un fichier JSON situé dans le dossier `data/`. Chaque utilisateur est identifié par un ID unique et peut stocker un nombre illimité de cartes dans ses favoris.

## Licence

Ce projet est développé à des fins éducatives et n'a aucun lien officiel avec The Pokémon Company. Toutes les images et noms de Pokémon sont la propriété de leurs détenteurs respectifs.

© 2025 PokeBooster - Projet étudiant
