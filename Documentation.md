# Guide détaillé du projet LetsGoBack

## Introduction

Ce document fournit une explication approfondie de la structure du projet, des packages, et des fonctions principales. Il est destiné à faciliter la prise en main rapide et la compréhension du code.

---

## Structure générale

Le projet est organisé en plusieurs packages, chacun ayant une responsabilité claire :

- **Context/** : Gestion du contexte applicatif, des réponses HTTP, de la sérialisation JSON, du JWT, etc.
- **Helpers/** : Fonctions utilitaires de validation.
- **Middleware/** : Middlewares pour la gestion des requêtes HTTP (CORS, authentification, logs, gestion des erreurs, etc.).
- **Router/** : Définition des routes, du serveur et de la logique de routage.

---

## Détail des packages et fonctions

### 1. Context/

#### context.go

Contient la structure centrale `Context` qui encapsule la requête, la réponse, les paramètres, le body, les headers, etc. Elle permet de :

- Stocker et partager des données tout au long du cycle de vie d'une requête.
- Fournir des méthodes pour accéder facilement aux paramètres, au body, aux headers, etc.
- Gérer le contexte utilisateur (authentification, permissions, etc.).

#### file.go

Fonctions pour la gestion avancée des fichiers :

- `SaveUploadedFile` : Sauvegarde un fichier uploadé sur le serveur.
- `GetFileMimeType` : Détecte le type MIME d'un fichier.
- `ReadFile` : Lit le contenu d'un fichier et le retourne sous forme de bytes ou de string.

#### json.go

Fonctions pour la sérialisation/désérialisation JSON :

- `BindJSON` : Parse le body d'une requête en une structure Go.
- `ToJSON` : Sérialise une structure Go en JSON.
- `FromJSON` : Désérialise du JSON en structure Go.

#### jwt.go

Gestion des tokens JWT :

- `GenerateJWT` : Génère un token JWT à partir d'un utilisateur ou d'une payload.
- `ValidateJWT` : Vérifie la validité d'un token JWT (signature, expiration, etc.).
- `ParseJWT` : Extrait les claims d'un token JWT.

#### response.go

Gestion des réponses HTTP :

- `JSON` : Envoie une réponse JSON avec le bon status code.
- `Error` : Envoie une réponse d'erreur formatée.
- `File` : Envoie un fichier en réponse HTTP.

#### types.go (Context)

Définit les types utilisés dans le package Context, comme les structures de réponse, les claims JWT, etc.

#### utils.go (Context)

Fonctions utilitaires diverses, par exemple :

- Génération d'UUID.
- Fonctions d'aide pour manipuler les headers ou les cookies.

---

### 2. Helpers/

#### validators.go

Contient toutes les fonctions de validation utilisées dans le projet :

- `IsEmailValid(email string) bool` : Vérifie si un email est valide.
- `IsPasswordStrong(password string) bool` : Vérifie la robustesse d'un mot de passe.
- `IsUUIDValid(uuid string) bool` : Vérifie si une chaîne est un UUID valide.
- Fonctions personnalisées pour valider des champs spécifiques selon les besoins métier.

---

### 3. Middleware/

#### cors.go

Middleware qui gère les headers CORS pour permettre ou restreindre l'accès à l'API depuis d'autres domaines.

- Ajoute les headers `Access-Control-Allow-Origin`, `Access-Control-Allow-Methods`, etc.

#### error.go

Middleware de gestion centralisée des erreurs :

- Intercepte les erreurs non gérées et renvoie une réponse formatée.
- Peut logger les erreurs pour le debug.

#### jwtAuth.go

Middleware d'authentification JWT :

- Vérifie la présence et la validité du token JWT dans les headers.
- Ajoute les informations utilisateur dans le contexte si le token est valide.

#### logger.go

Middleware de logging :

- Log chaque requête entrante (méthode, chemin, durée, code de réponse, etc.).
- Peut logger les erreurs ou les accès non autorisés.

#### recover.go

Middleware de récupération après panic :

- Intercepte les panics pour éviter que le serveur ne crash.
- Renvoie une réponse d'erreur 500 et log le panic.

#### requestID.go

Middleware qui génère un identifiant unique pour chaque requête :

- Ajoute un header `X-Request-ID` à chaque réponse.
- Permet de tracer les requêtes dans les logs.

#### types.go (Middleware)

Définit les types utilisés par les middlewares (ex : structure d'erreur, structure de log, etc.).

#### upload.go

Middleware pour la gestion de l'upload de fichiers :

- Gère la réception, la validation et le stockage des fichiers uploadés.

#### utils.go (Middleware)

Fonctions utilitaires pour les middlewares, par exemple :

- Extraction de l'IP client.
- Fonctions d'aide pour manipuler les headers.

---

### 4. Router/

#### router.go

Fichier principal de définition des routes :

- Déclare toutes les routes de l'API (GET, POST, etc.).
- Associe chaque route à un handler et aux middlewares nécessaires.

#### routerGroup.go

Permet de regrouper des routes par fonctionnalité ou par niveau de protection :

- Groupes publics, groupes protégés par authentification, etc.
- Application de middlewares spécifiques à un groupe de routes.

#### server.go

Initialisation et démarrage du serveur HTTP :

- Configure le port, les middlewares globaux, le router principal.
- Démarre le serveur et gère l'arrêt propre.

#### types.go (Router)

Définit les types utilisés pour le routage (ex : structure de route, configuration du serveur, etc.).

#### utils.go (Router)

Fonctions utilitaires pour le routage :

- Génération dynamique de routes.
- Fonctions d'aide pour la gestion des paramètres de route.
