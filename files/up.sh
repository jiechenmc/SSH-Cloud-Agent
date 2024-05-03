# sudo find . -mtime +1 -type f -name "*" -delete
# Add Docker's official GPG key:

# command -v docker &> /dev/null


# if [ $? -eq 0 ]; then
#   echo "Copy successful"
# else
#    sudo apt-get update
#     sudo apt-get install ca-certificates curl gnupg -y
#     sudo install -m 0755 -d /etc/apt/keyrings
#     curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
#     sudo chmod a+r /etc/apt/keyrings/docker.gpg

#     # Add the repository to Apt sources:
#     echo \
#     "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
#     $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
#     sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
#     sudo apt-get update
#     sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
# fi

# command -v screen &> /dev/null

# if [ $? -eq 0 ]; then
#   echo "Copy successful"
# else
#    sudo apt-get update
#    sudo apt install screen -y
# fi
   
# sudo apt update && sudo apt upgrade -y && sudo apt autoremove -y
# sudo docker compose down
# sudo docker compose up --build -d
# sudo docker system prune -f

sudo apt-get update -y; sudo apt-get upgrade -y; sudo apt-get install libuv1-dev -y 
# sudo ufw enable; sudo ufw allow 10000/tcp; sudo ufw allow 20000/tcp; sudo ufw allow 22/tcp; sudo ufw status verbose;
export POMPE_HOME=/users/chenj/Pompe-HS

sudo rm -rf $POMPE_HOME
git clone https://github.com/jiechenmc/Pompe-HS.git

sudo chmod +w Pompe-HS/**/

sudo rm -rf $POMPE_HOME/experiments/pompe/conf-distributed/*

echo "\nexport POMPE_HOME=/users/chenj/Pompe-HS" > .zshrc
echo "\nexport POMPE_HOME=/users/chenj/Pompe-HS" >> /etc/environment
source .zshrc

cd $POMPE_HOME
git submodule update --init --recursive
$POMPE_HOME/install_deps.sh

cd $POMPE_HOME/libhotstuff
sudo apt install doxygen -y
cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED=ON -DHOTSTUFF_PROTO_LOG=ON
make

cd examples
sudo make
