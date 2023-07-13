# hackaton
## Prerequisite:
  * Docker on Linux and Docker Desktop on Windows or MAC
  * Min 8GB RAM
  * Visual Code
## Repo content for PAL hackaton
  Working enviroment is made of DB container and dev Linux contaier.\
  Pls use provided yaml file to pull neccessary containers to form your working enviroment.\
  Create new folder and copy yaml file there, cd into the folder and run <docker compose -d> \
  Once images downoalded and started you can run bash shell inside by runing <docker exec -ti id-of-the-container bash>\

  .zip file contains Quarkus scafolding project (unzip it in dev container) and cd inside folder adn you can run <mvn compile quarkus:dev> \

  Visual Code downalod Remote Development extension to connect to dev container and dev code inside container \
