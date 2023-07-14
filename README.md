# hackaton
## Hypothesis:
  Quarkus Java and GoLang are made for cloud native devlopment and we would like to validate claims that they can beat traditional Java on speed of development and application run-time permance
## Prerequisite:
  * Docker on Linux and Docker Desktop on Windows or MAC
  * Min 8GB RAM
  * Visual Code
## Repo content for PAL hackaton and general info
  Working enviroment is made of DB container and dev Linux contaier.\
  Pls use provided yaml file to pull neccessary containers to form your working enviroment.\
  Create new folder and copy yaml file there, cd into the folder and run <docker compose -d> \
  Once images downoalded and started you can run bash shell inside by runing <docker exec -ti id-of-the-container bash>\

  .zip file contains Quarkus scafolding project (unzip it in dev container) and cd inside folder adn you can run <mvn compile quarkus:dev> 

  Visual Code downalod Remote Development extension to connect to dev container and dev code inside container \

  Assistance for container creation will be available on the day by support team 

  ## Postgres DB Information
  
CREATE TABLE transaction.transactions (\
	tx_id varchar(40) primary key,  /* transaction id*/ \
	acc_id varchar(40) not null, /* account id*/ \
	tx_ts timestamp not null, /* timestamp */ \
	status varchar(40) not null, /* can be ["pending","rejected","complete"] */ \
	amount int not null, /* integer - need divide by 100 for pennies*/ \
	merchantname text,  \
	merchant_id varchar(40), \
	tx_type varchar(40), /* ["online","atm","pos","dd","transfer"] */ \
	tx_details text);  /* random text*/ 

Example of rows: \
ed9b9bcf-731d-460d-add5-ea60fde769da|ADUG99161510903217|2023-01-27 09:27:46|pending|6792|Avery, Middleton and Foster|501-90443-X|online|Old great notice. \
6f66d3d1-25fb-432e-842e-b349ac1dde63|FBTV47112201868483|2023-02-03 17:34:46|rejected|8221|Arnold, Villa and Johnson|616-57004-3|transfer|Fire town worker. 

### Connection Details:
DB_Host ='pgdb' \
DB Port =5432 \
DB User ='postgres' \
DB Password ='postgres' \
DB Database ='workshop' 

## Rest API
GET method for 

### Must have:
(note happy to truncate to max 50 rows) \
http://{hostname}:{port}/**transactions/account/{acc_id}** \
returns all tx for account, if added optional status param (e.g. &status="complete") then filter also by status 


http://{hostname}:{port}/**transactions/account/{acc_id}?fromDateTime="yyyy-mm-dd hh:mm:ss"** \
returns all tx for account from datetime to "now", if added optional status param (e.g. &status="complete") then filter also by status  


http://{hostname}:{port}/**transactions/account/{acc_id}?fromDateTime="yyyy-mm-dd hh:mm:ss"&toDateTime="yyyy-mm-dd hh:mm:ss"**
returns all tx for account from datetime to datetime, if added optional status param (e.g. &status="complete") then filter also by status 


http://{hostname}:{port}/**transactions/{tx_id}** \
return single transaction


### Should have:
(on purpose less defined - leave to you how to do it) \
Pagination for calls returning max "pageSize" rows with query "pageSize" param and allows "page" inout param to select which "startPage" to show arbirtrary page 

## Goal
Docker image pushed in dockerhub with your application implementing Rest API

## Useful Links
### Quarkus
https://quarkus.io/guides/hibernate-orm /
https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=video&cd=&cad=rja&uact=8&ved=2ahUKEwjy5d_cio6AAxVJQkEAHR7AA7AQtwJ6BAgLEAI&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DkAui1-4KBrk&usg=AOvVaw3brk2lKl1_yHlVy-Pv7YrJ&opi=89978449 /
or  \
https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=video&cd=&cad=rja&uact=8&ved=2ahUKEwjy5d_cio6AAxVJQkEAHR7AA7AQtwJ6BAgHEAI&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3Dzs8aN5g0lr0&usg=AOvVaw0oNnVXFUoq4jU_cU2Ji7LW&opi=89978449 \


### GoLang
https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=video&cd=&cad=rja&uact=8&ved=2ahUKEwjtoemli46AAxXGS0EAHRX_Aq0QtwJ6BAgSEAI&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DqR0WnWL2o1Q&usg=AOvVaw3y7d5qfvRsFgzDKxaOc45v&opi=89978449 \
https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=video&cd=&cad=rja&uact=8&ved=2ahUKEwjtoemli46AAxXGS0EAHRX_Aq0QtwJ6BAgREAI&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D_GE-8Db1JNQ&usg=AOvVaw0pxFJ9QNhOdkESz-IHdBWd&opi=89978449 \
https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=video&cd=&cad=rja&uact=8&ved=2ahUKEwjXu6nIi46AAxUWUkEAHVkQC8QQtwJ6BAgLEAI&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DDdF5P7izAtE&usg=AOvVaw3JmtkCLeMABJmfkKxrtGTw&opi=89978449 \
https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=video&cd=&cad=rja&uact=8&ved=2ahUKEwjXu6nIi46AAxUWUkEAHVkQC8QQtwJ6BAgKEAI&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D8XtC4ha7SfI&usg=AOvVaw0QU2fVaD25m8DokNyDrI56&opi=89978449 \


### Visual Code Remote Dev
https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=video&cd=&cad=rja&uact=8&ved=2ahUKEwiZ98SNi46AAxWcQUEAHUo9C9oQtwJ6BAgMEAI&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DsakjpegUQsk&usg=AOvVaw1S6MvQP_SCv5Xc3WwQ56jF&opi=89978449 \
