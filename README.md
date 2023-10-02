# JobBoard Api Using GraphQL ,MongoDB, and Golang.
Create a new folder for the Project mkdir gql-yt

Mod init your project, give it whatever name you like go mod init github.com/akhil/gql-yt

Get gql gen for your project go get github.com/99designs/gqlgen

Add gqlgen to tools.go printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go

Get all the dependencies go mod tidy

Initialize your project go run github.com/99designs/gqlgen init

After you've written the graphql schema, run this - go run github.com/99designs/gqlgen generate

After you've built the project, these are the queries to interact with the API -

Get All Jobs
query GetAllJobs{ jobs{ _id title description company url } }

=======================

Create Job
mutation CreateJobListing($input: CreateJobListingInput!){ createJobListing(input:$input){ _id title description company url } }

{ "input": { "title": "Software Development Engineer - I", "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt", "company": "Google", "url": "www.google.com/" } }`

=========================

Get Job By Id
query GetJob($id: ID!){ job(id:$id){ _id title description url company } }

{ "id": "638051d7acc418c13197fdf7" }

=========================

Update Job By Id
mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) { updateJobListing(id:$id,input:$input){ title description _id company url } }

{ "id": "638051d3acc418c13197fdf6", "input": { "title": "Software Development Engineer - III" } }

=================================

Delete Job By Id
mutation DeleteQuery($id: ID!) { deleteJobListing(id:$id){ deletedJobId } }

{ "id": "638051d3acc418c13197fdf6" }

## Database Configration

--Start MongoDB.
    ---You can start the mongod process by issuing the following command:
     
      ```
      sudo systemctl start mongod
      ```

If you receive an error similar to the following when starting mongod:

Failed to start mongod.service: Unit mongod.service not found.

Run the following command first:

sudo systemctl daemon-reload

Then run the start command above again.

2
Verify that MongoDB has started successfully.
sudo systemctl status mongod

You can optionally ensure that MongoDB will start following a system reboot by issuing the following command:

sudo systemctl enable mongod

3
Stop MongoDB.
As needed, you can stop the mongod process by issuing the following command:

sudo systemctl stop mongod

4
Restart MongoDB.
You can restart the mongod process by issuing the following command:

sudo systemctl restart mongod

You can follow the state of the process for errors or important messages by watching the output in the /var/log/mongodb/mongod.log file.

5
Begin using MongoDB.
Start a 
mongosh
 session on the same host machine as the mongod. You can run 
mongosh
 without any command-line options to connect to a mongod that is running on your localhost with default port 27017.

mongosh

For more information on connecting using 
mongosh
, such as to connect to a mongod instance running on a different host and/or port, see the 
mongosh documentation.

To help you start using MongoDB, MongoDB provides Getting Started Guides in various driver editions. For the driver documentation, see 
Start Developing with MongoDB.
