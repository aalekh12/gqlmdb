# JobBoard 
In this Api we are performing crud operation using Graphql , Mongodb and Golang.



=======================

## Create Job

```
mutation CreateJobListing($input: CreateJobListingInput!){ createJobListing(input:$input){ _id title description company url } }
```
You can use this input

```
{ "input": { "title": "Software Development Engineer - I", "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt", "company": "Google", "url": "www.google.com/" } }`
```
=========================

## Get All Jobs

```
query GetAllJobs{ jobs{ _id title description company url } }
```
No need input Here



=========================

## Get Job By Id

```
query GetJob($id: ID!){ job(id:$id){ _id title description url company } }
```
You can use this input

```
{ "id": "638051d7acc418c13197fdf7" }
```

=========================

## Update Job By Id

```
mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) { updateJobListing(id:$id,input:$input){ title description _id company url } }
```

You can use this input 

```
{ "id": "638051d3acc418c13197fdf6", "input": { "title": "Software Development Engineer - III" } }
```

=================================

## Delete Job By Id
```
mutation DeleteQuery($id: ID!) { deleteJobListing(id:$id){ deletedJobId } }
```

You can use this input

```
{ "id": "638051d3acc418c13197fdf6" }
```

# Database Configration

--Start MongoDB--

## You can start the mongod process by issuing the following command:

```
sudo systemctl start mongod
```

## If you receive an error similar to the following when starting mongod:

Failed to start mongod.service: Unit mongod.service not found.

Run the following command first:

```
sudo systemctl daemon-reload
```
 Then run the start command above again.


## Verify that MongoDB has started successfully.

```
sudo systemctl status mongod
```

## You can optionally ensure that MongoDB will start following a system reboot by issuing the following command:

```
sudo systemctl enable mongod
```


## Stop MongoDB.
## As needed, you can stop the mongod process by issuing the following command:

```
sudo systemctl stop mongod
```

## Restart MongoDB.
## You can restart the mongod process by issuing the following command:

```
sudo systemctl restart mongod
```

You can follow the state of the process for errors or important messages by watching the output in the /var/log/mongodb/mongod.log file.


## Begin using MongoDB.
By Running Following command you are able to see the configration

```
mongosh
```

For more information on connecting using 
mongosh
, such as to connect to a mongod instance running on a different host and/or port, see the 
mongosh documentation.

To help you start using MongoDB, MongoDB provides Getting Started Guides in various driver editions. For the driver documentation, see Start Developing with MongoDB.
