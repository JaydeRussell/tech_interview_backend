# Tech interview Project

This project was created as a tech interview solution. It's separated into two parts: a [react frontend](https://github.com/JaydeRussell/tech_interview_frontend) and a [golang backend](https://github.com/JaydeRussell/tech_interview_backend). All data is created by the golang backend and persisted in a postgres database which is then maintained by the golang code.

NOTE: The README on both repos are identical to ensure the project is setup correctly.

# Requirements
These are the versions I built with, not necessarly the minimum requirements, but I won't gurantee anything below these versions.

```
docker  v20.10.22
Go      v1.19.5
NPM     v9.3.0
Node    v18.13.0
```

# Running the project
## Step 1: Starting the postgres database
In the backend project there's a `docker-compose.yml` file that contains the logic to start the postgres database. It can be run with the following command: `docker-compose -f docker-compose.yml up`.


## Step 2: Starting the Golang Backend
Starting the backend is simple, just navigate to the frontend directory then run `go get` and `go run main.go`

### Things to watch out for... 
If there's no database available, then the app will not start. Since this project is a glorified CRUD api, I consider this a feature. 

Also, by default the backend will attempt to connect to a database provided by the `DB_CONNECTION` environment variable. If one is not available then it will connect to a hardcoded DB connection string defined in the `config.go` file as a constant. This connection string points to the postgres database from the `docker-compose.yml` file available in the root of the project directory. 

**As a note:** Normally, pushing a db connection string to a public github repo is a big no-no, but I consider this permissible since it's a _very_ generic connection string to a locally defined database. I would never do something like this if it were any other circumstance.

## Step 3: Starting the react frontend
Starting the react frontend is also very simple, just navigate to the frontend directory then run `npm install` and `npm start`

### Things to watch out for...
The frontend _should_ be started after the backend is already up and running, however most (if not all) of the code accounts for the possibility of the backend not existing. The app probably won't look pretty, but it also shouldn't crash and you should be able to just refresh the browser page after the backend is up and running.

There are also a few hidden backend features that are hardcoded on the frontend just because I didn't have time to implement them. These include things like additional pagination options, creating questions, and adding answers to questions. Given an extra day I could probably add these features, and I honestly might just do it for fun later anyway, but I was already pushing the time limit. So, they'll have to wait.



# Technology usage explanations
Here's just a brief explanation for the decisions I made along the way. I feel like this is fairly comprehensive, but if you have any questions about what I did, or why, then feel free to reach out to me! I'm always happy to talk about programming!
### Frontend
#### React
For the frontend I decided to use react, or more specifically [Create React App](https://create-react-app.dev/), I did this for a couple of reasons. Mostly just because it handles all of the boilerplate of a react app, and while I probably wouldn't use it to create an optimized production ready website, it was more than enough for this project.

#### CSS
I let [react-boostrap](https://react-bootstrap.github.io/) do most of the css just because I don't consider frontend development to be my specialty and react-bootstrap made quick work of the ui elements. I had to modify some of it to get a slightly better look, but definitely not much.


### Backend
#### GO
[Go](https://go.dev/) is my go-to (ha!) backend language, and just one that I tend to prefer. It has a lot of built-in RESTful support and is really good at creating these sorts of micro-services.

#### Gin Server
I used [Gin](https://gin-gonic.com/) as a wrapper to the web-server just because it's lightweight, but still handles a significant amount of the boilerplate. It also makes adding middleware extremely easy. 

In theory, I could have gone for a lambda-based serverless architecture, but my only experience doing that has been through [AWS-lambda](https://aws.amazon.com/lambda/) and I wanted to make this project entirely self-contained.

#### Postgres
I used [Postgres](https://www.postgresql.org/) because to me the project felt fairly relational (questions having a predictable one-to-many relationship to answers), and Postgres was just the most recent SQL language that I've used.

#### ORM and SQL migrations
I used [GORM](https://gorm.io/) as a simple go ORM to make database queries simpler and more readable. It also does a pretty dang good job at preventing sql-injection, which is an obvious bonus.

Despite using GORM for the queries, I didn't use it to set up the table definitions. While it's very capable of building the tables based off of a provided object class, I prefer taking a migration based approach to maintaining database tables. This is because it makes it really simple to view  table definitions and track change history. 

Migrations also make it really easy to spin-up/down entire databases whenever you need them and populate data into a table without having to do so programmatically.