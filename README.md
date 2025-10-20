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


## Présentation du travail réalisé



## Comment jouer ?

#### Acceder au jeu
- Exécutez le fichier main.go dans un terminal
- Ouvrez votre navigateur et entrez l'adresse "localhost:8080" dans la barre de recherche par adresse
- Vous devrez arriver sur la page d'accueil du jeu

#### Lancer une partie
- Entrez dans les cases prévues à cet effet, le nom des 2 joueurs qui s'affronteront
- Cliquez sur suivant
- Choisissez la difficulté du jeu étant "Easy" par défault
- Cliquez sur "Start Game"

#### Déroulement d'une partie
- Arrivé devant le tableau, si ce n'est pas déjà le cas, dezoomer sur votre navigateur jusqu'à que votre résolution permet d'afficher l'ensemble des éléments (jusqu'aux 3 bouton d'en bas) sur l'écran sans besoin de scroll
- Le premier joueur choisi une colone sur laquelle il va jouer grace aux flèches en haut de chaque colone
- Une fois joué, c'est au tour du deuxième joueur et ainsi de suite
- Jouez votre meilleure stratégie
- Une fois un vainceur désigné ou la grille étant remplie, vous êtes redirigé sur un écran de victoire ou d'égalité.
- Vous disposez, comme pendant le déroulement de la partie de bouton permetant chaque de recommencer, changer la difficulté ou de revenir sur l'écran d'accueil
- Faites autant de parties que vous le souhaitez
