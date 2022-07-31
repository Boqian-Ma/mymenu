# MyMenu - a restaurant ordering system written in GoLang and React.js

MyMenu


## Installation 

#### Install and setup VirtualBox #### 
Follow installation and setup instructions
https://webcms3.cse.unsw.edu.au/static/uploads/course/COMP9900/21T2/f27d53fa0797fbf9320dbe4b8ebc860988729b919247e127ed856cb3948066f4/VirtualBoxGuide.pdf

#### Clone the git repo ####
```bash
cd ~
git clone https://github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar.git
```
    

#### Set up SSH (Optional) ####
Set up SSH to make dev easier with your own command line rather than command line on the VM GUI  
1. Before starting VM, open settings -> Network -> Adapter 1 -> Attached to and select Bridged Adapter
2. Update and upgrade
```bash
sudo apt update
sudo apt upgrade
```
3. Install openssh
```bash
sudo apt install openssh-server
```
4. Verify that SSH server is running
```bash
sudo service sshd status
```
5. Configure firewall and open port 22
```bash
sudo ufw enable
sudo ufw allow ssh
```
6. Check that SSH is correctly configured
```bash
sudo ufw status
```
7. Find IP address of VM
```bash
ip a | grep "inet 192"
# if more than 1 result is printed, either should work
```
8. Connect to VM via host machine with password 'lubuntu'
```bash
ssh lubuntu@<ip address of vm>
```

8. Add host SSH keys to ~/.ssh/authorized_keys to eliminate need to enter password

### Install packages ###
#### Update and Upgrade ####
```bash
sudo apt-get update
sudo apt-get upgrade
```
#### Git ####
```bash
sudo apt-get install git
```
#### npm and yarn ####
```bash
sudo apt-get install npm
sudo npm install -g yarn
```
#### curl ####
```bash
sudo apt install curl
```
Restart the shell
#### react-scripts and typescript ####
```bash
cd ~/capstoneproject-comp3900-w16a-jamar/frontend
npm install react-scripts
yarn add typescript
yarn add qrcode.react
```
#### Golang ####
1. Download Go binary
```bash
cd ~
wget https://dl.google.com/go/go1.16.5.linux-amd64.tar.gz
```
2. Verify Go tarball
```bash
sha256sum go1.16.5.linux-amd64.tar.gz
```
Output should look something like
```bash
Output
b12c23023b68de22f74c0524f10b753e7b08b1504cb7e417eccebdd3fae49061  go1.16.5.linux-amd64.tar.gz
```
3. Extract Go tarball
```bash 
sudo tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
```
4. Add Go to $PATH
```bash
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
```

5. Load new path into current shell
```bash
source ~/.profile
```

6. Verify Go installation
```bash
go version
```
Should have output
```bash
go version go1.16.5 linux/amd64
```

7. Install Golang database library
```bash
go get -u github.com/lib/pq
```

#### Go Migrate ####
1. Install migrate
```bash
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash

sudo apt install migrate
```

#### PostgreSQL ####
1. Install PostgreSQL
```bash
sudo apt install postgresql postgresql-contrib
```
2. Check PostgreSQL is running
```bash
sudo service postgresql status
```
If not on, run 
```bash
sudo service postgresql start
```
3. Login to postgres and open psql shell
```bash
sudo su postgres
psql
```
Can inspect users with `\du` in psql shell  
4. Give user postgres a password
```sql
ALTER USER postgres WITH PASSWORD 'password';
```
5. Exit back to user lubuntu
#### Docker ####
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04
1. Update
```bash
sudo apt update
```
2. Install prerequisite packages
```bash
sudo apt install apt-transport-https ca-certificates curl software-properties-common
```
3. Add GPG key
```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```
4. Add Docker repo
```bash
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"
```
5. Update with new Docker packages and install Docker
```bash
sudo apt install docker-ce
```
6. Verify that Docker is running
```bash
sudo service docker status
```
#### Docker Compose ####
1. Download docker-compose
```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```
2. Apply permissions
```bash
sudo chmod +x /usr/local/bin/docker-compose
sudo chmod 666 /var/run/docker.sock
```

#### Loading testing data from CSV Files (Optional) ####
A number of csv files containing testing data are located in 
```
~/capstoneproject-comp3900-w16a-jamar/backend/migrations
```

1. To load these data, first run 
```bash
cd ~/capstoneproject-comp3900-w16a-jamar/backend
sudo service postgresql stop
make db-stop
make db-start
make migrate-reset
make migrate-down
make migrate-up

make load-data # will load lots of dummy data
make core-run
```
