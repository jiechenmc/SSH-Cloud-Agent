# sudo find . -mtime +1 -type f -name "*" -delete
# Add Docker's official GPG key:

command -v docker &> /dev/null


if [ $? -eq 0 ]; then
  echo "Copy successful"
else
   sudo apt-get update
    sudo apt-get install ca-certificates curl gnupg -y
    sudo install -m 0755 -d /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    sudo chmod a+r /etc/apt/keyrings/docker.gpg

    # Add the repository to Apt sources:
    echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
fi

command -v screen &> /dev/null

if [ $? -eq 0 ]; then
  echo "Copy successful"
else
   sudo apt-get update
   sudo apt install screen -y
fi
   
sudo apt update && sudo apt upgrade -y && sudo apt autoremove -y
sudo docker compose down
sudo docker compose up --build -d
sudo docker system prune -f
