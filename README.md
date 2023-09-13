# github.com/E-kenny/StagetwoTask
* Welcome to the task two/rest generated docs.

- Clone the main branch
- Cd into the project directory
- Set the environment variable to connection string of a postgres database e.g DATABASE_URL = postgres://postgres:E_kenny246810@localhost:5432/StageTwoDB
- go run main.go 

## Host URL
- [https://stagetwotask.onrender.com/api](https://stagetwotask.onrender.com/api)

## Postman Documentation
- [Documentation](https://documenter.getpostman.com/view/11374327/2s9YC32Zfu)

## UML
- [UML Diagram](https://lucid.app/lucidchart/1d1e66f4-1055-45c7-95ee-0a2a0bf1d2b1/edit?viewport_loc=-1449%2C-774%2C2694%2C1192%2C0_0&invitationId=inv_17705617-2424-4865-ad84-6d3109508cea)


## Routes

<details>
<summary>`/api`</summary>

- [RequestID]()
- [RealIP]()
- [Logger]()
- [Recoverer]()
- **/api**
        - **/**
                - _GET_
                        - [Paginate]()
                        - [Listpersons]()
                - _POST_
                        - [Createperson]()

</details>
<details>
<summary>`/api/{param}`</summary>

- [RequestID]()
- [RealIP]()
- [Logger]()
- [Recoverer]()
- **/api**
        - **/{param}**
                - [PersonCtx]()
                - **/**
                        - _PUT_
                                - [Updateperson]()
                        - _DELETE_
                                - [Deleteperson]()
                        - _GET_
                                - [Getperson]()

</details>


