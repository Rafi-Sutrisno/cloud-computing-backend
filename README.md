# InsCure Backend RestFUL API : Golang, GORM (Object Relational Mapping), and GIN.

This repository contains all the code that we built to develop RestFUL API with go-gin-gorm. In order to replicate this whole project, make sure you are following this guide that we've provided below.

## Repository Contents

- [Requirements](#requirements)
- [Setting Up Google Cloud Platform](#setting-up-google-cloud-platform)
- [Setting Up Golang Environment in GCE](#setting-up-golang-environment-in-gce)
- [Deployment and Integration](#deployment-and-integration)
- [API Testing](#api-testing)

## Requirements

In order to run through this whole tutorial, make sure you meet these requirements:

- Running desktop/laptop with internet connection
- A Web Browser
- Postman

## Setting Up Google Cloud Platform

**New Project and Billings**
1. In your web browser, Go to [Google Cloud Platform](https://console.cloud.google.com/)
2. Register and Login with your account.
3. Create a new project.
4. Create a new billing account and redeem (if you haven't) the [free](https://cloud.google.com/free?hl=en) $300 worth of credits.

**Creating PostgreSQL Instance**
5. On the sidebar of the web, open Cloud SQL.
6. Create a new PostgreSQL instance with the credentials and settings as you wish. Then, create a new db called ``inscure``
7. For your information, this instance will store all of the tables that we define in ``/entity`` directory.

**Creating Cloud Storage Bucket**
8. On the sidebar of the web, open Cloud Storage -> Buckets.
9. Create a new bucket with the credentials and settings as you wish.
10. Inside that bucket, create 2 folders called ``prediction`` and ``profile_picture``. 
11. Upload a single file in the root directory for user default profile picture and rename it to ``default.png``
12. Its pretty self explanatory to know the purpose of this bucket by doing this.

**Creating Compute Engine Instance**
13. On the sidebar of the web, 
  - open IAM -> Service Account and create new account
  - open VPC -> Firewall and create a new firewall that allow tcp port 8080 with target name as you wish.
14. Go to Compute Engine -> VM Instances then Create a new instance with the credentials and settings as you wish.
  - make sure to use the firewall with the specified target-name.
15. After the process has finished, set a Reserve Static IP for the instance then connect to it via SSH.
16. Inside the SSH, run ``sudo apt-get update``, then install golang and pm2.
17. For go installation please refer to this [doc](https://go.dev/doc/install)
18. For pm2 installation,
  - Install nvm and nodeJS by referring to this [doc](https://gist.github.com/d2s/372b5943bce17b964a79)
  - Install pm2 by entering ``npm i pm2 -g``
19. Inside your root directory, clone this repository by using ``git clone <repo-url>``.

## Setting Up Golang Environment in GCE

20. Navigate to ``/inscure-fe`` directory, then enter ``nano .env``. Copy and fill this snippet based on your postgre instance,
```
DB_HOST=<your postgre public ip>
DB_USER=postgres
DB_PASS=<your instance password>
DB_NAME=inscure
DB_PORT=5432
```
21. Save and Close the file.
22. Go through all the files inside ``/service`` and ``/utils`` directory, then replace all strings containing ``example-bucket-test-cc-trw`` with your bucket name.

## Deployment and Integration

23. Navigate to ``/inscure-fe`` directory, and enter ``nano package.json`` then copy this snippet inside it.
```json
{
  "apps": [
    {
      "name": "main",
      "script": "go run main.go"
    }
  ]
}
```
23. Save and close the file.
24. Inside the same directory, run ``pm2 start package.json``
25. Voila! You've managed to deploy our RestFUL API by with GCP.

## API Testing

To test the API, copy the external IP address of your compute engine instance, then refer to this link for available api endpoints
[Postman Documentation](https://documenter.getpostman.com/view/25927897/2sA3XV7eCN)

