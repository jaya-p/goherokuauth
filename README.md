# goherokuauth
This git repository is forked from jaya-p/gowebservice. The purpose is for adaptation for deploying to heroku.com platform.  

## Deploy to Heroku (Using Heroku CLI)
heroku login  
heroku container:push web -a <app-name>  
heroku container:release web -a <app-name>  
heroku logs --tail -a <app-name>

## Access deployed application
https://<app-name>.herokuapp.com/  
https://<app-name>.herokuapp.com/api/v1/auth
