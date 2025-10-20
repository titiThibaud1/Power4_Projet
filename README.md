# Projet Puissance 4

## Qu'est ce que le Puissance 4 ?

Le Puissance 4 est un jeu de stratégie.
Le but du jeu est d'aligner une suite de 4 pions de même couleur sur une grille comptant 6 rangées et 7 colonnes.

Chaque joueur dispose de 21 pions d'une couleur (par convention, en général jaune ou rouge).
Tour à tour, les deux joueurs placent un pion dans la colonne de leur choix, le pion coulisse alors jusqu'à la position la plus basse possible dans ladite colonne à la suite de quoi c'est à l'adversaire de jouer. 

Le vainqueur est le joueur qui réalise le premier un alignement (horizontal, vertical ou diagonal) consécutif d'au moins quatre pions de sa couleur. 
Si, alors que toutes les cases de la grille de jeu sont remplies, aucun des deux joueurs n'a réalisé un tel alignement, la partie est déclarée nulle.


## Présentation du projet

Nous devons créer un jeux de Puissance 4 en language Golang et HTML.

Ce jeu doit respecter les conditions suivantes :

<img width="698" height="450" alt="image" src="https://github.com/user-attachments/assets/d3ee69c3-4b9f-4211-a293-79409aeaa008" />

<img width="840" height="537" alt="image" src="https://github.com/user-attachments/assets/639ca0f3-c7ed-454b-a804-1e553c0647cc" />


## Comment jouer ?

#### Acceder au jeu
- Exécutez le fichier main.go dans un terminal via la commande "go run main.go"
- Ouvrez votre navigateur et entrez l'adresse "localhost:8080" dans la barre de recherche par adresse
- Vous devrez arriver sur la page d'accueil du jeu

#### Lancer une partie
- Entrez dans les cases prévues à cet effet, le nom des 2 joueurs qui s'affronteront
- Cliquez sur "Next"
- Choisissez la difficulté du jeu étant "Easy" par défault
- Cliquez sur "Start Game"

#### Déroulement d'une partie
- Arrivé devant le tableau, si ce n'est pas déjà le cas, dezoomer sur votre navigateur jusqu'à que votre résolution permet d'afficher l'ensemble des éléments (jusqu'aux 3 bouton d'en bas) sur l'écran sans besoin de scroll
- Le premier joueur choisi une colone sur laquelle il va jouer grace aux flèches en haut de chaque colone
- Une fois joué, c'est au tour du deuxième joueur et ainsi de suite
- Jouez votre meilleure stratégie
- Une fois un vainceur désigné ou la grille étant remplie, vous êtes redirigé sur un écran de victoire ou d'égalité.
- Vous disposez, comme pendant le déroulement de la partie de boutons permetant chacun de recommencer, changer la difficulté ou de revenir sur l'écran d'accueil
- Faites autant de parties que vous le souhaitez


## Présentation du travail réalisé

### Page d'accueil
Sur la page d'accueil vous pouvez y entrer les nom des 2 joueurs qui vont s'affronter sur le jeu de Puissance 4.

Le bouton "Next" permet de passer a l'étape suivante et enregistrer les informations écrites.


### Page du choix de la difficulté
Sur la page du choix des difficulté, vous avez un appercu des nom des joueurs écrit lors de l'étape précédente.

Vous avez également, grace à des input "radio", le choix entre 3 difficulté qui chacune impacte la taille de la grille de jeu :

- Easy : 6*7
- Medium : 6*9
- Hard : 7*8


### Page de jeu
Sur la page de jeu, vous avez en haut de l'écran un rappel de la difficulté dans laquelle vous joué.

Vous avez également une grille de jeu, qui en cliquant sur les bouton en haut de chaque colone, permet de place un jeton de la couleur du joueur actuel dans la case ibre la plus basse de la colone en question.

Si un joueur alligne 4 jeton de la même couleur sur la grille, vous êtes redirigé vers une page de victoire.

Si la grille est complète sans qu'un joueur ait alligné 4 jetons de la même couleur, vous êtes redirigé vers une page d'égalité.

Enfin, en bas de la page vous avez des bouton permettant chacun de :

- "Reset"               -->     Recommencer la partie de 0 (conserver la difficulté et les joueurs)
- "Change difficulty"   -->     Changer la difficulté (conserver les joueurs)
- "Home"                -->     Revenir sur l'écran d'accueil


### Page de victoire / égalité
En cas de victoire d'un joueur, vous avez à l'écran le pseudo du gagnant ansi que les 3 même boutons que précedement
En cas d'égalité, seulement un message "Draw" est affiché, toujours avec les 3 boutons.


## Information générales

- Vous devez naviguer entre les pages uniquement avec les boutons prévus à cet effet, ne pas cliquer sur la flèche de retour de votre navigateur ou alors de passer par les routes via l'URL au risque de perdre les informations des pages précédentes.
