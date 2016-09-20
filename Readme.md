# Gantry - Ported to go

[Originally](https://github.com/Jason-Hendry/gantry) written in bash scripts this project aims to create a robust cross platform DevOps tools to bootstrap, build, deploy and manage a variable of web and mobile projects.

## Roadmap

### New Project
* [ ] Symfony (MySQL/Postgres)
* [ ] Wordpress
* [ ] Cordova / Iconic / Phonegap
* [ ] Select staging env
* [ ] Select production env (ECS, GCP, Kubernetes etc)
* [ ] Polyglot setups (Node + PHP + go + ...)

### Provisioning
* [ ] Cloudflare
* [ ] AWS (EC2, ECS)
* [ ] GCP
* [ ] Heroku
* [ ] Ubuntu (apt, files/folders etc)

### Dev Tools
* [ ] Console
* [ ] Console access to other containers eg gantry console api
* [ ] Open in web
* [ ] DB Tools (phpMyAdmin, mysql, psql)
* [ ] Execute SQL on DB
* [ ] Mock Email / SMS services for dev

### Ops
* [ ] Config services like DB, Cache, Queue, Jobs
* [ ] Backup DB (Prod, Staging, Dev)
* [ ] Restore DB (Prod, Staging, Dev)
* [ ] Sync DB Prod > Staging
* [ ] Sanatize DB (Staging, Dev)

### Build
* [ ] Generate build configs (CircleCI, docker?)
* [ ] OSX Build (CirclCI)

### Deployments
* [ ] Simple Docker - Push to Repo (tag/latest), Pull on server, Scale up new image, Scale back old images (replicate existing process)
* [ ] ECS update containers in AWS
* [ ] Kubernetes (GCP or hosted)
* [ ] Google ?
* [ ] Heroku ?

### Docker command wrapper

* [ ] AWS-Cli
* [ ] composer (php)
* [ ] npm
* [ ] node
* [ ] sass
* [ ] cap ? 
* [ ] ansible ? Maybe for provisions
* [ ] behat ?
* [ ] bower
* [ ] gulp
* [ ] cordova

### Rapid Prototyping

* [ ] Symfony Entity / Crud command wrapper
* [ ] Symfony Migration tools
* [ ] Symfony Create / Update users
