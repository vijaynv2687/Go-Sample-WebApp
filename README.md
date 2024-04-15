## Go-Sample-WebApp


This repo provides all the necessary files and code to be able to understand the basic requirements for deploying a docker image of sample registration form, with HTML file, backed golang code and dockerfile.

# Install all required dependencies

##
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz  
##
    yum -y install docker  
##
    yum -y install nginx  


# Build docker image  

##
    docker build . -f dockerfile -t webapp  

# Start Web Server  
##
    systemctl start nginx  
##
    systemctl status nginx  
# Test  

# To list if there any docker images running  
##
        docker ps 
# To list if there are any docker images running or exited  
##
        docker ps -a 
# To list the docker iamges built  
##
        docker images  
# To run the docker image to bring up the webapp  
##
        docker run -d -p 8080:3000 localhost/webapp
#Open browser to launch the webapp(Registration page loads up) locally --> use docker host IP Address in the below URL:  

http://ipadd:8080
