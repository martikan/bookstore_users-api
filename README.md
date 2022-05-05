# Microservice for Bookstore - Users-API

## Index

1 [Setup local development environment](#setup-local-development-environment)<br/>
2 [Run the API with Docker](#run-the-api-with-docker)<br/>
3 [Run the API with Kubernetes](#run-the-api-with-kubernetes)<br/>

## Setup local development environment

### Setup vscode as workspace

* Create launch.json file for running GO

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}",
            "env": {
                "mysql_users_db_username": "<<your_database_username>>",
                "mysql_users_db_password": "<<your_database_password>>",
                "mysql_users_db_uri": "<<your_database_host>>",
                "mysql_users_db_database": "<<your_database_name>>"
            }
        }
    ]
}
```
* Add the working directory as a workspace to vscode: `File >> Add Folder to Workspace...`

* You can run the project with `Control + F5` & You can run it in debug mode `F5`

## Run the API with Docker

* Not implemented

## Run the API with Kubernetes

* Not implemented